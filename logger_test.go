package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	config := DefaultConfig()
	Init(config)

	s := "Hello World!"
	Debug(s)
	Info(s)
	Warn(s)
	Error(s)
	DebugF("%v", s)
	InfoF("%v", s)
	WarnF("%v", s)
	ErrorF("%v", s)

	if err := Sync(); err != nil {
		Error(err)
	}
}
