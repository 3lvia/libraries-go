package elvia

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type shutdownFunc func(context.Context) error

// Service configures standard service components and manages their lifecycle.
type Service struct {
	logger        *slog.Logger
	apiServer     *http.Server
	shutdownFuncs []shutdownFunc
}

// New creates a new service with the given service name and options.
func New(ctx context.Context, serviceName string, opts ...Opt) (*Service, error) {
	cfg := defaultConfig(serviceName)
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

	shutdownFuncs := make([]shutdownFunc, 0)
	otel.SetTextMapPropagator(cfg.otelPropagator)

	traceProvider, err := cfg.otelNewTraceProvider(ctx, cfg.env, trace.WithResource(r))
	if err != nil {
		return nil, err
	}
	if traceProvider != nil {
		shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)
		otel.SetTracerProvider(traceProvider)
	}

	loggerProvider, err := cfg.otelNewLoggerProvider(ctx, cfg.env)
	if err != nil {
		return nil, err
	}
	if loggerProvider != nil {
		shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
		global.SetLoggerProvider(loggerProvider)

		logger = runtime.LoggerFanout(logger, otelslog.NewHandler(serviceName, otelslog.WithLoggerProvider(loggerProvider)))
	}

	metricProvider, err := cfg.otelNewMetricProvider(ctx, cfg.env)
	if err != nil {
		return nil, err
	}
	if metricProvider != nil {
		shutdownFuncs = append(shutdownFuncs, metricProvider.Shutdown)
		otel.SetMeterProvider(metricProvider)
	}

	if cfg.withApiAddr != DisableAPI {
		engine := cfg.withApiEngine(cfg.env)
		if engine != nil {
			for _, endpoint := range cfg.withApiEndpoints {
				endpoint(engine)
			}
		}

		apiServer := &http.Server{
			Addr:    cfg.withApiAddr,
			Handler: engine,
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

	return &Service{
		logger:        logger,
		shutdownFuncs: shutdownFuncs,
		apiServer:     cfg.withHTTPServer,
	}, nil
}

// Run starts the service and blocks until the service is stopped.
// The stopping mechanism is when the context is cancelled or a SIGINT or SIGTERM signal is received.
func (s *Service) Run(ctx context.Context) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	var apiErrChan chan error
	if s.apiServer != nil {
		apiErrChan = make(chan error, 1)
		go func() {
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

	return s.Stop(ctx)
}

// Stop stops the service and all its components.
func (s *Service) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	var err error
	for _, fn := range s.shutdownFuncs {
		err = errors.Join(err, fn(ctx))
	}
	s.shutdownFuncs = nil

	slog.InfoContext(ctx, "service stopped")

	return err
}
