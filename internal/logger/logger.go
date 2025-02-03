package logger

import (
	"go.uber.org/zap"
)

const (
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
)

type Logger struct {
	zap   *zap.Logger
	Sugar *zap.SugaredLogger
}

func NewLogger(level string) (*Logger, error) {
	var config zap.Config

	config = zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	if level == LogLevelDebug {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	zap, err := config.Build()
	return &Logger{zap: zap, Sugar: zap.Sugar()}, err
}

func (l *Logger) Close() {
	l.zap.Sync()
}
