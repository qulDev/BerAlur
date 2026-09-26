package config

import (
	"errors"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DatabaseURL string
	HTTPAddr    string
	Environment string
}

func Load() (Config, error) {
	cfg := Config{DatabaseURL: os.Getenv("DATABASE_URL"), HTTPAddr: os.Getenv("HTTP_ADDR"), Environment: os.Getenv("APP_ENV")}
	if cfg.HTTPAddr == "" {
		cfg.HTTPAddr = ":8080"
	}
	if cfg.Environment == "" {
		cfg.Environment = "local"
	}
	if cfg.Environment != "local" && cfg.Environment != "dev" && cfg.Environment != "prod" {
		return Config{}, errors.New("APP_ENV must be local, dev, or prod")
	}
	u, err := url.Parse(cfg.DatabaseURL)
	if err != nil || u.Hostname() == "" || strings.Trim(u.Path, "/") == "" || (u.Scheme != "postgres" && u.Scheme != "postgresql") {
		return Config{}, errors.New("DATABASE_URL must be a PostgreSQL URL with host and database")
	}
	_, port, err := net.SplitHostPort(cfg.HTTPAddr)
	number, portErr := strconv.Atoi(port)
	if err != nil || portErr != nil || number < 1 || number > 65535 {
		return Config{}, errors.New("HTTP_ADDR must contain a host and port between 1 and 65535")
	}
	return cfg, nil
}
