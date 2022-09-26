// Package zap implements the zap logger.
//
// RU:
// Package zap реализует логгер zap.
package zap

import (
	"github.com/ovsinc/multilog/common"

	origzap "go.uber.org/zap"
)

// New constructor of a logger that wraps the original logger.
//
// RU:
// New конструтор интерфейс для использования логгера zap
// Оборачивает логгер zap l.
func New(l ...*origzap.Logger) common.Logger {
	logger, _ := origzap.NewProduction()
	if len(l) > 0 {
		logger = l[0]
	}
	return &zaplogger{
		logger: logger,
	}
}

type zaplogger struct {
	logger *origzap.Logger
}

func (l *zaplogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(common.Format(format, args...))
}

func (l *zaplogger) Infof(format string, args ...interface{}) {
	l.logger.Info(common.Format(format, args...))
}

func (l *zaplogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(common.Format(format, args...))
}

func (l *zaplogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(common.Format(format, args...))
}
