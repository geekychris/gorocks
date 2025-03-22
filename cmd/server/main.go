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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "rocksdb-service/api/proto"
	"rocksdb-service/internal/db"
)

type server struct {
	pb.UnimplementedRocksDBServiceServer
	dbManager *db.DBManager
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	database, err := s.dbManager.GetDB(req.DatabaseName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get database: %v", err)
	}

	err = database.Put(req.Key, req.Value)
	if err != nil {
		return &pb.PutResponse{Success: false, Error: err.Error()}, nil
	}
	return &pb.PutResponse{Success: true}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	database, err := s.dbManager.GetDB(req.DatabaseName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get database: %v", err)
	}

	value, exists, err := database.Get(req.Key)
	if err != nil {
		return &pb.GetResponse{Found: false, Error: err.Error()}, nil
	}
	return &pb.GetResponse{Value: value, Found: exists}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	database, err := s.dbManager.GetDB(req.DatabaseName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get database: %v", err)
	}

	err = database.Delete(req.Key)
	if err != nil {
		return &pb.DeleteResponse{Success: false, Error: err.Error()}, nil
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (s *server) StreamGet(req *pb.StreamGetRequest, stream pb.RocksDBService_StreamGetServer) error {
	database, err := s.dbManager.GetDB(req.DatabaseName)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to get database: %v", err)
	}

	var ch chan db.KeyValuePair

	switch query := req.Query.(type) {
	case *pb.StreamGetRequest_Prefix:
		ch = database.GetByPrefix(query.Prefix)
	case *pb.StreamGetRequest_Keys:
		ch = database.GetMultiple(query.Keys.Keys)
	default:
		return fmt.Errorf("invalid query type")
	}

	for pair := range ch {
		if pair.Err != nil {
			return status.Errorf(codes.Internal, "stream error: %v", pair.Err)
		}

		err := stream.Send(&pb.StreamGetResponse{
			Key:   pair.Key,
			Value: pair.Value,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send response: %v", err)
		}
	}

	return nil
}

func main() {
	var (
		port   = flag.Int("port", 50051, "The server port")
		dbPath = flag.String("db-path", "/data/rocksdb", "Path to RocksDB data directory")
	)
	flag.Parse()

	// Initialize DBManager
	dbManager := db.NewDBManager(*dbPath)
	defer dbManager.Close()

	// Initialize gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRocksDBServiceServer(s, &server{dbManager: dbManager})

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
