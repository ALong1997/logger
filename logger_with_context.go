package logger

import (
	"context"
	"go.uber.org/zap"
)

const TraceIDKey = "trace_id"

func GetLoggerWithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx != nil {
		if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
			return GetLogger().With(zap.String(TraceIDKey, traceID))
		}
	}

	return GetLogger()
}

func DebugWithContext(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Debug(args)
}

func InfoWithContext(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Info(args)
}

func WarnWithContext(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Warn(args)
}

func ErrorWithContext(ctx context.Context, args ...interface{}) {
	GetLoggerWithContext(ctx).Error(args)
}

func DebugFWithContext(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Debugf(format, args)
}

func InfoFWithContext(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Infof(format, args)
}

func WarnFWithContext(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Warnf(format, args)
}

func ErrorFWithContext(ctx context.Context, format string, args ...interface{}) {
	GetLoggerWithContext(ctx).Errorf(format, args)
}
