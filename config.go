package logger

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		JsonEncoder bool // json or console

		LogLevel zapcore.Level // log level

		Console bool // output to console

		GoroutineID bool // output GoroutineID

		FileConfig FileConfig // log file

		once sync.Once
	}

	// FileConfig reference lumberjack.Logger
	FileConfig struct {
		FilePath   string
		FileName   string
		MaxSize    int
		MaxAge     int
		MaxBackups int
		Compress   bool
	}
)

func DefaultConfig() *Config {
	return &Config{
		JsonEncoder: false,
		LogLevel:    zapcore.DebugLevel,
		Console:     false,
		GoroutineID: false,
		FileConfig:  DefaultFileConfig(),
		once:        sync.Once{},
	}
}

func DefaultFileConfig() FileConfig {
	return FileConfig{
		FilePath:   "/opt/log",
		FileName:   "app.log",
		MaxSize:    100,
		MaxAge:     10,
		MaxBackups: 10,
		Compress:   false,
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

	c.FileConfig.check()
}

// check use default field to replace invalid field
func (fc *FileConfig) check() {
	if fc == nil {
		panic("invalid config")
	}

	defaultConfig := DefaultFileConfig()

	if len(fc.FilePath) == 0 {
		fc.FilePath = defaultConfig.FilePath
	}

	if len(fc.FileName) == 0 {
		fc.FileName = defaultConfig.FileName
	}

	if fc.MaxSize <= 0 {
		fc.MaxSize = defaultConfig.MaxSize
	}

	if fc.MaxAge <= 0 {
		fc.MaxAge = defaultConfig.MaxAge
	}

	if fc.MaxBackups <= 0 {
		fc.MaxBackups = defaultConfig.MaxBackups
	}
}
