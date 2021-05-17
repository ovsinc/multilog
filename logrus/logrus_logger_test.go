// Package logrus реализует логгер logrus.
package logrus

import (
	"multilog/mock"
	"regexp"
	"testing"

	origlogrus "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

const (
	rx = "time=\"[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}\\+[0-9]{2}:00\" level=(debug|info|warning|error) msg=\"hello world\""
)

func TestNew(t *testing.T) {
	w := mock.NewWW()

	ol := origlogrus.New()
	ol.SetOutput(w)
	ol.SetLevel(origlogrus.DebugLevel)

	l := New(ol)

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
