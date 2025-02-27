package logger

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options is the logging options.
type Options struct {
	// Level is the log level. options: debug, info, warn, error, dpanic, panic, fatal (default: info)
	Level string `json:"level" yaml:"level" mapstructure:"level"`

	// Format is the log format. options: json, text (default: text)
	Format string `json:"format" yaml:"format" mapstructure:"format"`
}

// NewZapLogger initializes the logging instance.
func NewZapLogger(options Options) (*zap.Logger, error) {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(options.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(zapConfig)
	if options.Format == "json" {
		zapConfig = zap.NewProductionEncoderConfig()
		enc = zapcore.NewJSONEncoder(zapConfig)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	zap.ReplaceGlobals(logger)

	return logger, nil
}

// NewSlogLogger initializes the logging instance.
func NewSlogLogger(options Options) (*slog.Logger, error) {
	var level slog.Level
	switch options.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
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
