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
	logger.Debug(getPrefix(), fmt.Sprintf(format, args...))
}

func InfoF(format string, args ...interface{}) {
	logger.Info(getPrefix(), fmt.Sprintf(format, args...))
}

func WarnF(format string, args ...interface{}) {
	logger.Warn(getPrefix(), fmt.Sprintf(format, args...))
}

func ErrorF(format string, args ...interface{}) {
	logger.Error(getPrefix(), fmt.Sprintf(format, args...))
}

func Sync() error {
	return logger.Sync()
}
