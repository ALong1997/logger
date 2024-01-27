package logger

import (
	"go.uber.org/zap"
)

func getLogger() *zap.SugaredLogger {
	if globalLogger.withGoID {
		return globalLogger.With(zap.String(goIDKey, GoID()))
	}

	return globalLogger.SugaredLogger
}

func Sync() error {
	return globalLogger.Sync()
}

func Debug(args ...interface{}) {
	getLogger().Debug(args)
}

func Info(args ...interface{}) {
	getLogger().Info(args)
}

func Warn(args ...interface{}) {
	getLogger().Warn(args)
}

func Error(args ...interface{}) {
	getLogger().Error(args)
}

func DebugF(format string, args ...interface{}) {
	getLogger().Debugf(format, args)
}

func InfoF(format string, args ...interface{}) {
	getLogger().Infof(format, args)
}

func WarnF(format string, args ...interface{}) {
	getLogger().Warnf(format, args)
}

func ErrorF(format string, args ...interface{}) {
	getLogger().Errorf(format, args)
}
