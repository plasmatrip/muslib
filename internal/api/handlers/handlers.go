package handlers

import (
	"net/http"
	"time"

	"github.com/plasmatrip/muslib/internal/config"
	"github.com/plasmatrip/muslib/internal/logger"
	"github.com/plasmatrip/muslib/internal/storage"
)

type Handlers struct {
	Config config.Config
	Logger logger.Logger
	Stor   storage.Repository
	Client http.Client
}

func NewHandlers(cfg config.Config, l logger.Logger, db storage.Repository) *Handlers {
	return &Handlers{
		Config: cfg,
		Logger: l,
		Stor:   db,
		Client: http.Client{Timeout: cfg.ClientTimeout * time.Second},
	}
}
