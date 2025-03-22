# RocksDB gRPC Service

This service provides a gRPC interface to RocksDB, offering key-value storage operations with support for streaming queries.

## Features

- Basic operations:
  - Put: Store a key-value pair
  - Get: Retrieve a value by key
  - Delete: Remove a key-value pair
- Advanced operations:
  - StreamGet: Stream multiple key-value pairs using:
    - Prefix search (key*)
    - Multiple exact keys

## Requirements

- Go 1.21 or later
- RocksDB and its dependencies
- Protocol Buffers compiler (protoc)

## Building from Source

1. Install RocksDB:

   ```bash
   # On macOS with Homebrew
   brew install rocksdb

   # On Ubuntu/Debian
   sudo apt-get install -y librocksdb-dev \
       libsnappy-dev \
       liblz4-dev \
       libzstd-dev \
       libgflags-dev \
       libbz2-dev
   ```

2. Verify RocksDB installation:

   ```bash
   # On macOS
   brew list rocksdb  # This will show the installation paths
   
   # On Ubuntu/Debian
   dpkg -L librocksdb-dev
   ```
2. Install Protocol Buffers compiler:

   ```bash
   # Ubuntu/Debian
   sudo apt-get install -y protobuf-compiler 

   # macOS
   brew install protobuf
   ```

3. Install Go protobuf and gRPC plugins:

   ```bash
   # Install protoc plugins for Go
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

   # Add Go bin directory to PATH (add this to your .bashrc or .zshrc for permanence)
   export PATH="$PATH:$(go env GOPATH)/bin"

   # Verify installations
   which protoc-gen-go
   which protoc-gen-go-grpc
   ```

4. Generate protobuf code (run from project root directory):

   ```bash
   # Verify you are in the project root
   ls api/proto/rocksdb.proto

   # Generate the protobuf and gRPC code
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       api/proto/rocksdb.proto

   # Verify the generated files exist
   ls -l api/proto/rocksdb.pb.go api/proto/rocksdb_grpc.pb.go
   ```

5. Build the service:

   First, ensure you have the correct RocksDB paths:

   ```bash
   # On macOS with Homebrew, find RocksDB path
   export ROCKSDB_PATH=$(brew --prefix rocksdb)
   
   # Build the service
   CGO_CFLAGS="-I${ROCKSDB_PATH}/include" \
   CGO_LDFLAGS="-L${ROCKSDB_PATH}/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
   go build -o rocksdb-service ./cmd/server

   # Alternative: if you know the exact paths, you can use them directly:
   # For Apple Silicon (M1/M2) Macs:
   CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/9.7.4/include" \
   CGO_LDFLAGS="-L/opt/homebrew/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
   go build -o rocksdb-service ./cmd/server

   # For Intel Macs:
   CGO_CFLAGS="-I/usr/local/Cellar/rocksdb/9.7.4/include" \
   CGO_LDFLAGS="-L/usr/local/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
   go build -o rocksdb-service ./cmd/server

   # For Ubuntu/Debian:
   CGO_CFLAGS="-I/usr/include/rocksdb" \
   CGO_LDFLAGS="-L/usr/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
   go build -o rocksdb-service ./cmd/server
   ```

   Note: The build uses the `github.com/linxGnu/grocksdb` package which is actively maintained and compatible with recent RocksDB versions.

## Docker Build and Run

1. Build the Docker image:

   ```bash
   docker build -t rocksdb-service .
   ```

2. Run the container:

   ```bash
   docker run -d \
       -p 50051:50051 \
       -v rocksdb-data:/data/rocksdb \
       rocksdb-service
   ```

## Usage

The service listens on port 50051 by default. You can modify the port and database path using command-line flags:

```bash
./rocksdb-service --port 50051 --db-path /data/rocksdb
```

### API Examples

Using a gRPC client, you can:

1. Store a value:
   ```protobuf
   PutRequest {
       key: "user:1",
       value: <bytes>
   }
   ```

2. Retrieve a value:
   ```protobuf
   GetRequest {
       key: "user:1"
   }
   ```

3. Stream values by prefix:
   ```protobuf
   StreamGetRequest {
       prefix: "user:"
   }
   ```

4. Stream multiple exact keys:
   ```protobuf
   StreamGetRequest {
       keys: {
           keys: ["user:1", "user:2", "user:3"]
       }
   }
   ```

## Implementation Details

The service uses:
- RocksDB for persistent key-value storage
- gRPC for the network protocol
- Protocol Buffers for data serialization
- Streaming responses for prefix searches and multi-key requests
- Graceful shutdown handling

The RocksDB wrapper (`internal/db/rocksdb.go`) provides a clean interface for database operations, while the gRPC service (`cmd/server/main.go`) handles the network protocol and request processing.

# gorocks
