package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "google.golang.org/grpc"
    pb "rocksdb-service/api/proto"
    "rocksdb-service/internal/db"
)

type server struct {
    pb.UnimplementedRocksDBServiceServer
    db *db.RocksDB
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
    err := s.db.Put(req.Key, req.Value)
    if err != nil {
        return &pb.PutResponse{Success: false, Error: err.Error()}, nil
    }
    return &pb.PutResponse{Success: true}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
    value, exists, err := s.db.Get(req.Key)
    if err != nil {
        return &pb.GetResponse{Found: false, Error: err.Error()}, nil
    }
    return &pb.GetResponse{Value: value, Found: exists}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
    err := s.db.Delete(req.Key)
    if err != nil {
        return &pb.DeleteResponse{Success: false, Error: err.Error()}, nil
    }
    return &pb.DeleteResponse{Success: true}, nil
}

func (s *server) StreamGet(req *pb.StreamGetRequest, stream pb.RocksDBService_StreamGetServer) error {
    var ch chan db.KeyValuePair

    switch query := req.Query.(type) {
    case *pb.StreamGetRequest_Prefix:
        ch = s.db.GetByPrefix(query.Prefix)
    case *pb.StreamGetRequest_Keys:
        ch = s.db.GetMultiple(query.Keys.Keys)
    default:
        return fmt.Errorf("invalid query type")
    }

    for pair := range ch {
        if pair.Err != nil {
            return stream.Send(&pb.StreamGetResponse{
                Error: pair.Err.Error(),
            })
        }

        err := stream.Send(&pb.StreamGetResponse{
            Key:   pair.Key,
            Value: pair.Value,
        })
        if err != nil {
            return err
        }
    }

    return nil
}

func main() {
    var (
        port = flag.Int("port", 50051, "The server port")
        dbPath = flag.String("db-path", "/data/rocksdb", "Path to RocksDB data directory")
    )
    flag.Parse()

    // Initialize RocksDB
    rocksDB, err := db.NewRocksDB(*dbPath)
    if err != nil {
        log.Fatalf("Failed to initialize RocksDB: %v", err)
    }
    defer rocksDB.Close()

    // Initialize gRPC server
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterRocksDBServiceServer(s, &server{db: rocksDB})

    // Handle graceful shutdown
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        log.Printf("Server listening at %v", lis.Addr())
        if err := s.Serve(lis); err != nil {
            log.Fatalf("Failed to serve: %v", err)
        }
    }()

    <-stop
    log.Println("Shutting down server...")
    s.GracefulStop()
}

