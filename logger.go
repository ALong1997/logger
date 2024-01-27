package logger

import (
	"go.uber.org/zap"
)

func getLogger() *zap.SugaredLogger {
	if withGoID {
		return logger.With(zap.String(goIDKey, GoID()))
	}

	return logger
}

func Sync() error {
	return logger.Sync()
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
