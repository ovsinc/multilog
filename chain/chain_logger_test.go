// Package chain реализаует цепочку логгирования.
package chain

import (
	gosystemlog "log"
	"regexp"
	"testing"

	golog "github.com/ovsinc/multilog/golog"
	"github.com/ovsinc/multilog/logrus"
	"github.com/ovsinc/multilog/mock"

	origlogrus "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

var (
	prefix = "test/log "
	rx2    = prefix + "[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2} (DEBUG|INFO|ERR|WARN): hello world"
	rx1    = "time=\"[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}\\+[0-9]{2}:00\" level=(debug|info|warning|error) msg=\"hello world\""
)

func TestNew(t *testing.T) {
	w1 := mock.NewWW()
	ol1 := origlogrus.New()
	ol1.SetOutput(w1)
	ol1.SetLevel(origlogrus.DebugLevel)
	l1 := logrus.New(ol1)

	w2 := mock.NewWW()
	l2 := golog.New(gosystemlog.New(w2, prefix, gosystemlog.LstdFlags))

	l := New(l1, l2)

	re1 := regexp.MustCompile(rx1)
	re2 := regexp.MustCompile(rx2)

	l.Debugf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re1, string(w1.Read()))
	assert.Regexp(t, re2, string(w2.Read()))

	l.Infof(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re1, string(w1.Read()))
	assert.Regexp(t, re2, string(w2.Read()))

	l.Warnf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re1, string(w1.Read()))
	assert.Regexp(t, re2, string(w2.Read()))

	l.Errorf(mock.MsgFmt, mock.Msg)
	assert.Regexp(t, re1, string(w1.Read()))
	assert.Regexp(t, re2, string(w2.Read()))
}
