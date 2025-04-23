package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenHttpPort string `envconfig:"LISTEN_HTTP_PORT" default:"8080"`
	PostgresAddr   string `envconfig:"POSTGRES_ADDR" default:""`
	JWTSecret      string `envconfig:"JWT_SECRET" default:"secret"`
	ListenGRPCPort int32  `envconfig:"LISTEN_GRPC_PORT" default:"8081"`
}

func New() *Config {
	return &Config{}
}

func (r *Config) Init() {
	if err := envconfig.Process("", r); err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}
}
