// +build !windows,!plan9

// Package syslog реализует логгер syslog.
package syslog

import (
	gosystemsyslog "log/syslog"
	"multilog/common"
)

// New конструтор интерфейс для использования логгера syslog.
// Оборачивает writer интерфейс.
func New(w *gosystemsyslog.Writer) common.Logger {
	return &sysloglogger{
		writer: w,
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
