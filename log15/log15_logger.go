// Package log15 implements the log15 logger.
//
// RU:
// Package log15 реализует логгер log15.
package log15

import (
	"github.com/ovsinc/multilog/common"
	log "github.com/ovsinc/multilog/common"

	log15orig "github.com/inconshreveable/log15"
)

// New constructor of a logger that wraps the original logger.
//
// RU:
// New конструктор log15 логгера.
// Оборачивает log15 логгер l.
func New(l ...log15orig.Logger) log.Logger {
	logger := log15orig.New()
	if len(l) > 0 {
		logger = l[0]
	}
	return &log15logger{
		logger: logger,
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
