package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/v1tbrah/media-service/config"
	"github.com/v1tbrah/media-service/internal/api"
	"github.com/v1tbrah/media-service/internal/storage"
)

func main() {
	newConfig := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	if err := newConfig.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	ctxStart, ctxStartCancel := context.WithCancel(context.Background())

	newStorage, err := storage.Init(ctxStart, newConfig.Minio)
	if err != nil {
		log.Fatal().Err(err).Interface("config", newConfig.Minio).Msg("storage.Init")
	} else {
		log.Info().Msg("storage initialized")
	}

	newAPI := api.New(newConfig.HTTP, newStorage)

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errServingCh := make(chan error)
	go func() {
		errServing := newAPI.StartServing(ctxStart, newConfig.GRPC, shutdownSig)
		errServingCh <- errServing
	}()

	select {
	case shutdownSigValue := <-shutdownSig:
		close(shutdownSig)
		log.Info().Msgf("Shutdown signal received: %s", strings.ToUpper(shutdownSigValue.String()))
	case errServing := <-errServingCh:
		if errServing != nil {
			log.Error().Err(errServing).Msg("newAPI.StartServing")
		}
	}

	ctxStartCancel()

	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = newAPI.GracefulStop(ctxClose); err != nil {
		log.Error().Err(err).Msg("gRPC and HTTP server graceful stop")
		if err == context.DeadlineExceeded {
			return
		}
	} else {
		log.Info().Msg("gRPC and HTTP server gracefully stopped")
	}
}
