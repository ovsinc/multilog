package multilog_test

import (
	"os"

	"github.com/ovsinc/multilog"
	"github.com/ovsinc/multilog/logrus"
	origlogrus "github.com/sirupsen/logrus"
)

func ExampleDefaultLogger() {
	l := origlogrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(origlogrus.DebugLevel)
	l.SetFormatter(&origlogrus.TextFormatter{
		DisableTimestamp: true,
	})

	multilog.DefaultLogger = logrus.New(l)
	multilog.Infof("Hello %s", "world")

	// Output:
	// level=debug msg="Hello world"
}
