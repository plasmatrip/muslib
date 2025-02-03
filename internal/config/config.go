package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const clientTimeout = time.Second * 5 //таймаут запроса к внешнему сервису

type Config struct {
	Host          string        `env:"RUN_ADDRESS"`          //адрес веб-сервера
	Database      string        `env:"DATABASE_URI"`         //DSN базы данных
	InfoService   string        `env:"INFO_SERVICE_ADDRESS"` //адрес внешнего сервиса
	LogLevel      string        `env:"LOG_LEVEL"`            //уровень логирования
	ClientTimeout time.Duration //таймаут запроса к внешнему сервису
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		ClientTimeout: clientTimeout,
	}

	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}

	//пытаемся загрузить .env файл
	if err := godotenv.Load(filepath.Dir(ex) + "/.env"); err != nil {
		return nil, errors.New(".env not found")
	}

	// читаем переменные окружения, при ошибке прокидываем ее наверх
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to read environment variable: %w", err)
	}

	if _, exist := os.LookupEnv("RUN_ADDRESS"); !exist {
		return nil, errors.New("RUN_ADDRESS not found")
	}

	if _, exist := os.LookupEnv("DATABASE_URI"); !exist {
		return nil, errors.New("DATABASE_URI not found")
	}

	if _, exist := os.LookupEnv("INFO_SERVICE_ADDRESS"); !exist {
		return nil, errors.New("INFO_SERVICE_ADDRESS not found")
	}

	if _, exist := os.LookupEnv("LOG_LEVEL"); !exist {
		return nil, errors.New("LOG_LEVEL not found")
	}

	return cfg, nil
}
