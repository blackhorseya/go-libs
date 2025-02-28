package httpx

import (
	"log/slog"
	"net/http"

	"github.com/blackhorseya/go-libs/logger"
	"github.com/gin-gonic/gin"
)

// GinServer is a wrapper around the http.Server and gin.Engine
type GinServer struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewGinServer creates a new GinServer
func NewGinServer(log *slog.Logger, debug bool) *GinServer {
	middlewares := []gin.HandlerFunc{
		logger.GinTraceLoggingMiddleware(log),
	}

	gin.SetMode(gin.ReleaseMode)
	if debug {
		gin.SetMode(gin.DebugMode)
		middlewares = append(middlewares, gin.Logger())
	}

	router := gin.New()
	router.Use(middlewares...)

	return &GinServer{
		httpserver: nil,
		Router:     router,
	}
}
