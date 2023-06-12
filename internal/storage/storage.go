package storage

import (
	"context"
	"net"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"

	"github.com/v1tbrah/media-service/config"
)

type Minio struct {
	cli *minio.Client

	cfg config.Minio
}

func Init(ctx context.Context, cfg config.Minio) (*Minio, error) {
	cli, err := minio.New(
		net.JoinHostPort(cfg.Host, cfg.Port),
		&minio.Options{Creds: credentials.NewStaticV2(cfg.AccessKey, cfg.SecretKey, "")},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "minio.New, cfg:\n%+v", cfg)
	}

	if cli.IsOffline() {
		return nil, errors.New("cli is offline")
	}

	isExists, err := cli.BucketExists(ctx, cfg.PostBucketName)
	if err != nil {
		return nil, errors.Wrapf(err, "cli.BucketExists %s", cfg.PostBucketName)
	}

	if !isExists {
		if err = cli.MakeBucket(ctx, cfg.PostBucketName, minio.MakeBucketOptions{}); err != nil {
			return nil, errors.Wrapf(err, "cli.MakeBucket %s", cfg.PostBucketName)
		}
	}

	return &Minio{
		cli: cli,
		cfg: cfg,
	}, nil
}
