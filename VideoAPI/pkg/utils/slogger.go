package utils

import (
	"log/slog"
	"os"
	"time"
)

func NewSlog() error{
	f, err := os.OpenFile("VideoAPI\\pkg\\utils\\videoapi.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil{
		return err
	}
	logger := slog.New(slog.NewJSONHandler(f,
		&slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					a.Value = slog.StringValue(time.Now().Format("yyyy-MM-dd HH:mm:ss.SSS"))
				}
				return a
			},
		}))
	slog.SetDefault(logger)
	return nil
}
