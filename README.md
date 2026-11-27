<p align="center">
  <img src="assets/logo.png" alt="Gorocks logo" width="180">
</p>

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

- Go 1.24 or later
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

## Backup / restore

`cmd/backup` is a standalone CLI that takes a consistent RocksDB checkpoint and
uploads it to S3 (or any S3-compatible store), and can later rebuild a
database directory from a stored snapshot. Replication protects against node
loss but not against logical corruption or an operator error replicated to
every replica; offsite snapshots cover that gap.

Build it the same way as `cmd/server` (it needs the same RocksDB cgo
dependencies):

```bash
CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/9.7.4/include" \
CGO_LDFLAGS="-L/opt/homebrew/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
go build -o backup ./cmd/backup
```

### Flags

- `-op`: `snapshot` or `restore` (required)
- `-db-path`: path to the RocksDB data directory to snapshot (snapshot only; default `/data/rocksdb`)
- `-db`: logical database name, used in the snapshot's S3 key (snapshot only, required)
- `-bucket`: S3 bucket to upload to / restore from (required)
- `-prefix`: S3 key prefix snapshots are stored under (default `snapshots`)
- `-region`: AWS region to sign requests for (default `us-east-1`)
- `-endpoint`: S3-compatible endpoint override, e.g. a local MinIO URL (default: real AWS S3)
- `-path-style`: use path-style S3 addressing (most S3-compatible stores require this)
- `-key`: snapshot S3 key to restore (restore only, required)
- `-out`: destination directory to rebuild the database into; must not already exist (restore only, required)

Credentials are resolved the normal AWS SDK way (environment variables,
shared config/credentials files, instance/task role, etc.) — there is no
credentials flag.

### Example: snapshot and restore against a local MinIO

Start MinIO and create a bucket, then take a snapshot:

```bash
./backup -op=snapshot \
  -db-path=/data/rocksdb/users \
  -db=users \
  -bucket=gorocks-snapshots \
  -endpoint=http://localhost:9000 \
  -path-style
# prints the snapshot's key, e.g.: snapshots/users/20261126T120000Z.tar.gz
```

Restore it into a fresh directory:

```bash
./backup -op=restore \
  -key=snapshots/users/20261126T120000Z.tar.gz \
  -out=/data/restore/users \
  -bucket=gorocks-snapshots \
  -endpoint=http://localhost:9000 \
  -path-style
```

`restore` refuses to run if `-out` already exists, so it never clobbers a
directory a running server has open — stop the server, or restore to a fresh
path, before pointing it at the restored data.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
