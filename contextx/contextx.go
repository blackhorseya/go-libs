package contextx

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
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

// SpanFromContext is a function that returns a new Contextx with the context of the request
func SpanFromContext(c context.Context, spanName string, opts ...trace.SpanStartOption) (Contextx, trace.Span) {
	next, span := otel.Tracer("contextx").Start(c, spanName, opts...)
	return Contextx{
		next,
		zap.L(),
	}, span
}
