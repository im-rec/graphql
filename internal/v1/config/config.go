package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env"
)

type configurations struct {
	AppName                   string `env:"APP_NAME"`
	AppDescription            string `env:"APP_DESCRIPTION"`
	AppDebug                  bool   `env:"APP_DEBUG"`
	ServerRestPort            int    `env:"SERVER_REST_PORT"`
	ServerRestReadTimeout     int    `env:"SERVER_REST_RTO"`
	ServerRestWriteTimeout    int    `env:"SERVER_REST_WTO"`
	ServerRestIdleTimeout     int    `env:"SERVER_REST_ITO"`
	ServerGraphqlPort         int    `env:"SERVER_GRAPHQL_PORT"`
	ServerGraphqlReadTimeout  int    `env:"SERVER_GRAPHQL_RTO"`
	ServerGraphqlWriteTimeout int    `env:"SERVER_GRAPHQL_WTO"`
	ServerGraphqlIdleTimeout  int    `env:"SERVER_GRAPHQL_ITO"`
}

var (
	configuration configurations
	mutex         sync.Once
)

func GetConfig() configurations {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() configurations {
	var cfg = configurations{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("got an error while parsing config, error: %s", err)
	}

	return cfg
}
