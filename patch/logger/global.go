package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	_defaultLogger *Logger
	_levelLogger   *levelLogger
	m              sync.RWMutex
)

func init() {
	_defaultLogger = NewLogger(os.Stdout, WithSource(),
		WithTime(), WithLevel(slog.LevelDebug))
}

func setGlobalLevelLogger(ll *levelLogger) {
	m.Lock()
	defer m.Unlock()
	_levelLogger = ll
}

func getLogger(lv slog.Level) *Logger {
	m.RLock()
	defer m.RUnlock()
	if _levelLogger == nil {
		return _defaultLogger
	}
	logger := _levelLogger.getLogger(lv)
	if logger == nil {
		return _defaultLogger
	}
	return logger
}

func GetOutput(lv slog.Level) io.Writer {
	m.RLock()
	defer m.RUnlock()
	if _levelLogger == nil {
		return os.Stdout
	}
	return _levelLogger.getOutput(lv)
}

func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.Enabled(ctx, level) {
		return
	}
	var pc uintptr

	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.Handler().Handle(ctx, r)
}

func Debug(msg string, args ...any) {
	getLogger(slog.LevelDebug).log(context.Background(), slog.LevelDebug, msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	getLogger(slog.LevelDebug).log(ctx, slog.LevelDebug, msg, args...)
}

func Info(msg string, args ...any) {
	getLogger(slog.LevelInfo).log(context.Background(), slog.LevelInfo, msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	getLogger(slog.LevelInfo).log(ctx, slog.LevelInfo, msg, args...)
}

func Warn(msg string, args ...any) {
	getLogger(slog.LevelWarn).log(context.Background(), slog.LevelWarn, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	getLogger(slog.LevelWarn).log(ctx, slog.LevelWarn, msg, args...)
}

func Error(msg string, args ...any) {
	getLogger(slog.LevelError).log(context.Background(), slog.LevelError, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	getLogger(slog.LevelError).log(ctx, slog.LevelError, msg, args...)
}

func Panic(msg string, args ...any) {
	getLogger(slog.LevelError).log(context.Background(), slog.LevelError, msg, args...)
	panic(msg)
}

func PanicContext(ctx context.Context, msg string, args ...any) {
	getLogger(slog.LevelError).log(ctx, slog.LevelError, msg, args...)
	panic(msg)
}
