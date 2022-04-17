package config

import (
	"fmt"
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port            int
	DSN             string
	ServiceName     string
	LogLevel        uint32
	Salt            string
	TokenTTL        time.Duration
	RefreshTokenTTL time.Duration
	TokenSecret     string
	RedisURL        string
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
