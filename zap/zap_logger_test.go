// Package zap реализует логгер zap.
package zap

import (
	"regexp"
	"testing"

	"github.com/ovsinc/multilog/mock"

	"github.com/stretchr/testify/assert"

	origzap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const rx = "(debug|info|error|warn)\thello world"

func TestNew(t *testing.T) {
	w := mock.NewWW()

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), w, origzap.DebugLevel)

	l := New(origzap.New(core))

	re := regexp.MustCompile(rx)

	l.Debugf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Infof(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Warnf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Errorf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))
}
