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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	defer log.Close()

	db, err := storage.NewRepository(ctx, cfg.Database, *log)
	if err != nil {
		log.Sugar.Infow("database connection error: ", err)
		os.Exit(1)
	}
	defer db.Close()

	// ctrl := controller.NewController(cfg.ClientTimeout, *deps)
	// ctrl.StartWorkers(ctx)
	// ctrl.StartOrdersProcessor(ctx)

	server := http.Server{
		Addr: cfg.Host,
		Handler: func(next http.Handler) http.Handler {
			log.Sugar.Infow("The Music Library server is running. ", "Server address", cfg.Host, "Music info service address", cfg.InfoService)
			return next
		}(router.NewRouter(*cfg, *log, *db)),
	}

	go server.ListenAndServe()

	<-ctx.Done()

	server.Shutdown(context.Background())

	log.Sugar.Infow("The server has been shut down gracefully")

	os.Exit(0)
}
