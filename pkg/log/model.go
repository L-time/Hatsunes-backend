package log

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var netFormat = nested.Formatter{
	TimestampFormat: "2006-01-02 15:04:05",
	HideKeys:        false,
}

var ormFormat = nested.Formatter{
	TimestampFormat: "2006-01-02 15:04:05",
	HideKeys:        false,
}

var sysFormat = nested.Formatter{
	TimestampFormat: "2006-01-02 15:04:05",
	HideKeys:        false,
}

type Handel interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
	Fatalf(string, ...interface{})
	Debugf(string, ...interface{})
}

type NetLogger struct {
	Logger *logrus.Logger
}

type SysLogger struct {
	Logger *logrus.Logger
}
