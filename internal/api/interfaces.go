package api

import (
	"context"
)

//go:generate mockery --name Storage
type Storage interface {
	AddPost(ctx context.Context, data []byte) (guid string, err error)
	GetPost(ctx context.Context, guid string) (data []byte, err error)
}
