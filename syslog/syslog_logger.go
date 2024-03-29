//go:build !windows && !plan9
// +build !windows,!plan9

// Package syslog implements the syslog logger.
//
// RU:
// Package syslog реализует логгер syslog.
package syslog

import (
	gosystemsyslog "log/syslog"

	"github.com/ovsinc/multilog/common"
)

// New constructor of a logger that wraps the original syslog writer.
//
// RU:
// New конструтор интерфейс для использования логгера syslog.
// Оборачивает writer интерфейс.
func New(w ...*gosystemsyslog.Writer) common.Logger {
	writer, _ := gosystemsyslog.New(gosystemsyslog.LOG_DEBUG|gosystemsyslog.LOG_LOCAL0, "")
	if len(w) > 0 {
		writer = w[0]
	}
	return &sysloglogger{
		writer: writer,
	}
}

type sysloglogger struct {
	writer *gosystemsyslog.Writer
}

func (l *sysloglogger) Debugf(format string, args ...interface{}) {
	_ = l.writer.Debug(common.Format(format, args...))
}

func (l *sysloglogger) Infof(format string, args ...interface{}) {
	_ = l.writer.Info(common.Format(format, args...))
}

func (l *sysloglogger) Warnf(format string, args ...interface{}) {
	_ = l.writer.Warning(common.Format(format, args...))
}

func (l *sysloglogger) Errorf(format string, args ...interface{}) {
	_ = l.writer.Err(common.Format(format, args...))
}
