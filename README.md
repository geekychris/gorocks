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

Note: Adjust the RocksDB include path according to your installed version.

## Usage

Start the service:
```bash
./rocksdb-service [flags]
```

### Available Flags

- `--port`: The server port (default: 50051)
- `--db-path`: Path to RocksDB data directory (default: /data/rocksdb)

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
