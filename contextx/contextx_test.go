package contextx

import (
	"context"
	"log/slog"
	"testing"
)

func TestWithContext(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	c := context.Background()
	ctx := WithContext(c)

	ctx.Debug("debug message")
	ctx.Info("info message")
	ctx.Warn("warn message")
	ctx.Error("error message")
}
