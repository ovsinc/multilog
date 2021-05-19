package chain_test

import (
	"bytes"
	"io"
	"os"

	origlogrus "github.com/sirupsen/logrus"

	log15orig "github.com/inconshreveable/log15"

	"github.com/ovsinc/multilog/chain"
	"github.com/ovsinc/multilog/log15"
	"github.com/ovsinc/multilog/logrus"
)

func LogfmtFormat() log15orig.Format {
	return log15orig.FormatFunc(func(r *log15orig.Record) []byte {
		buf := &bytes.Buffer{}

		_, _ = buf.Write([]byte("lvl="))
		_, _ = io.WriteString(buf, r.Lvl.String())

		_, _ = buf.Write([]byte(" msg=\""))
		_, _ = io.WriteString(buf, r.Msg)
		_, _ = buf.Write([]byte("\""))

		return buf.Bytes()
	})
}

func ExampleNew() {
	l1 := origlogrus.New()
	l1.SetOutput(os.Stdout)
	l1.SetLevel(origlogrus.DebugLevel)
	l1.SetFormatter(&origlogrus.TextFormatter{
		DisableTimestamp: true,
	})

	l2 := log15orig.New()
	l2.SetHandler(log15orig.StreamHandler(os.Stdout, LogfmtFormat()))

	chainLog := chain.New(logrus.New(l1), log15.New(l2))
	chainLog.Infof("Hello %s", "world")

	// Output:
	// level=debug msg="Hello world"
	// lvl=info msg="Hello world"
}
