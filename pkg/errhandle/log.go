package errhandle

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	Log = NewLogger()
	cb  = context.Background()
)

type Logger struct {
	*slog.Logger
}

type LoggerOptions func(*Logger)

func NewLogger(opts ...LoggerOptions) *Logger {
	l := &Logger{
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// Remove time.
				if a.Key == slog.TimeKey && len(groups) == 0 {
					return slog.Attr{}
				}
				// Remove the directory from the source's filename.
				if a.Key == slog.SourceKey {
					source := a.Value.Any().(*slog.Source)
					source.File = filepath.Base(source.File)
				}
				if a.Key == slog.MessageKey {
					return slog.Attr{
						Key:   "",
						Value: a.Value,
					}
				}
				return a
			},
		})),
	}
	for _, o := range opts {
		o(l)
	}
	return l
}

func (l *Logger) Errorln(s ...any) {
	if !l.Enabled(cb, slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintln(s...), pcs[0])
	_ = l.Handler().Handle(cb, r)
}

func (l *Logger) Errorf(format string, s ...any) {
	if !l.Enabled(cb, slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, s...), pcs[0])
	_ = l.Handler().Handle(cb, r)
}

func (l *Logger) Fatal(s ...any) {
	l.Errorln(s...)
	os.Exit(1)
}
