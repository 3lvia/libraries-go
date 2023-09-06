package hashivault

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log"
	"net/http"
	"sync"
)

// New returns a new SecretsManager and also a channel that will send errors that may arise in the concurrent internal
// goroutines that will run in the whole lifetime of the service after this function. The returned error indicates that
// something went wrong during initialization, and the service will not be able to run (if it is not nil).
func New(ctx context.Context, opts ...Option) (SecretsManager, <-chan error, error) {
	c := &optionsCollector{}
	for _, opt := range opts {
		opt(c)
	}

	l := c.logger
	if l == nil {
		l = log.New(nullWriter(1), "", log.LstdFlags)
	}

	tracerName = c.otelTracerName
	if tracerName == "" {
		tracerName = defaultTracerName
	}

	l.Printf("starting hashivault secrets manager with tracer: %s", tracerName)

	tracer := otel.GetTracerProvider().Tracer(tracerName)
	spanCtx, span := tracer.Start(ctx, "hashivault.New")
	defer span.End()

	if err := c.build(); err != nil {
		traceError(span, err, l)
		return nil, nil, fmt.Errorf("invalid options: %w", err)
	}

	span.SetAttributes(attribute.String("vault_address", c.vaultAddress))
	l.Printf("using vault address: %s", c.vaultAddress)

	errChan := make(chan error)
	client := c.client
	if client == nil {
		client = &http.Client{}
	}

	tokenGetter := func() string {
		return c.vaultToken
	}
	if c.vaultToken != "" {
		l.Print("using static vault token")
	}
	if c.vaultToken == "" {
		// initializedChan is used to signal that the tokenGetter has been initialized. This ensures that secrets are not
		// requested before we have a valid token. This channel should only be used once, and no actual message will ever
		// be sent on it. Instead, it will be closed when the tokenGetter has been initialized.
		initializedChan := make(chan struct{})

		tokenGetter = startTokenJob(spanCtx, c, errChan, initializedChan, client, l)

		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(ch <-chan struct{}, w *sync.WaitGroup) {
			defer w.Done()
			<-ch
		}(initializedChan, wg)

		wg.Wait()
	}

	l.Print("hashivault secrets manager initialized, ready to go!")

	m := newManager(c.vaultAddress, tokenGetter, errChan, l)
	return m, errChan, nil
}
