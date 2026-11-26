package backup

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Config describes where snapshots are stored and how to reach the bucket.
type S3Config struct {
	// Bucket is the S3 bucket snapshots are uploaded to and restored from.
	Bucket string
	// Prefix is prepended to every object key, e.g. "snapshots".
	Prefix string
	// Region is the AWS region to sign requests for.
	Region string
	// Endpoint overrides the default AWS endpoint, for S3-compatible stores
	// such as MinIO or an in-memory test fake. Empty means use real AWS S3.
	Endpoint string
	// PathStyle forces path-style addressing (bucket in the URL path rather
	// than the host), which most S3-compatible stores require.
	PathStyle bool
}

// NewS3Client builds an S3 client from cfg, overriding the endpoint and
// addressing style when cfg.Endpoint is set so the same code drives both
// real AWS S3 and local S3-compatible mocks.
func NewS3Client(ctx context.Context, cfg S3Config) (*s3.Client, error) {
	if cfg.Bucket == "" {
		return nil, fmt.Errorf("S3Config.Bucket must not be empty")
	}

	opts := []func(*config.LoadOptions) error{}
	if cfg.Region != "" {
		opts = append(opts, config.WithRegion(cfg.Region))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
		}
		o.UsePathStyle = cfg.PathStyle
	})
	return client, nil
}
