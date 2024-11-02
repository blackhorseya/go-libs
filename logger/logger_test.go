package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := Init(Options{
		Level:  "debug",
		Format: "console",
	})
	if err != nil {
		t.Errorf("failed to init logger: %v", err)
	}

	// Test the logger
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}
