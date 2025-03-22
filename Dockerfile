# Build stage
FROM golang:1.21-bullseye AS builder

# Install RocksDB dependencies and protoc
RUN apt-get update && apt-get install -y \
    librocksdb-dev \
    libsnappy-dev \
    liblz4-dev \
    libzstd-dev \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf code
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/proto/rocksdb.proto

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o /rocksdb-service ./cmd/server

# Final stage
FROM debian:bullseye-slim

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    librocksdb-dev \
    libsnappy-dev \
    liblz4-dev \
    libzstd-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /rocksdb-service .

# Create data directory
RUN mkdir -p /data/rocksdb && \
    chmod 777 /data/rocksdb

EXPOSE 50051
VOLUME ["/data/rocksdb"]

ENTRYPOINT ["./rocksdb-service"]

