package config

import "os"

const (
	defaultMinioHost           = "127.0.0.1"
	defaultMinioPort           = "9000"
	defaultMinioAccessKey      = "minioadmin"
	defaultMinioSecretKey      = "minioadmin"
	defaultMinioPostBucketName = "post"
)
const (
	envNameMinioHost           = "MINIO_HOST"
	envNameMinioPort           = "MINIO_PORT"
	envNameMinioAccessKey      = "MINIO_ACCESS_KEY"
	envNameMinioSecretKey      = "MINIO_SECRET_KEY"
	envNameMinioPostBucketName = "MINIO_POST_BUCKET_NAME"
)

type Minio struct {
	Host           string
	Port           string
	AccessKey      string
	SecretKey      string
	PostBucketName string
}

func newDefaultMinioConfig() Minio {
	return Minio{
		Host:           defaultMinioHost,
		Port:           defaultMinioPort,
		AccessKey:      defaultMinioAccessKey,
		SecretKey:      defaultMinioSecretKey,
		PostBucketName: defaultMinioPostBucketName,
	}
}

func (c *Minio) parseEnv() {
	envHost := os.Getenv(envNameMinioHost)
	if envHost != "" {
		c.Host = envHost
	}

	envPort := os.Getenv(envNameMinioPort)
	if envPort != "" {
		c.Port = envPort
	}

	envAccessKey := os.Getenv(envNameMinioAccessKey)
	if envAccessKey != "" {
		c.AccessKey = envAccessKey
	}

	envSecretKey := os.Getenv(envNameMinioSecretKey)
	if envSecretKey != "" {
		c.SecretKey = envSecretKey
	}

	envPostBucketName := os.Getenv(envNameMinioPostBucketName)
	if envPostBucketName != "" {
		c.PostBucketName = envPostBucketName
	}
}
