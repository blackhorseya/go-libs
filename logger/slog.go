package logger

import (
	"log/slog"
	"os"
)

// NewSlogLogger initializes the logging instance.
func NewSlogLogger(options Options) (*slog.Logger, error) {
	var level slog.Level
	err := level.UnmarshalText([]byte(options.Level))
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}

	var handler slog.Handler
	switch options.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger, nil
}
