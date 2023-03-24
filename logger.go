package configcatpreevaluate

// From: https://github.com/configcat/go-sdk/blob/218279137924bf88947c8cc8120aec6287b33104/logger.go

import (
	"github.com/sirupsen/logrus"
)

// Logger defines the interface this library logs with.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
}

// DefaultLogger creates the default logger (logrus.New()).
func DefaultLogger() Logger {
	return logrus.New()
}
