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
