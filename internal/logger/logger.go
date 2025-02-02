package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zap   *zap.Logger
	Sugar *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	zap, err := zap.NewDevelopment()
	return &Logger{zap: zap, Sugar: zap.Sugar()}, err
}

func (l *Logger) Close() {
	l.zap.Sync()
}
