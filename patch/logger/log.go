package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(w io.Writer, opts ...Opt) *Logger {
	ctx := newDefaultOptCtx()
	for _, opt := range opts {
		opt(ctx)
	}

	handleOpts := &slog.HandlerOptions{}
	handleOpts.Level = ctx.level
	handleOpts.AddSource = ctx.source

	replaceTime := func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key != slog.TimeKey ||
			attr.Value.Kind() != slog.KindTime ||
			len(groups) != 0 {
			return attr
		}

		if !ctx.time {
			return slog.Attr{}
		}
		if ctx.timeFormat != "" {
			v := attr.Value.Time().Format(ctx.timeFormat)
			return slog.String(slog.TimeKey, v)
		}
		return attr
	}

	handleOpts.ReplaceAttr = func(groups []string, attr slog.Attr) slog.Attr {
		return replaceTime(groups, attr)
	}

	var handler slog.Handler
	if ctx.json {
		handler = slog.NewJSONHandler(w, handleOpts)
	} else {
		handler = slog.NewTextHandler(w, handleOpts)
	}

	log := slog.New(handler)
	return &Logger{Logger: log}
}

type OutputCtrl struct {
	DebugOut io.Writer
	InfoOut  io.Writer
	WarnOut  io.Writer
	ErrorOut io.Writer
}

type levelLogger struct {
	OutputCtrl
	debugLogger *Logger
	infoLogger  *Logger
	warnLogger  *Logger
	errLogger   *Logger
}

func (l *levelLogger) getOutput(level slog.Level) io.Writer {
	var out io.Writer
	switch level {
	case slog.LevelDebug:
		out = l.DebugOut
	case slog.LevelInfo:
		out = l.InfoOut
	case slog.LevelWarn:
		out = l.WarnOut
	case slog.LevelError:
		out = l.ErrorOut
	}
	if out != nil {
		return out
	}
	return os.Stdout
}

func (l *levelLogger) getLogger(level slog.Level) *Logger {
	switch level {
	case slog.LevelDebug:
		return l.debugLogger
	case slog.LevelInfo:
		return l.infoLogger
	case slog.LevelWarn:
		return l.warnLogger
	case slog.LevelError:
		return l.errLogger
	default:
		panic(fmt.Sprintf("bad level: %d", level))
	}
}

func (l *levelLogger) setLogger(level slog.Level, logger *Logger) {
	switch level {
	case slog.LevelDebug:
		l.debugLogger = logger
	case slog.LevelInfo:
		l.infoLogger = logger
	case slog.LevelWarn:
		l.warnLogger = logger
	case slog.LevelError:
		l.errLogger = logger
	default:
		panic(fmt.Sprintf("bad level: %d", level))
	}
}

func SetDefault(level2Out OutputCtrl, opts ...Opt) {
	ctx := newDefaultOptCtx()
	for _, opt := range opts {
		opt(ctx)
	}
	ll := &levelLogger{}
	ll.OutputCtrl = level2Out

	outputs := make(map[io.Writer][]slog.Level)
	levels := []slog.Level{
		slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError,
	}

	for _, level := range levels {
		out := ll.getOutput(level)
		list, ok := outputs[out]
		if !ok {
			list = make([]slog.Level, 0, 1)
		}
		list = append(list, level)
		outputs[out] = list
	}

	for output, lvs := range outputs {
		if len(lvs) == 0 {
			continue
		}
		logger := NewLogger(output, opts...)
		for _, lv := range lvs {
			ll.setLogger(lv, logger)
		}
	}
	setGlobalLevelLogger(ll)
}
