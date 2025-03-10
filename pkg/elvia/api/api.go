package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/3lvia/libraries-go/pkg/elvia/api/problemdetails"
	"github.com/3lvia/libraries-go/pkg/elvia/probe"
	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	sloggin "github.com/samber/slog-gin"
)

// NewDefaultServer creates a new HTTP server with default settings.
func NewDefaultServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

// NewDefaultEngine creates a new gin engine with default middleware.
// The environment is used to set the gin mode.
func NewDefaultEngine(env runtime.Env) *gin.Engine {
	switch env {
	case runtime.Development:
		gin.SetMode(gin.DebugMode)
	case runtime.Test:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(sloggin.NewWithConfig(slog.Default(), sloggin.Config{
		WithSpanID:  true,
		WithTraceID: true,
		Filters: []sloggin.Filter{
			sloggin.IgnorePath("/metrics", "/probe"),
		},
	}))
	engine.Use(gin.Recovery())

	return engine
}

// ConfigureStandardEndpoints configures standard endpoints for the API.
// This includes a 404 handler, a 405 handler, and a favicon handler.
func ConfigureStandardEndpoints(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		c.Render(http.StatusNotFound, problemdetails.U(
			http.StatusNotFound,
			"Not Found",
			fmt.Sprintf("path %s not found", c.Request.URL.Path)))
	})

	engine.NoMethod(func(c *gin.Context) {
		c.Render(http.StatusMethodNotAllowed, problemdetails.U(
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
			fmt.Sprintf("method %s not allowed", c.Request.Method)))
	})

	engine.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

// ConfigureStandardMetricsEndpoint configures a standard metrics endpoint for the API.
// The endpoint is exposed at /metrics and provides Prometheus metrics.
func ConfigureStandardMetricsEndpoint(engine *gin.Engine) {
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// ConfigureStandardHealthEndpoint configures a standard probe endpoint for the API.
// The endpoint is exposed at /probe and returns a JSON response.
func ConfigureStandardHealthEndpoint(engine *gin.Engine, check probe.HealthChecksFunc) {
	engine.GET("/health", func(c *gin.Context) {
		healths := check()
		summary := probe.Check(healths)

		httpStatus := http.StatusOK
		if summary.Status == probe.Unhealthy {
			httpStatus = http.StatusServiceUnavailable
		}

		c.JSON(httpStatus, summary)
	})
}