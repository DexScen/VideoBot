package utils

import (
	"log/slog"
	"os"
	"time"
)

func NewSlog(f *os.File) {
	logger := slog.New(slog.NewJSONHandler(f,
		&slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					a.Value = slog.StringValue(time.Now().Format("2006-01-02 15:04:05.000"))
				}
				return a
			},
		}))
	slog.SetDefault(logger)
}
