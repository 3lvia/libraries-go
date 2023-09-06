package hashivault

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"sync"
	"time"
)

func newEvergreen(path, vaultAddress string, sec *secret, tokenGetter tokenGetterFunc, client *http.Client, errChan chan<- error, l *log.Logger) *evergreenSecret {
	eg := &evergreenSecret{
		path:         path,
		sec:          sec,
		mux:          &sync.Mutex{},
		client:       client,
		vaultAddress: vaultAddress,
		tokenGetter:  tokenGetter,
		l:            l,
	}

	go eg.start(errChan)

	return eg
}

type evergreenSecret struct {
	path         string
	vaultAddress string
	client       *http.Client
	sec          *secret
	mux          *sync.Mutex
	tokenGetter  tokenGetterFunc
	l            *log.Logger
}

func (e *evergreenSecret) get() map[string]any {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.sec.data()
}

func (e *evergreenSecret) start(errChan chan<- error) {
	for {
		<-time.After(time.Duration(e.sec.LeaseDuration) * time.Second)
		e.mux.Lock()

		tracer := otel.GetTracerProvider().Tracer(tracerName)
		ctx, span := tracer.Start(
			context.Background(),
			"hashivault.evergreenSecret.start",
			trace.WithAttributes(attribute.String("path", e.path)))
		defer span.End()

		sec, err := get(ctx, e.path, e.vaultAddress, e.tokenGetter(), e.client, e.l)
		if err != nil {
			errChan <- err
			e.mux.Unlock()
			continue
		}
		e.sec = sec
		e.mux.Unlock()
	}
}
