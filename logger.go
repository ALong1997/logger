package logger

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func DebugF(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func InfoF(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func WarnF(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func ErrorF(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Sync() error {
	return logger.Sync()
}
