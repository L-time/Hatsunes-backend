package log

import "github.com/sirupsen/logrus"

func NewSysLogger() *SysLogger {
	s := SysLogger{
		Logger: logrus.New(),
	}
	s.Logger.SetFormatter(&sysFormat)
	return &s
}

func (sys *SysLogger) Infof(s string, i ...interface{}) {
	sys.Logger.Infof(s, i...)
}

func (sys *SysLogger) Errorf(s string, i ...interface{}) {
	sys.Logger.Errorf(s, i...)
}

func (sys *SysLogger) Warnf(s string, i ...interface{}) {
	sys.Logger.Warnf(s, i...)
}

func (sys *SysLogger) Fatalf(s string, i ...interface{}) {
	sys.Logger.Fatalf(s, i...)
}

func (sys *SysLogger) Debugf(s string, i ...interface{}) {
	sys.Logger.Debugf(s, i...)
}
