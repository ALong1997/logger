package logger

import (
	"fmt"
)

func Debug(args ...interface{}) {
	logger.Debug(getPrefix(), args)
}

func Info(args ...interface{}) {
	logger.Info(getPrefix(), args)
}

func Warn(args ...interface{}) {
	logger.Warn(getPrefix(), args)
}

func Error(args ...interface{}) {
	logger.Error(getPrefix(), args)
}

func DebugF(format string, args ...interface{}) {
	logger.Debugf(getPrefix(), fmt.Sprintf(format, args...))
}

func InfoF(format string, args ...interface{}) {
	logger.Infof(getPrefix(), fmt.Sprintf(format, args...))
}

func WarnF(format string, args ...interface{}) {
	logger.Warnf(getPrefix(), fmt.Sprintf(format, args...))
}

func ErrorF(format string, args ...interface{}) {
	logger.Errorf(getPrefix(), fmt.Sprintf(format, args...))
}

func Sync() error {
	return logger.Sync()
}
