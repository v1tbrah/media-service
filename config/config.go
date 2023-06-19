package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	defaultLogLvl = zerolog.InfoLevel
	envNameLogLvl = "LOG_LVL"
)

type Config struct {
	GRPC   GRPC
	HTTP   HTTP
	Minio  Minio
	LogLvl zerolog.Level
}

func NewDefaultConfig() Config {
	return Config{
		GRPC:   newDefaultGRPCConfig(),
		HTTP:   newDefaultHTTPConfig(),
		Minio:  newDefaultMinioConfig(),
		LogLvl: defaultLogLvl,
	}
}

func (c *Config) ParseEnv() error {
	c.GRPC.parseEnv()

	c.HTTP.parseEnv()

	c.Minio.parseEnv()

	if err := c.parseEnvLogLvl(); err != nil {
		return errors.Wrap(err, "c.parseEnvLogLvl")
	}

	return nil
}

func (c *Config) parseEnvLogLvl() error {
	envLogLvl := os.Getenv(envNameLogLvl)
	if envLogLvl != "" {
		logLevel, err := zerolog.ParseLevel(envLogLvl)
		if err != nil {
			return errors.Wrapf(err, "parse log lvl: %s", envLogLvl)
		}
		c.LogLvl = logLevel
	}
	return nil
}
