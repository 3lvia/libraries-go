package runtime

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

// NewLogger creates a new logger with the given log level.
func NewLogger(level slog.Level) *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

// LoggerFanout creates a new logger that fans out log messages to multiple handlers.
func LoggerFanout(logger *slog.Logger, handlers ...slog.Handler) *slog.Logger {
	handlers = append([]slog.Handler{logger.Handler()}, handlers...)
	handler := slogmulti.Fanout(handlers...)

	logger = slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
