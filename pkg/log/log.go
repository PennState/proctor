package log

import (
	"github.com/sirupsen/logrus"
)

// Logger returns a new logrus Logger that outputs its data to the
// testing package's Log() method.
func Logger(t TestingT) *logrus.Logger {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	l, err := logrus.ParseLevel(*level)
	if err != nil {
		t.Log(err)
	}
	logger := logrus.New()
	logger.SetLevel(l)
	logger.SetOutput(logWriter{t: t})
	return logger
}

type logWriter struct {
	t TestingT
}

// Write implements the io.Writer interface to allow the testing
// package's logger to receive
func (w logWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
