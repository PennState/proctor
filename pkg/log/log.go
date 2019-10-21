package log

import (
	"github.com/sirupsen/logrus"
)

type logWriter struct {
	t TestingT
}

// Write implements the io.Writer interface to allow the testing
// package's logger to receive
func (w logWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}

// Logger wraps the logrus.Logger methods as test helpers
type Logger struct {
	logger *logrus.Logger
	Debug  log  // Debug logs the provided arguments at debug level
	Debugf logf // Debugf logs the format produced string at debug level
	Error  log
	Errorf logf
	Fatal  log
	Fatalf logf
	Info   log
	Infof  logf
	Panic  log
	Panicf logf
	Trace  log
	Tracef logf
	Warn   log
	Warnf  logf
}

// Logger returns a new logrus Logger that outputs its data to the
// testing package's Log() method.
func New(t TestingT) *Logger {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	logger := logrus.New()
	return Wrap(t, logger)
}

// Wrap encapsulates an existing logrus.Logger inside methods that
// are known to be test helpers
func Wrap(t TestingT, logger *logrus.Logger) *Logger {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	l, err := logrus.ParseLevel(*level)
	if err != nil {
		t.Log(err)
	}
	logger.SetLevel(l)
	logger.SetOutput(logWriter{t: t})
	return &Logger{
		logger: logger,
		Debug:  wrap(t, logger, logrus.DebugLevel),
		Debugf: wrapf(t, logger, logrus.DebugLevel),
		Error:  wrap(t, logger, logrus.ErrorLevel),
		Errorf: wrapf(t, logger, logrus.ErrorLevel),
		Fatal:  wrap(t, logger, logrus.FatalLevel),
		Fatalf: wrapf(t, logger, logrus.FatalLevel),
		Info:   wrap(t, logger, logrus.InfoLevel),
		Infof:  wrapf(t, logger, logrus.InfoLevel),
		Panic:  wrap(t, logger, logrus.PanicLevel),
		Panicf: wrapf(t, logger, logrus.PanicLevel),
		Trace:  wrap(t, logger, logrus.TraceLevel),
		Tracef: wrapf(t, logger, logrus.TraceLevel),
		Warn:   wrap(t, logger, logrus.WarnLevel),
		Warnf:  wrapf(t, logger, logrus.WarnLevel),
	}
}

type log func(...interface{})
type logf func(format string, args ...interface{})

func wrap(t TestingT, logger *logrus.Logger, level logrus.Level) log {
	return func(args ...interface{}) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}

		logger.Log(level, args)
	}
}

func wrapf(t TestingT, logger *logrus.Logger, level logrus.Level) logf {
	return func(format string, args ...interface{}) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}

		logger.Logf(level, format, args)
	}
}

// var (
// 	Debug reporter = f
// )

// func Debug(args interface{}) {
// 	if h, ok := t.(tHelper); ok {
// 		h.Helper()
// 	}

// }
