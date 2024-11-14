package contextx

import (
	"context"

	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()
	defer func() {
		_ = logger.Sync()
	}()
}

// Contextx is a struct that holds the context of the request
type Contextx struct {
	context.Context
	*zap.Logger
}

// WithContext is a function that returns a new Contextx with the context of the request
func WithContext(c context.Context) Contextx {
	return Contextx{
		c,
		zap.L(),
	}
}
