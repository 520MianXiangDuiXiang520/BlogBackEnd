package logger

import (
	"log/slog"
)

type optContext struct {
	level      slog.Level
	timeFormat string
	time       bool
	source     bool

	json bool
}

type Opt func(ctx *optContext)

func newDefaultOptCtx() *optContext {
	return &optContext{
		level: slog.LevelDebug,
	}
}

func WithJson() Opt {
	return func(ctx *optContext) {
		ctx.json = true
	}
}

func WithLevel(level slog.Level) Opt {
	return func(ctx *optContext) {
		ctx.level = level
	}
}

func WithTime() Opt {
	return func(ctx *optContext) {
		ctx.time = true
	}
}

func WithSource() Opt {
	return func(ctx *optContext) {
		ctx.source = true
	}
}

func WithTimeFormat(format string) Opt {
	return func(ctx *optContext) {
		ctx.timeFormat = format
	}
}
