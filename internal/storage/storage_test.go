package storage

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/v1tbrah/media-service/config"
)

func TestInit(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(cfg.LogLvl)

	if err := cfg.ParseEnv(); err != nil {
		t.Fatalf("config.ParseEnv: %v", err)
	}
	zerolog.SetGlobalLevel(cfg.LogLvl)

	_, err := Init(ctx, cfg.Minio)
	if err != nil {
		t.Fatalf("init cache: %v", err)
	}
}
