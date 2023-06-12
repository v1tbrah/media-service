package storage

import (
	"bytes"
	"context"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

func (s *Minio) AddPost(ctx context.Context, data []byte) (guid string, err error) {
	var buf bytes.Buffer
	size, err := buf.Write(data)
	if err != nil {
		return "", errors.Wrap(err, "buf.Write")
	}

	guid = uuid.NewString()

	_, err = s.cli.PutObject(ctx, s.cfg.PostBucketName, guid, &buf, int64(size), minio.PutObjectOptions{})
	if err != nil {
		return "", errors.Wrap(err, "cli.PutObject")
	}

	return guid, nil
}
