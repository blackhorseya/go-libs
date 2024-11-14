package contextx

import (
	"context"
	"testing"
)

func TestWithContext(t *testing.T) {
	c := context.Background()
	ctx := WithContext(c)

	ctx.Info("info message")
	ctx.Warn("warn message")
	ctx.Error("error message")
}

func TestSpanFromContext(t *testing.T) {
	c := context.Background()
	ctx, span := SpanFromContext(c, "span name")
	defer span.End()

	ctx.Info("info message")
	ctx.Warn("warn message")
	ctx.Error("error message")
}
