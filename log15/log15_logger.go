// Package log15 hеализует логгер log15.
package log15

import (
	"multilog/common"
	log "multilog/common"

	log15orig "github.com/inconshreveable/log15"
)

// New конструктор log15 логгера.
// Оборачивает log15 логгер l.
func New(l log15orig.Logger) log.Logger {
	return &log15logger{
		logger: l,
	}
}

type log15logger struct {
	logger log15orig.Logger
}

func (l *log15logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(common.Format(format, args...))
}

func (l *log15logger) Infof(format string, args ...interface{}) {
	l.logger.Info(common.Format(format, args...))
}

func (l *log15logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(common.Format(format, args...))
}

func (l *log15logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(common.Format(format, args...))
}
