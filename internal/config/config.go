package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	host          = "http://localhost:8080"
	infoService   = "http://localhost:8081"
	database      = "postgres://gratify:password@localhost:5432/gratify?sslmode=disable"
	workers       = 5
	workBuffer    = 5
	clientTimeout = time.Second * 5
)

type Config struct {
	Host              string `env:"RUN_ADDRESS"`
	Database          string `env:"DATABASE_URI"`
	InfoService       string `env:"INFO_SERVICE_ADDRESS"`
	ClientTimeout     time.Duration
	Workers           int
	WorkBuffer        int
	ProcessorInterval int
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		ClientTimeout: clientTimeout,
		Workers:       workers,
		WorkBuffer:    workBuffer,
	}

	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}

	if err := godotenv.Load(filepath.Dir(ex) + "/.env"); err != nil {
		return nil, errors.New(".env not found")
	}

	cl := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// читаем переменные окружения, при ошибке прокидываем ее наверх
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to read environment variable: %w", err)
	}

	var fHost string
	cl.StringVar(&fHost, "a", host, "server address host:port")

	var fDatabase string
	cl.StringVar(&fDatabase, "d", database, "database DSN")

	var fInfoService string
	cl.StringVar(&fInfoService, "s", infoService, "music info service address host:port")

	if err := cl.Parse(os.Args[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	if _, exist := os.LookupEnv("RUN_ADDRESS"); !exist {
		cfg.Host = fHost
	}

	if _, exist := os.LookupEnv("DATABASE_URI"); !exist {
		cfg.Database = fDatabase
	}

	if _, exist := os.LookupEnv("INFO_SERVICE_ADDRESS"); !exist {
		cfg.InfoService = fInfoService
	}

	// if err := parseAddress(cfg); err != nil {
	// 	return nil, fmt.Errorf("port parsing error: %w", err)
	// }

	return cfg, nil
}

// func parseAddress(cfg *Config) error {
// 	var parts []string
// 	_, addr, found := strings.Cut(cfg.Host, "://")
// 	if found {
// 		parts = strings.Split(addr, ":")
// 	} else {
// 		parts = strings.Split(cfg.Host, ":")
// 	}

// 	if len(parts) == 2 {
// 		if len(parts[0]) == 0 || len(parts[1]) == 0 {
// 			cfg.Host = host + ":" + port
// 			return nil
// 		}

// 		_, err := strconv.ParseInt(parts[1], 10, 64)
// 		return err
// 	}
// 	cfg.Host = host + ":" + port
// 	return nil
// }
