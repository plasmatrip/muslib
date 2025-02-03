package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/plasmatrip/muslib/internal/config"
	"github.com/plasmatrip/muslib/internal/logger"
	"github.com/plasmatrip/muslib/internal/router"
	"github.com/plasmatrip/muslib/internal/storage"
)

func main() {
	// для грейсфул шатдауна слушаем сигнал ОС
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// загружаем конфиг
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// инициализируем логгер
	log, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	defer log.Close()

	// инициализируем БД
	db, err := storage.NewRepository(ctx, cfg.Database, *log)
	if err != nil {
		log.Sugar.Infow("database connection error: ", err)
		os.Exit(1)
	}
	defer db.Close()

	// запускаем веб-сервер
	server := http.Server{
		Addr: cfg.Host,
		Handler: func(next http.Handler) http.Handler {
			log.Sugar.Infow("The Music Library server is running. ", "Server address", cfg.Host, "Music info service address", cfg.InfoService)
			return next
		}(router.NewRouter(*cfg, *log, *db)),
	}

	go server.ListenAndServe()

	// ждем сигнал ОС
	<-ctx.Done()

	server.Shutdown(context.Background())

	log.Sugar.Infow("The server has been shut down gracefully")

	os.Exit(0)
}
