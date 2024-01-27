package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"sync"
)

type logger struct {
	*zap.SugaredLogger

	withGoID bool

	once sync.Once
}

var globalLogger *logger

func init() {
	globalLogger = &logger{}
}

func Init(c *Config) {
	c.check()
	globalLogger.once.Do(func() {
		core := zapcore.NewCore(getEncoder(c.JsonEncoder), getLogWriter(c.FileConfig, c.Console), c.LogLevel)
		globalLogger.SugaredLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
		globalLogger.withGoID = c.GoroutineID
	})
}

func getEncoder(jsonEncoder bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if jsonEncoder {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(fc FileConfig, console bool) zapcore.WriteSyncer {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(fc.FilePath, fc.FileName),
		MaxAge:     fc.MaxAge,
		MaxSize:    fc.MaxSize,
		MaxBackups: fc.MaxBackups,
		Compress:   fc.Compress,
	})
	if console {
		return zapcore.NewMultiWriteSyncer(fileWriter, zapcore.Lock(os.Stdout))
	}
	return fileWriter
}
