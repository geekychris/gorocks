package backup

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"rocksdb-service/internal/db"
)

// Snapshotter takes RocksDB checkpoints and moves them to and from S3.
type Snapshotter struct {
	S3  *s3.Client
	Cfg S3Config
}

// NewSnapshotter builds a Snapshotter backed by client, storing snapshots
// per cfg.
func NewSnapshotter(client *s3.Client, cfg S3Config) *Snapshotter {
	return &Snapshotter{S3: client, Cfg: cfg}
}

// Snapshot takes a consistent checkpoint of rdb and uploads it to S3 as a
// single tar.gz object under a timestamped key. It returns the key the
// snapshot was stored at.
func (s *Snapshotter) Snapshot(ctx context.Context, rdb *db.RocksDB, dbName string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "gorocks-snapshot-*")
	if err != nil {
		return "", fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	checkpointDir := filepath.Join(tmpDir, "checkpoint")
	if err := rdb.Checkpoint(checkpointDir); err != nil {
		return "", fmt.Errorf("checkpoint database %s: %w", dbName, err)
	}

	archivePath := filepath.Join(tmpDir, "snapshot.tar.gz")
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf("create archive file: %w", err)
	}
	defer archiveFile.Close()

	if err := tarGzDir(checkpointDir, archiveFile); err != nil {
		return "", fmt.Errorf("archive checkpoint for %s: %w", dbName, err)
	}
	if _, err := archiveFile.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("seek archive file: %w", err)
	}

	key := s.objectKey(dbName)

	uploader := manager.NewUploader(s.S3)
	if _, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Cfg.Bucket),
		Key:    aws.String(key),
		Body:   archiveFile,
	}); err != nil {
		return "", fmt.Errorf("upload snapshot %s: %w", key, err)
	}

	return key, nil
}

// Restore downloads the snapshot stored at key and rebuilds a database
// directory from it at destDir. destDir must not already exist: Restore
// only targets cold, unopened directories and never overwrites one.
func (s *Snapshotter) Restore(ctx context.Context, key, destDir string) error {
	if _, err := os.Stat(destDir); err == nil {
		return fmt.Errorf("restore destination %s already exists", destDir)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat restore destination %s: %w", destDir, err)
	}

	tmpFile, err := os.CreateTemp("", "gorocks-restore-*.tar.gz")
	if err != nil {
		return fmt.Errorf("create temp file for download: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	downloader := manager.NewDownloader(s.S3)
	if _, err := downloader.Download(ctx, tmpFile, &s3.GetObjectInput{
		Bucket: aws.String(s.Cfg.Bucket),
		Key:    aws.String(key),
	}); err != nil {
		return fmt.Errorf("download snapshot %s: %w", key, err)
	}

	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("seek downloaded snapshot: %w", err)
	}

	if err := untarGz(tmpFile, destDir); err != nil {
		return fmt.Errorf("extract snapshot %s into %s: %w", key, destDir, err)
	}

	rdb, err := db.NewRocksDB(destDir)
	if err != nil {
		return fmt.Errorf("open restored database at %s for verification: %w", destDir, err)
	}
	rdb.Close()

	return nil
}

// objectKey returns the timestamped S3 key a snapshot of dbName is stored
// at, e.g. "snapshots/mydb/20060102T150405Z.tar.gz".
func (s *Snapshotter) objectKey(dbName string) string {
	ts := time.Now().UTC().Format("20060102T150405Z")
	prefix := strings.Trim(s.Cfg.Prefix, "/")
	if prefix == "" {
		return fmt.Sprintf("%s/%s.tar.gz", dbName, ts)
	}
	return fmt.Sprintf("%s/%s/%s.tar.gz", prefix, dbName, ts)
}
