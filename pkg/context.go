package gotime

import "context"

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	// gotime log args
	ContextKeyLogType = contextKey("logType")
	// gotime config
	ContextKeyGoTimeDir = contextKey("gotimeDir")
)

func GetGoTimeDir(ctx context.Context) string {
	return ctx.Value(ContextKeyGoTimeDir).(string)
}

func GetLogType(ctx context.Context) LogType {
	return ctx.Value(ContextKeyLogType).(LogType)
}
