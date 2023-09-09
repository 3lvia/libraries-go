package mschema

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"sync"
)

func newRegistry(collector *optionsCollector) Registry {
	tracerName := collector.otelTracerName
	if tracerName == "" {
		tracerName = defaultTracerName
	}

	schemaRegistryURL := collector.schemaRegistryURL
	user := collector.user
	password := collector.password

	currentHTTPClient := collector.client
	if currentHTTPClient == nil {
		currentHTTPClient = &http.Client{}
	}

	return &registry{
		url:         schemaRegistryURL,
		user:        user,
		password:    password,
		client:      currentHTTPClient,
		descriptors: map[string]Descriptor{},
		mux:         &sync.RWMutex{},
		tracerName:  tracerName,
	}
}

type registry struct {
	url      string
	user     string
	password string

	client *http.Client

	descriptors map[string]Descriptor
	mux         *sync.RWMutex

	tracerName string
}

func (r *registry) GetByID(ctx context.Context, id int) (Descriptor, error) {
	tracer := otel.GetTracerProvider().Tracer(r.tracerName)
	spanCtx, span := tracer.Start(
		ctx,
		"mschema.registry.GetByID",
		trace.WithAttributes(attribute.Int("id", id)))
	defer span.End()

	r.mux.RLock()
	k := fmt.Sprintf("ID-%d", id)
	if d, ok := r.descriptors[k]; ok {
		r.mux.RUnlock()
		return d, nil
	}
	r.mux.RUnlock()

	d, err := getById(spanCtx, id, r.url, r.user, r.password, r.client, r.tracerName)
	if err != nil {
		return nil, err
	}

	r.storeDescriptor(d)
	return d, nil
}

func (r *registry) GetBySubject(ctx context.Context, subject string) (Descriptor, error) {
	tracer := otel.GetTracerProvider().Tracer(r.tracerName)
	spanCtx, span := tracer.Start(
		ctx,
		"mschema.registry.GetBySubject",
		trace.WithAttributes(attribute.String("subject", subject)))
	defer span.End()

	r.mux.RLock()
	if d, ok := r.descriptors[subject]; ok {
		r.mux.RUnlock()
		return d, nil
	}
	r.mux.RUnlock()

	d, err := get(spanCtx, subject, r.url, r.user, r.password, r.client, r.tracerName)
	if err != nil {
		return nil, err
	}

	r.storeDescriptor(d)
	return d, nil
}

func (r *registry) storeDescriptor(d Descriptor) {
	r.mux.Lock()
	defer r.mux.Unlock()

	k := fmt.Sprintf("ID-%d", d.ID())

	r.descriptors[k] = d
	r.descriptors[d.Subject()] = d
}
