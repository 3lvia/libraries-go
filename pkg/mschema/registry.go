package mschema

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"sync"
)

func newRegistry(url string, collector *optionsCollector) (Registry, error) {
	tracerName := collector.otelTracerName
	if tracerName == "" {
		tracerName = defaultTracerName
	}

	if url == "" {
		return nil, fmt.Errorf("schema registry URL is required")
	}

	user := collector.user
	password := collector.password

	currentHTTPClient := collector.client
	if currentHTTPClient == nil {
		currentHTTPClient = &http.Client{}
	}

	r := &registry{
		url:         url,
		user:        user,
		password:    password,
		client:      currentHTTPClient,
		descriptors: map[string]Descriptor{},
		mux:         &sync.RWMutex{},
		tracerName:  tracerName,
	}

	return r, nil
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

func (r *registry) List(ctx context.Context) ([]Descriptor, error) {
	tracer := otel.GetTracerProvider().Tracer(r.tracerName)
	_, span := tracer.Start(
		ctx,
		"mschema.registry.List")
	defer span.End()

	url := fmt.Sprintf("%s/schemas?deleted=false&latestOnly=true", r.url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.user, r.password)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var descriptors []descriptor
	err = json.Unmarshal(body, &descriptors)
	if err != nil {
		return nil, err
	}

	var res []Descriptor
	for _, d := range descriptors {
		res = append(res, d)
	}
	return res, nil
}

func (r *registry) storeDescriptor(d Descriptor) {
	r.mux.Lock()
	defer r.mux.Unlock()

	k := fmt.Sprintf("ID-%d", d.ID())

	r.descriptors[k] = d
	r.descriptors[d.Subject()] = d
}
