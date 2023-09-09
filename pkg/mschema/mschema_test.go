package mschema

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew_noSchemaRegistryURL(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "schema registry URL is required" {
		t.Errorf("expected error 'schema registry URL is required', got '%s'", err.Error())
	}
}

func TestNew_internalServerError(t *testing.T) {
	ctx := context.Background()

	exporter := setTestOtel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	server.Client()

	r, err := New(
		server.URL,
		WithClient(server.Client()),
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.GetBySubject(ctx, "private.dp.edna.examples")
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "unexpected status code: 500" {
		t.Errorf("expected error 'unexpected status code: 500', got '%s'", err.Error())
	}

	expectedSpans := []string{"mschema.get", "mschema.registry.GetBySubject"}
	checkSpans(expectedSpans, exporter, t)
}

func TestNew_happyDays(t *testing.T) {
	ctx := context.Background()

	exporter := setTestOtel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/subjects/private.dp.edna.examples-value/versions/latest" {
			t.Fatal("wrong path")
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(jsonSchema)); err != nil {
			t.Fatal(err)
		}
	}))

	r, err := New(
		server.URL,
		WithClient(server.Client()),
	)
	if err != nil {
		t.Fatal(err)
	}

	d, err := r.GetBySubject(ctx, "private.dp.edna.examples")
	if err != nil {
		t.Fatal(err)
	}
	if d.ID() != 100172 {
		t.Errorf("expected ID 100172, got %d", d.ID())
	}

	expectedSpans := []string{"mschema.get", "mschema.registry.GetBySubject"}
	checkSpans(expectedSpans, exporter, t)
}

func checkSpans(expectedSpans []string, exporter *tracetest.InMemoryExporter, t *testing.T) {
	m := map[string]tracetest.SpanStub{}
	for _, span := range exporter.GetSpans() {
		m[span.Name] = span
	}

	for _, spanName := range expectedSpans {
		if _, ok := m[spanName]; !ok {
			t.Errorf("expected span '%s' to be in spanMap", spanName)
		}
	}
}

func setTestOtel() *tracetest.InMemoryExporter {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(
		trace.WithSyncer(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)
	return exporter
}

const (
	jsonSchema = `{"subject":"private.dp.edna.examples-value","version":6,"id":100172,"schema":"{\"type\":\"record\",\"name\":\"Person\",\"namespace\":\"dp.demoapp\",\"fields\":[{\"name\":\"Cars\",\"type\":[\"null\",{\"type\":\"array\",\"items\":[\"null\",{\"type\":\"record\",\"name\":\"Car\",\"fields\":[{\"name\":\"Color\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Gearbox\",\"type\":[\"null\",{\"type\":\"enum\",\"name\":\"Gearbox\",\"symbols\":[\"Automatic\",\"Manual\"]}],\"default\":null},{\"name\":\"model\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Person\",\"type\":[\"null\",\"Person\"],\"default\":null}]}]}],\"default\":null},{\"name\":\"Houses\",\"type\":[\"null\",{\"type\":\"array\",\"items\":[\"null\",{\"type\":\"record\",\"name\":\"House\",\"fields\":[{\"name\":\"Address\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Buildingtype\",\"type\":{\"type\":\"enum\",\"name\":\"Buildingtype\",\"symbols\":[\"House\",\"Apartment\",\"Cabin\",\"Other\"]}},{\"name\":\"Color\",\"type\":[\"null\",\"string\"],\"default\":null}]}]}],\"default\":null},{\"name\":\"Id\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Name\",\"type\":[\"null\",\"string\"],\"default\":null},{\"name\":\"Updated\",\"type\":[\"null\",\"string\"],\"default\":null}]}"}`
)
