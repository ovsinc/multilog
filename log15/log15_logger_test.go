// Package log15 hеализует логгер log15.
package log15

import (
	"regexp"
	"testing"

	"github.com/ovsinc/multilog/mock"

	log15orig "github.com/inconshreveable/log15"
	"github.com/stretchr/testify/assert"
)

const (
	rx = "t=[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}\\+[0-9]{2}00 lvl=(dbug|info|warn|eror) msg=\"hello world\""
)

func TestNew(t *testing.T) {
	w := mock.NewWW()

	re := regexp.MustCompile(rx)

	ol := log15orig.New()
	ol.SetHandler(log15orig.StreamHandler(w, log15orig.LogfmtFormat()))

	l := New(ol)

	l.Debugf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Infof(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Warnf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Errorf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))
}
