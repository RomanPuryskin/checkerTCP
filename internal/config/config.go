package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ServerAddress string        `env:"SERVER_ADDRESS" envDefault:"localhost"`
	ServerPort    string        `env:"SERVER_PORT" envDefault:"8080"`
	ReadTimeout   time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
	IdleTimeout   time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
}

type LoggerConfig struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

type CLIConfig struct {
	ServerConnectTimeout string `env:"CLI_SERVER_CONNECT_TIMEOUT" envDefault:"30s"`
	TCPConnectTimeout    string `env:"CLI_TCP_CONNECT_TIMEOUT" envDefault:"30s"`
}

func load(cfg any) error {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}

	return nil
}

func LoadSeverCofig() (*ServerConfig, error) {
	sc := &ServerConfig{}
	if err := load(sc); err != nil {
		return nil, fmt.Errorf("[LoadSeverCofig]: %w", err)
	}
	return sc, nil
}

func LoadLoggerCofig() (*LoggerConfig, error) {
	lc := &LoggerConfig{}
	if err := load(lc); err != nil {
		return nil, fmt.Errorf("[LoadLoggerCofig]: %w", err)
	}
	return lc, nil
}

func LoadCLICofig() (*CLIConfig, error) {
	cc := CLIConfig{}
	if err := load(&cc); err != nil {
		return nil, fmt.Errorf("[LoadCLICofig]: %w", err)
	}
	return &cc, nil
}
