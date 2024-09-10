package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

type OrmLogger struct {
	Logger *logrus.Logger
}

func NewOrmLogger() *OrmLogger {
	o := &OrmLogger{
		Logger: logrus.New(),
	}
	o.Logger.SetFormatter(&ormFormat)
	return o
}

func (o *OrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *o
	switch level {
	case logger.Silent:
		newLogger.Logger.SetLevel(logrus.PanicLevel)
	case logger.Error:
		newLogger.Logger.SetLevel(logrus.ErrorLevel)
	case logger.Warn:
		newLogger.Logger.SetLevel(logrus.WarnLevel)
	case logger.Info:
		newLogger.Logger.SetLevel(logrus.InfoLevel)
	default:
		newLogger.Logger.SetLevel(logrus.DebugLevel)
	}
	return &newLogger
}

func (o *OrmLogger) Info(ctx context.Context, s string, i ...interface{}) {
	o.Logger.WithContext(ctx).Infof(s, i...)
}

func (o *OrmLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	o.Logger.WithContext(ctx).Warnf(s, i...)
}

func (o *OrmLogger) Error(ctx context.Context, s string, i ...interface{}) {
	o.Logger.WithContext(ctx).Errorf(s, i...)
}

func (o *OrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	entry := o.Logger.WithContext(ctx).WithFields(logrus.Fields{
		"duration":      elapsed,
		"rows_affected": rows,
		"sql":           sql,
	})

	if err != nil {
		entry.WithField("error", err).Error("trace error")
	} else {
		entry.Info("trace success")
	}
}
