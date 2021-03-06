package config

import (
	"fmt"
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port               int
	AppHost            string
	DSN                string
	ServiceName        string
	LogLevel           uint32
	Salt               string
	TokenTTL           time.Duration
	RefreshTokenTTL    time.Duration
	TokenSecret        string
	LogFile            string        `envconfig:"default=./vars/logs/main.log"`
	FileStore          string        `envconfig:"default=./file_store"`
	RequestMaxDuration time.Duration `envconfig:"default=30s"`
	FtpHost            string
	FtpLogin           string
	FtpPass            string
	FtpPath            string
	BitrixHook         string
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
