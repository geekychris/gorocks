// Command backup takes RocksDB checkpoints and ships them to S3, or
// restores a database directory from a previously stored snapshot.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"rocksdb-service/internal/backup"
	"rocksdb-service/internal/db"
)

func main() {
	var (
		op        = flag.String("op", "", "operation to perform: snapshot|restore")
		dbPath    = flag.String("db-path", "/data/rocksdb", "path to the RocksDB data directory to snapshot (snapshot only)")
		dbName    = flag.String("db", "", "logical database name, used in the snapshot's S3 key (snapshot only)")
		bucket    = flag.String("bucket", "", "S3 bucket to upload snapshots to or restore them from")
		prefix    = flag.String("prefix", "snapshots", "S3 key prefix snapshots are stored under")
		region    = flag.String("region", "us-east-1", "AWS region to sign S3 requests for")
		endpoint  = flag.String("endpoint", "", "S3-compatible endpoint override, e.g. a local MinIO URL (default: real AWS S3)")
		pathStyle = flag.Bool("path-style", false, "use path-style S3 addressing (required by most S3-compatible stores)")
		key       = flag.String("key", "", "snapshot S3 key to restore (restore only)")
		out       = flag.String("out", "", "destination directory to rebuild the database into; must not already exist (restore only)")
	)
	flag.Parse()

	ctx := context.Background()

	cfg := backup.S3Config{
		Bucket:    *bucket,
		Prefix:    *prefix,
		Region:    *region,
		Endpoint:  *endpoint,
		PathStyle: *pathStyle,
	}

	client, err := backup.NewS3Client(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to build S3 client: %v", err)
	}
	snapshotter := backup.NewSnapshotter(client, cfg)

	switch *op {
	case "snapshot":
		if *dbName == "" {
			log.Fatal("-db is required for -op=snapshot")
		}
		rdb, err := db.NewRocksDB(*dbPath)
		if err != nil {
			log.Fatalf("failed to open database at %s: %v", *dbPath, err)
		}
		defer rdb.Close()

		snapshotKey, err := snapshotter.Snapshot(ctx, rdb, *dbName)
		if err != nil {
			log.Fatalf("snapshot failed: %v", err)
		}
		fmt.Println(snapshotKey)

	case "restore":
		if *key == "" {
			log.Fatal("-key is required for -op=restore")
		}
		if *out == "" {
			log.Fatal("-out is required for -op=restore")
		}
		if err := snapshotter.Restore(ctx, *key, *out); err != nil {
			log.Fatalf("restore failed: %v", err)
		}
		fmt.Printf("restored %s to %s\n", *key, *out)

	default:
		log.Fatalf("unknown -op %q: must be \"snapshot\" or \"restore\"", *op)
	}
}
