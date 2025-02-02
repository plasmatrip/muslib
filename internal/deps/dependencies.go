package deps

import (
	"github.com/plasmatrip/muslib/internal/config"
	"github.com/plasmatrip/muslib/internal/logger"
	"github.com/plasmatrip/muslib/internal/storage"
)

type Dependencies struct {
	Config config.Config
	Logger logger.Logger
	Stor   storage.Repository
}
