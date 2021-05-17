package golog

import (
	gosystemlog "log"
	"multilog/mock"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	prefix = "test/log "
	rx     = prefix + "[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2} (DEBUG|INFO|ERR|WARN): hello world"
)

func TestNew(t *testing.T) {
	w := mock.NewWW()
	l := New(gosystemlog.New(w, prefix, gosystemlog.LstdFlags))

	re := regexp.MustCompile(rx)

	l.Debugf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Infof(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Errorf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))

	l.Warnf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re, string(w.Read()))
}