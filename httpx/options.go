package httpx

import (
	"time"
)

// ServerOption is a function that configures the server
type ServerOption func(opts *serverOptions)

type serverOptions struct {
	Host            string
	Port            int
	ShutdownTimeout time.Duration
}

// WithHost sets the host of the server
func WithHost(host string) ServerOption {
	return func(opts *serverOptions) {
		opts.Host = host
	}
}

// WithPort sets the port of the server
func WithPort(port int) ServerOption {
	return func(opts *serverOptions) {
		opts.Port = port
	}
}

// WithShutdownTimeout sets the shutdown timeout of the server
func WithShutdownTimeout(timeout time.Duration) ServerOption {
	return func(opts *serverOptions) {
		opts.ShutdownTimeout = timeout
	}
}
