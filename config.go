package logger

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	// reference lumberjack.Logger
	LogLevel   zapcore.Level
	FilePath   string
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool

	// print to console
	Console bool

	// print GoroutineID
	GoroutineID bool

	once sync.Once
}

func DefaultConfig() *Config {
	return &Config{
		LogLevel:   zapcore.DebugLevel,
		FilePath:   "/opt/log",
		FileName:   "app.log",
		MaxSize:    100,
		MaxAge:     10,
		MaxBackups: 10,
		Compress:   false,
		Console:    false,
	}
}

// check use default field to replace invalid field
func (c *Config) check() {
	if c == nil {
		panic("invalid config")
	}

	defaultConfig := DefaultConfig()

	if c.LogLevel < zapcore.DebugLevel || c.LogLevel >= zapcore.InvalidLevel {
		c.LogLevel = defaultConfig.LogLevel
	}

	if len(c.FilePath) == 0 {
		c.FilePath = defaultConfig.FilePath
	}

	if len(c.FileName) == 0 {
		c.FileName = defaultConfig.FileName
	}

	if c.MaxSize <= 0 {
		c.MaxSize = defaultConfig.MaxSize
	}

	if c.MaxAge <= 0 {
		c.MaxAge = defaultConfig.MaxAge
	}

	if c.MaxBackups <= 0 {
		c.MaxBackups = defaultConfig.MaxBackups
	}
}
