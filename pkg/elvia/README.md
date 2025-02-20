# Elvia Service

Elvia Service is a library for building services in Elvia with common functionality.

Common functionality includes:
- Logging using slog
- OpenTelemetry tracing, logging and metrics
- Health and metrics endpoints
- Graceful shutdown

## Installation

```bash
go get github.com/3lvia/libraries-go/pkg/elvia
```

## Usage

```go
package main

import (
	"context"
	"log/slog"

	"github.com/3lvia/libraries-go/pkg/elvia"
)

func main() {
	ctx := context.Background()

	// Note: the system and service name will be used for OpenTelemetry to identify the service
	svc, err := elvia.NewService(ctx, "my-system", "my-service")
	if err != nil {
        slog.Error("failed to create service", "error", err)
        panic(err)
    }

	// Run is a blocking call and will only return on error or,
	// if the service is stopped by the ctx or SIGTERM/SIGINT
	err = svc.Run(ctx)
	if err != nil {
		slog.Error("failed to run service", "error", err)
		panic(err)    
	}
	
	// Call Stop to gracefully shutdown the service
	err = svc.Stop(ctx)
	if err != nil {
        slog.Error("failed to stop service", "error", err)
        panic(err)    
    }
}
```

Even though the service can run on its own, it is likely that you will want additional components for your service.
Wrap the service in a struct and add additional components as needed.

```go
package main

import (
	"context"
	"log/slog"

	"github.com/3lvia/libraries-go/pkg/elvia"
	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"golang.org/x/telemetry/internal/config"
)

type MyService struct {
	*elvia.Service

	// Add additional components here
}

func NewMyService(ctx context.Context, cfg *Config) (*MyService, error) {
	// Configure the elvia service with elvia.With... options
	opts := []elvia.ServiceOpt{
		elvia.WithEnvLoggerLevel(cfg.Env),
		elvia.WithAPI(cfg.APIAddr),
	}

	svc, err := elvia.NewService(ctx, "my-system", "my-service", opts...)
	if err != nil {
		return nil, err
	}

	// Create additional components here

	return &MyService{svc}, nil
}

func (s *MyService) Run(ctx context.Context) error {
	// Start additional components here

	return s.Service.Run(ctx)
}

func (s *MyService) Stop(ctx context.Context) error {
	// Stop additional components here

	return s.Service.Stop(ctx)
}

type Config struct {
	Env runtime.Env
    APIAddr string
}

func main() {
	// Load your configuration
	cfg := &Config{
        Env: runtime.Development,
        APIAddr: ":8080",
    }
	
	ctx := context.Background()

	svc, err := NewMyService(ctx, cfg)
	if err != nil {
		slog.Error("failed to create service", "error", err)
		panic(err)
	}

	err = svc.Run(ctx)
	if err != nil {
		slog.Error("failed to run service", "error", err)
		panic(err)
	}

	err = svc.Stop(ctx)
	if err != nil {
		slog.Error("failed to stop service", "error", err)
		panic(err)
	}
}
```