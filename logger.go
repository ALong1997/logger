package logger

import (
	"go.uber.org/zap"
)

func GetLogger() *zap.SugaredLogger {
	if globalLogger.withGoID {
		return globalLogger.With(zap.String(goIDKey, GoID()))
	}

	return globalLogger.SugaredLogger
}

func Sync() error {
	return globalLogger.Sync()
}

func Debug(args ...interface{}) {
	GetLogger().Debug(args)
}

func Info(args ...interface{}) {
	GetLogger().Info(args)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args)
}

func Error(args ...interface{}) {
	GetLogger().Error(args)
}

func DebugF(format string, args ...interface{}) {
	GetLogger().Debugf(format, args)
}

func InfoF(format string, args ...interface{}) {
	GetLogger().Infof(format, args)
}

func WarnF(format string, args ...interface{}) {
	GetLogger().Warnf(format, args)
}

func ErrorF(format string, args ...interface{}) {
	GetLogger().Errorf(format, args)
}
