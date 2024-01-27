package logger

import (
	"context"
	"go.uber.org/zap"
)

const TraceIDKey = "trace_id"

func getLoggerWithContext(ctx context.Context) *zap.SugaredLogger {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if ok {
		return getLogger().With(zap.String(TraceIDKey, traceID))
	}

	return getLogger()
}

func DebugWithContext(ctx context.Context, args ...interface{}) {
	getLoggerWithContext(ctx).Debug(args)
}

func InfoWithContext(ctx context.Context, args ...interface{}) {
	getLoggerWithContext(ctx).Info(args)
}

func WarnWithContext(ctx context.Context, args ...interface{}) {
	getLoggerWithContext(ctx).Warn(args)
}

func ErrorWithContext(ctx context.Context, args ...interface{}) {
	getLoggerWithContext(ctx).Error(args)
}

func DebugFWithContext(ctx context.Context, format string, args ...interface{}) {
	getLoggerWithContext(ctx).Debugf(format, args)
}

func InfoFWithContext(ctx context.Context, format string, args ...interface{}) {
	getLoggerWithContext(ctx).Infof(format, args)
}

func WarnFWithContext(ctx context.Context, format string, args ...interface{}) {
	getLoggerWithContext(ctx).Warnf(format, args)
}

func ErrorFWithContext(ctx context.Context, format string, args ...interface{}) {
	getLoggerWithContext(ctx).Errorf(format, args)
}
