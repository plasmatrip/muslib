package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/plasmatrip/muslib/internal/api/handlers"
	"github.com/plasmatrip/muslib/internal/api/middleware"
	"github.com/plasmatrip/muslib/internal/config"
	"github.com/plasmatrip/muslib/internal/logger"
	"github.com/plasmatrip/muslib/internal/storage"
)

func NewRouter(cfg config.Config, log logger.Logger, stor storage.Repository) *chi.Mux {

	r := chi.NewRouter()

	handlers := handlers.Handlers{Config: cfg, Logger: log, Stor: stor}

	r.Use(middleware.WithLogging(log), middleware.WithCompression)

	r.Route("/info", func(r chi.Router) {
		r.Get("/", handlers.Info)
	})

	r.Route("/song", func(r chi.Router) {
		r.Post("/", handlers.AddSong)
		r.Put("/", handlers.UpdateSong)
		r.Delete("/", handlers.DeleteSong)
	})

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", handlers.GetSongs)
	})

	r.Route("/lyrics", func(r chi.Router) {
		r.Get("/", handlers.GetLyrics)
	})

	return r
}
