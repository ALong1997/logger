package logger

import (
	"go.uber.org/zap"
)

func Debug(args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Debug(args)
	} else {
		logger.Debug(args)
	}
}

func Info(args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Info(args)
	} else {
		logger.Info(args)
	}
}

func Warn(args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Warn(args)
	} else {
		logger.Warn(args)
	}
}

func Error(args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Error(args)
	} else {
		logger.Error(args)
	}
}

func DebugF(format string, args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Debugf(format, args)
	} else {
		logger.Debugf(format, args)
	}
}

func InfoF(format string, args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Infof(format, args)
	} else {
		logger.Infof(format, args)
	}
}

func WarnF(format string, args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Warnf(format, args)
	} else {
		logger.Warnf(format, args)
	}
}

func ErrorF(format string, args ...interface{}) {
	if goID {
		logger.With(zap.String("GoID", GoID())).Errorf(format, args)
	} else {
		logger.Errorf(format, args)
	}
}

func Sync() error {
	return logger.Sync()
}
