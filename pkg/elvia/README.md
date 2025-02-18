# Elvia Service

Elvia Service is a library for building services in Elvia with common functionality.

Common functionality includes:
- Logging using slog
- OpenTelemetry tracing, logging and metrics
- Health and metrics endpoints
- Graceful shutdown

## Installation

```bash
go get github.com/3lvia/libraires-go/pkg/elvia
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

```go
package main

type MyService struct {
	*elvia.Service
	// Add additional components here
}

func NewMyService(ctx context.Context) (*MyService, error) {
	svc, err := elvia.NewService(ctx, "my-system", "my-service")
	if err != nil {
		return nil, err
	}

	// Create additional components here

	return &MyService{svc}, nil
}

func (s *MyService) Run(ctx context.Context) error {
	// Do something before running the service
	err := s.Service.Run(ctx)
	if err != nil {
		return err
	}

	// Do something after running the service
	return nil
}

func (s *MyService) Stop(ctx context.Context) error {
	// Do something before stopping the service
	err := s.Service.Stop(ctx)
	if err != nil {
		return err
	}

	// Do something after stopping the service
	return nil
}

func main() {
	ctx := context.Background()

	svc, err := NewMyService(ctx)
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

## Options

The `elvia.NewService` function takes in optional `elvia.ServiceOpt` options. Use the `elvia.With...` functions to add them.

```go
func main() {
    opts := []elvia.ServiceOpt{
		elvia.WithEnvironment(runtime.Development),
		elvia.WithLoggerLevel(slog.LevelDebug),
    }
    
    svc, err := elvia.NewService(ctx, "my-system", "my-service", opts...)
}
```