package elvia

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/3lvia/libraries-go/pkg/elvia/api"
	"github.com/3lvia/libraries-go/pkg/elvia/probe"
	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

type ShutdownFunc func(context.Context) error

// Service configures standard service components and manages their lifecycle.
type Service struct {
	logger        *slog.Logger
	apiServer     *http.Server
	apiEngine     *gin.Engine
	shutdownFuncs []ShutdownFunc
	healthChecks  probe.HealthChecks
}

// NewService creates a new service with the given service name and options.
func NewService(ctx context.Context, systemName, serviceName string, opts ...ServiceOpt) (*Service, error) {

	cfg := defaultConfig(systemName, serviceName)
	for _, opt := range opts {
		opt(&cfg)
	}

	logger := runtime.NewLogger(cfg.loggerLevel)

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			cfg.otelAttributes...,
		),
	)
	if err != nil {
		return nil, err
	}

	shutdownFuncs := make([]ShutdownFunc, 0)

	if cfg.otelEnabled {
		otel.SetTextMapPropagator(cfg.otelPropagator)

		traceProvider, err := cfg.otelNewTraceProvider(ctx, cfg.env, trace.WithResource(r))
		if err != nil {
			return nil, err
		}
		if traceProvider != nil {
			shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)
			otel.SetTracerProvider(traceProvider)
		}

		loggerProvider, err := cfg.otelNewLoggerProvider(ctx, cfg.env, log.WithResource(r))
		if err != nil {
			return nil, err
		}
		if loggerProvider != nil {
			shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
			global.SetLoggerProvider(loggerProvider)

			logger = runtime.LoggerFanout(logger, otelslog.NewHandler(serviceName, otelslog.WithLoggerProvider(loggerProvider)))
		}

		metricProvider, err := cfg.otelNewMetricProvider(ctx, cfg.env, metric.WithResource(r))
		if err != nil {
			return nil, err
		}
		if metricProvider != nil {
			shutdownFuncs = append(shutdownFuncs, metricProvider.Shutdown)
			otel.SetMeterProvider(metricProvider)
		}
	}

	var apiServer *http.Server
	var apiEngine *gin.Engine
	if cfg.withApiAddr != DisableAPI {
		slog.Info("API server is enabled", "addr", cfg.withApiAddr)

		apiEngine = cfg.withApiEngine(cfg.env)
		if apiEngine != nil {
			if cfg.withApiOTELMiddleware != nil {
				cfg.withApiOTELMiddleware(apiEngine)
			}

			for _, endpoint := range cfg.withApiEndpoints {
				endpoint(apiEngine)
			}
		}

		if cfg.withHTTPServer != nil {
			apiServer = cfg.withHTTPServer
		} else {
			apiServer = api.NewDefaultServer(cfg.withApiAddr, apiEngine)
		}

		shutdownFuncs = append(shutdownFuncs, func(ctx context.Context) error {
			slog.InfoContext(ctx, "API server is shutting down")
			defer slog.InfoContext(ctx, "API server has shut down")

			if err := apiServer.Shutdown(ctx); err != nil {
				_ = apiServer.Close()
				slog.Error("could not stop the API server gracefully", "error", err)
				return err
			}
			return nil
		})
	}

	svc := &Service{
		logger:        logger,
		shutdownFuncs: shutdownFuncs,
		apiServer:     apiServer,
		apiEngine:     apiEngine,
		healthChecks:  make(probe.HealthChecks),
	}

	if apiEngine != nil && cfg.withApiHealthEndpoint != nil {
		cfg.withApiHealthEndpoint(apiEngine, func() probe.HealthReports {
			reports := make(probe.HealthReports)
			for name, check := range svc.healthChecks {
				reports[name] = check()
			}
			return reports
		})
	}

	return svc, nil
}

// RegisterShutdown registers a function that will be run when the service is stopped.
func (s *Service) RegisterShutdown(fn ShutdownFunc) {
	s.shutdownFuncs = append(s.shutdownFuncs, fn)
}

// RegisterHealthCheck registers a health check function that will affect the status of the health endpoint.
func (s *Service) RegisterHealthCheck(name string, fn probe.HealthReportFunc) {
	s.healthChecks[name] = fn
}

// Run starts the service and blocks until the service is stopped.
// The stopping mechanism is when the context is cancelled or a SIGINT or SIGTERM signal is received.
// Stop should be called to stop the service.
func (s *Service) Run(ctx context.Context) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	var apiErrChan chan error
	if s.apiServer != nil {
		apiErrChan = make(chan error, 1)
		go func() {
			slog.InfoContext(ctx, "API server started", "addr", s.apiServer.Addr)
			err := s.apiServer.ListenAndServe()
			apiErrChan <- err
		}()
	}

	select {
	case err := <-apiErrChan:
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			slog.InfoContext(ctx, "API server stopped")
		} else {
			slog.ErrorContext(ctx, "API server stopped unexpectedly", "error", err)
		}
	case err := <-ctx.Done():
		slog.InfoContext(ctx, "shutting down service", "error", err)
	case <-done:
		slog.InfoContext(ctx, "shutting down service")
	}

	return nil
}

// Stop stops the service and all its components.
func (s *Service) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var err error
	for _, fn := range slices.Backward(s.shutdownFuncs) {
		err = errors.Join(err, fn(ctx))
	}
	s.shutdownFuncs = nil

	slog.InfoContext(ctx, "service stopped")

	return err
}