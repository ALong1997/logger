package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
)

var (
	logger *zap.SugaredLogger

	goID   bool
	prefix string // use getPrefix() to read it
)

func Init(c *Config) {
	c.once.Do(func() {
		c.check()
		core := zapcore.NewCore(getEncoder(), getLogWriter(c), c.LogLevel)
		logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

		goID = c.GoroutineID
	})
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(c *Config) zapcore.WriteSyncer {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(c.FilePath, c.FileName),
		MaxAge:     c.MaxAge,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		Compress:   c.Compress,
	})
	if c.Console {
		return zapcore.NewMultiWriteSyncer(fileWriter, zapcore.Lock(os.Stdout))
	}
	return fileWriter
}

func setPrefix(s string) {
	prefix = s
}

func getPrefix() string {
	if goID {
		return strings.TrimSpace("GoID:" + GoID() + " " + prefix)
	} else {
		return prefix
	}
}
