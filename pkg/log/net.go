package log

import (
	"github.com/sirupsen/logrus"
)

func NewNetLogger() *NetLogger {
	n := &NetLogger{
		Logger: logrus.New(),
	}
	n.Logger.SetFormatter(&netFormat)
	return n
}

func (n *NetLogger) Infof(s string, i ...interface{}) {
	n.Logger.Infof(s, i...)
}

func (n *NetLogger) Errorf(s string, i ...interface{}) {
	n.Logger.Errorf(s, i)
}

func (n *NetLogger) Warnf(s string, i ...interface{}) {
	n.Logger.Warnf(s, i...)
}

func (n *NetLogger) Fatalf(s string, i ...interface{}) {
	n.Logger.Fatalf(s, i...)
}

func (n *NetLogger) Debugf(s string, i ...interface{}) {
	n.Logger.Debugf(s, i...)
}
