# RocksDB gRPC Service

A gRPC service providing a robust interface to RocksDB, offering key-value storage operations with support for streaming queries.

## Features

- Basic operations:
  - Put: Store a key-value pair
  - Get: Retrieve a value by key
  - Delete: Remove a key-value pair
- Advanced operations:
  - StreamGet: Stream multiple key-value pairs using:
    - Prefix search (key*)
    - Multiple exact keys

## Prerequisites

- Go 1.21 or later
- RocksDB installed (via Homebrew on macOS)
- Protocol Buffers compiler

## Installation

### Install RocksDB (macOS)

```bash
brew install rocksdb
```

### Install Protocol Buffers compiler

```bash
brew install protobuf
```

### Install Go dependencies

```bash
go mod tidy
```

## Building

1. Generate Protocol Buffer code:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/proto/rocksdb.proto
```

2. Build the service:
```bash
CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/9.7.4/include" \
CGO_LDFLAGS="-L/opt/homebrew/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
go build -o rocksdb-service ./cmd/server
```

3. Build the client:
```bash
go build -o rocksdb-client ./cmd/client
```

Note: Adjust the RocksDB include path according to your installed version.

## Usage

### Server

Start the service:
```bash
./rocksdb-service [flags]
```

Available flags:
- `--port`: The server port (default: 50051)
- `--db-path`: Path to RocksDB data directory (default: /data/rocksdb)

### Client Usage

The service includes a command-line client that supports basic operations.

The client supports the following operations:

1. Put a value:
```bash
./rocksdb-client -op put -key mykey -value "my value" [-db mydb] [-server localhost:50051]
```

2. Get a value:
```bash
./rocksdb-client -op get -key mykey [-db mydb] [-server localhost:50051]
```

3. Delete a value:
```bash
./rocksdb-client -op delete -key mykey [-db mydb] [-server localhost:50051]
```

Available flags:
- `-server`: The server address (default: localhost:50051)
- `-db`: Database name to use (default: default)
- `-key`: Key to operate on (required)
- `-value`: Value to put (required for put operation)
- `-op`: Operation to perform: put, get, or delete (required)

## Multi-Database Support

The service supports multiple RocksDB databases. Each database is stored in a separate subdirectory under the main data directory.

### Database Names
- Each request must specify a database name
- Database names must be non-empty and should contain only valid filesystem characters
- Databases are created automatically on first use
- Each database is isolated from others and stored in its own directory

### Example Usage
When making requests to the service, include the database name in each request:

```proto
// Store data in "users" database
rpc Put {
    database_name: "users"
    key: "user:1"
    value: "..."
}

// Retrieve data from "products" database
rpc Get {
    database_name: "products"
    key: "product:1"
}
```

The databases will be created in subdirectories under the specified data directory:
```
/data/
    /users/
        [RocksDB files]
    /products/
        [RocksDB files]
```

## API

For detailed API documentation, refer to the protobuf definitions in `api/proto/rocksdb.proto`.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
