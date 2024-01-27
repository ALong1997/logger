package logger

import (
	"context"
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

	traceID := "123456"
	ctx := context.WithValue(context.Background(), TraceIDKey, traceID)
	DebugWithContext(ctx, s)
	InfoWithContext(ctx, s)
	WarnWithContext(ctx, s)
	ErrorWithContext(ctx, s)
	DebugFWithContext(ctx, "%v", s)
	InfoFWithContext(ctx, "%v", s)
	WarnFWithContext(ctx, "%v", s)
	ErrorFWithContext(ctx, "%v", s)

	if err := Sync(); err != nil {
		Error(err)
	}
}
