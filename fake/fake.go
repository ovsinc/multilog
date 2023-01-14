package fake

import (
	"github.com/ovsinc/multilog/common"
)

type fakeLogger struct{}

func NewfakeLogger() common.Logger {
	return fakeLogger{}
}

func (fakeLogger) Debugf(format string, args ...interface{}) {}
func (fakeLogger) Infof(format string, args ...interface{})  {}
func (fakeLogger) Warnf(format string, args ...interface{})  {}
func (fakeLogger) Errorf(format string, args ...interface{}) {}
func (fakeLogger) Fatalf(format string, args ...interface{}) {}
