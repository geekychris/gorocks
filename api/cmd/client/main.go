package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "rocksdb-service/api/proto"
)

func main() {
	var (
		serverAddr = flag.String("server", "localhost:50051", "The server address in the format of host:port")
		dbName     = flag.String("db", "default", "Database name to use")
		operation  = flag.String("op", "", "Operation to perform: put, get, delete, or prefix")
		key        = flag.String("key", "", "Key to operate on")
		value      = flag.String("value", "", "Value to put (only used with put operation)")
		prefix     = flag.String("prefix", "", "Key prefix to search for (only used with prefix operation)")
	)
	flag.Parse()

	if *operation == "" {
		log.Fatal("Operation is required")
	}

	if (*operation == "put" || *operation == "get" || *operation == "delete") && *key == "" {
		log.Fatal("Key is required for put, get, and delete operations")
	}

	if *operation == "prefix" && *prefix == "" {
		log.Fatal("Prefix is required for prefix operation")
	}

	if *operation == "put" && *value == "" {
		log.Fatal("Value is required for put operation")
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRocksDBServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	switch *operation {
	case "put":
		resp, err := client.Put(ctx, &pb.PutRequest{
			DatabaseName: *dbName,
			Key:          *key,
			Value:        []byte(*value),
		})
		if err != nil {
			log.Fatalf("Put failed: %v", err)
		}
		if !resp.Success {
			log.Fatalf("Put failed: %s", resp.Error)
		}
		fmt.Println("Put successful")

	case "get":
		resp, err := client.Get(ctx, &pb.GetRequest{
			DatabaseName: *dbName,
			Key:          *key,
		})
		if err != nil {
			log.Fatalf("Get failed: %v", err)
		}
		if !resp.Found {
			fmt.Println("Key not found")
			return
		}
		fmt.Printf("Value: %s\n", string(resp.Value))

	case "delete":
		resp, err := client.Delete(ctx, &pb.DeleteRequest{
			DatabaseName: *dbName,
			Key:          *key,
		})
		if err != nil {
			log.Fatalf("Delete failed: %v", err)
		}
		if !resp.Success {
			log.Fatalf("Delete failed: %s", resp.Error)
		}
		fmt.Println("Delete successful")

	case "prefix":
		stream, err := client.StreamGet(ctx, &pb.StreamGetRequest{
			DatabaseName: *dbName,
			Query: &pb.StreamGetRequest_Prefix{
				Prefix: *prefix,
			},
		})
		if err != nil {
			log.Fatalf("StreamGet failed: %v", err)
		}

		count := 0
		fmt.Println("Keys with prefix:", *prefix)
		for {
			resp, err := stream.Recv()
			if err != nil {
				// End of stream
				break
			}
			if resp.Error != "" {
				fmt.Printf("Error for key %s: %s\n", resp.Key, resp.Error)
				continue
			}
			fmt.Printf("Key: %s, Value: %s\n", resp.Key, string(resp.Value))
			count++
		}
		fmt.Printf("Found %d key-value pairs with prefix: %s\n", count, *prefix)

	default:
		log.Fatalf("Unknown operation: %s", *operation)
	}
}
