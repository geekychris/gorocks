package backup

import (
	"context"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"

	"rocksdb-service/internal/db"
)

const testBucket = "gorocks-test-bucket"

// newTestS3 starts an in-memory, in-process S3-compatible server and
// returns a client configured against it with testBucket already created.
func newTestS3(t *testing.T) *s3.Client {
	t.Helper()

	backend := s3mem.New()
	faker := gofakes3.New(backend)
	ts := httptest.NewServer(faker.Server())
	t.Cleanup(ts.Close)

	client := s3.New(s3.Options{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(ts.URL),
		UsePathStyle: true,
		Credentials:  credentials.NewStaticCredentialsProvider("KEY", "SECRET", ""),
	})

	ctx := context.Background()
	if _, err := client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(testBucket)}); err != nil {
		t.Fatalf("CreateBucket: %v", err)
	}
	return client
}

func seedRocksDB(t *testing.T, path string) map[string][]byte {
	t.Helper()

	rdb, err := db.NewRocksDB(path)
	if err != nil {
		t.Fatalf("NewRocksDB: %v", err)
	}
	defer rdb.Close()

	data := map[string][]byte{
		"user/1":    []byte("alice"),
		"user/2":    []byte("bob"),
		"user/20":   []byte("overlap"),
		"order/1":   []byte("widget"),
		"zzz/empty": []byte(""),
	}
	for k, v := range data {
		if err := rdb.Put(k, v); err != nil {
			t.Fatalf("Put(%s): %v", k, err)
		}
	}
	return data
}

func assertRocksDBContents(t *testing.T, path string, want map[string][]byte) {
	t.Helper()

	rdb, err := db.NewRocksDB(path)
	if err != nil {
		t.Fatalf("NewRocksDB(%s): %v", path, err)
	}
	defer rdb.Close()

	for k, wantV := range want {
		gotV, exists, err := rdb.Get(k)
		if err != nil {
			t.Fatalf("Get(%s): %v", k, err)
		}
		if !exists {
			t.Fatalf("Get(%s): key missing after restore", k)
		}
		if string(gotV) != string(wantV) {
			t.Fatalf("Get(%s) = %q, want %q", k, gotV, wantV)
		}
	}

	gotByPrefix := map[string][]byte{}
	for pair := range rdb.GetByPrefix("user/") {
		if pair.Err != nil {
			t.Fatalf("GetByPrefix: %v", pair.Err)
		}
		gotByPrefix[pair.Key] = pair.Value
	}
	wantByPrefix := map[string][]byte{
		"user/1":  want["user/1"],
		"user/2":  want["user/2"],
		"user/20": want["user/20"],
	}
	if len(gotByPrefix) != len(wantByPrefix) {
		t.Fatalf("GetByPrefix(user/) returned %d keys, want %d", len(gotByPrefix), len(wantByPrefix))
	}
	for k, v := range wantByPrefix {
		if string(gotByPrefix[k]) != string(v) {
			t.Fatalf("GetByPrefix(user/)[%s] = %q, want %q", k, gotByPrefix[k], v)
		}
	}
}

func TestSnapshotThenRestoreRoundTrip(t *testing.T) {
	baseDir := t.TempDir()
	client := newTestS3(t)
	snapshotter := NewSnapshotter(client, S3Config{Bucket: testBucket, Prefix: "snapshots"})

	srcPath := filepath.Join(baseDir, "src")
	want := seedRocksDB(t, srcPath)

	rdb, err := db.NewRocksDB(srcPath)
	if err != nil {
		t.Fatalf("NewRocksDB(reopen): %v", err)
	}

	key, err := snapshotter.Snapshot(context.Background(), rdb, "mydb")
	rdb.Close()
	if err != nil {
		t.Fatalf("Snapshot: %v", err)
	}
	if key == "" {
		t.Fatalf("Snapshot returned empty key")
	}

	destPath := filepath.Join(baseDir, "restored")
	if err := snapshotter.Restore(context.Background(), key, destPath); err != nil {
		t.Fatalf("Restore: %v", err)
	}

	assertRocksDBContents(t, destPath, want)
}

func TestRestoreRefusesExistingDestDir(t *testing.T) {
	baseDir := t.TempDir()
	client := newTestS3(t)
	snapshotter := NewSnapshotter(client, S3Config{Bucket: testBucket, Prefix: "snapshots"})

	srcPath := filepath.Join(baseDir, "src")
	seedRocksDB(t, srcPath)

	rdb, err := db.NewRocksDB(srcPath)
	if err != nil {
		t.Fatalf("NewRocksDB(reopen): %v", err)
	}
	key, err := snapshotter.Snapshot(context.Background(), rdb, "mydb")
	rdb.Close()
	if err != nil {
		t.Fatalf("Snapshot: %v", err)
	}

	destPath := filepath.Join(baseDir, "existing")
	if err := os.MkdirAll(destPath, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	sentinel := filepath.Join(destPath, "sentinel.txt")
	if err := os.WriteFile(sentinel, []byte("do not touch"), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	if err := snapshotter.Restore(context.Background(), key, destPath); err == nil {
		t.Fatalf("Restore into an existing directory should fail")
	}

	data, err := os.ReadFile(sentinel)
	if err != nil {
		t.Fatalf("ReadFile(sentinel): %v", err)
	}
	if string(data) != "do not touch" {
		t.Fatalf("existing destDir contents were modified: %q", data)
	}
}
