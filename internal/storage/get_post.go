package storage

import (
	"context"
	"io"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var ErrNotFoundByGUID = errors.New("not found by guid")

func (s *Minio) GetPost(ctx context.Context, guid string) ([]byte, error) {
	obj, err := s.cli.GetObject(ctx, s.cfg.PostBucketName, guid, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "cli.GetObject")
	}

	defer func() {
		if err = obj.Close(); err != nil {
			log.Error().Err(err).Msg("obj.Close")
		}
	}()

	info, err := obj.Stat()
	if err != nil {
		if mErr, ok := err.(minio.ErrorResponse); ok && mErr.StatusCode == http.StatusNotFound {
			return nil, ErrNotFoundByGUID
		}
		return nil, errors.Wrap(err, "obj.Stat")
	}

	data := make([]byte, info.Size)
	_, err = obj.Read(data)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, errors.Wrap(err, "obj.Read")
	}

	return data, nil
}
