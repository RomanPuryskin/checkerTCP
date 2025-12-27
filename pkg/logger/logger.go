package logger

import (
	"log/slog"
	"os"

	"github.com/console_TCP/internal/config"
)

func InitLogger(cfg *config.LoggerConfig) *slog.Logger {

	var level slog.Level
	switch cfg.LogLevel {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)

}
