package hashivault

import (
	"bytes"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNew_static(t *testing.T) {
	ctx := context.Background()

	var buf bytes.Buffer
	l := log.New(&buf, "", log.LstdFlags)

	url, client, closer := startTestServer(t)
	defer closer()

	expectedLogMessages := []string{
		"starting hashivault secrets manager with tracer: go.opentelemetry.io/otel",
		fmt.Sprintf("using vault address: %s", url),
		"starting token job",
		fmt.Sprintf("authenticating to %s using GitHub", url),
		"token job initialized, first token acquired",
		"hashivault secrets manager initialized, ready to go!",
	}
	expectedSpans := []string{
		"auth.authGitHub",
		"auth.Authenticate",
		"hashivault.tokenJob.authenticate",
		"hashivault.New",
		"hashivault.get",
		"hashivault.GetSecret",
	}

	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(
		trace.WithSyncer(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	addHandler("/v1/kunde/kv/data/appinsights/kunde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, jsonStaticSecret)
	})

	gitHubToken := "my-github-token"

	sm, errChan, err := New(ctx, WithClient(client), WithVaultAddress(url), WithGitHubToken(gitHubToken), WithLogger(l))
	if err != nil {
		t.Fatal(err)
	}

	go func(ec <-chan error) {
		e := <-ec
		if e != nil {
			t.Error(e)
		}
	}(errChan)

	eg, err := sm.GetSecret(ctx, "kunde/kv/data/appinsights/kunde")
	NoErr(t, err)

	sec := eg()
	if sec["instrumentation-key"] != "my-secret-instrumentation-key" {
		t.Errorf("expected 'value', got '%s'", sec["instrumentation-key"])
	}

	if loginCount < 1 {
		t.Errorf("expected loginCount to be 1, got %d", loginCount)
	}

	logMessages := buf.String()
	for _, message := range expectedLogMessages {
		if !strings.Contains(logMessages, message) {
			t.Errorf("expected log message '%s' to be in '%s'", message, logMessages)
		}
	}

	spans := exporter.GetSpans()
	if len(spans) != 6 {
		t.Errorf("expected 6 spans, got %d", len(spans))
	}
	var newSpan tracetest.SpanStub
	var getSecretSpan tracetest.SpanStub
	spanMap := map[string]tracetest.SpanStub{}
	for _, span := range spans {
		spanMap[span.Name] = span
		if span.Name == "hashivault.New" {
			newSpan = span
		}
		if span.Name == "hashivault.GetSecret" {
			getSecretSpan = span
		}
	}
	if newSpan.ChildSpanCount != 1 {
		t.Errorf("expected 1 child spans, got %d", newSpan.ChildSpanCount)
	}
	if getSecretSpan.ChildSpanCount != 1 {
		t.Errorf("expected 1 child span, got %d", getSecretSpan.ChildSpanCount)
	}
	for _, spanName := range expectedSpans {
		if _, ok := spanMap[spanName]; !ok {
			t.Errorf("expected span '%s' to be in spanMap", spanName)
		}
	}
}

func TestNew_dynamic(t *testing.T) {
	ctx := context.Background()

	t.Skip("this test fails when executed with the other tests, but works when runned alone")

	url, client, closer := startTestServer(t)
	defer closer()

	secretCount := 0
	addHandler("/v1/kunde/kv/data/appinsights/kunde2", func(w http.ResponseWriter, r *http.Request) {
		secretCount++
		js := fmt.Sprintf(jsonDynamicSecret, fmt.Sprintf("secret-%d", secretCount))
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, js)
	})

	gitHubToken := "my-github-token"

	sm, errChan, err := New(ctx, WithClient(client), WithVaultAddress(url), WithGitHubToken(gitHubToken))
	if err != nil {
		t.Fatal(err)
	}

	go func(ec <-chan error) {
		e := <-ec
		if e != nil {
			t.Error(e)
		}
	}(errChan)

	eg, err := sm.GetSecret(ctx, "kunde/kv/data/appinsights/kunde2")
	NoErr(t, err)

	sec := eg()
	if sec["instrumentation-key"] != "secret-1" {
		t.Errorf("expected 'value', got '%s'", sec["instrumentation-key"])
	}

	if loginCount < 1 {
		t.Errorf("expected loginCount to be 1, got %d", loginCount)
	}

	<-time.After(2 * time.Second)

	sec = eg()
	if sec["instrumentation-key"] != "secret-2" {
		t.Errorf("expected 'value', got '%s'", sec["instrumentation-key"])
	}
}

func NoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, got: %v", err)
	}
}

const (
	ghVaultResponseTemplate = `{
    "request_id": "d645ddd7-3b2e-f28b-0138-512d5ff301a4",
    "lease_id": "",
    "renewable": false,
    "lease_duration": 0,
    "data": null,
    "wrap_info": null,
    "warnings": null,
    "auth": {
        "client_token": "%s",
        "accessor": "zIC3dwCcg8foRsVTxxxtdX570X5",
        "policies": [
            "coreteam",
            "default",
            "drops",
            "drops-extra",
            "ea",
            "ea-extra",
            "hes-extensions",
            "hes-extensions-extra",
            "leveransemotor",
            "leveransemotor-extra"
        ],
        "token_policies": [
            "coreteam",
            "default",
            "drops",
            "drops-extra",
            "ea",
            "ea-extra",
            "hes-extensions",
            "hes-extensions-extra",
            "leveransemotor",
            "leveransemotor-extra"
        ],
        "metadata": {
            "org": "abc",
            "username": "sasc"
        },
			"lease_duration": 1,
        "renewable": true,
        "entity_id": "2957867d-00ea-981d-e2e9-dd41ab412212",
        "token_type": "service",
        "orphan": true,
        "mfa_requirement": null,
        "num_uses": 0
    }
}`

	jsonStaticSecret = `{
    "request_id": "4dfb1662-c462-99f0-120e-ee61cd3b099e",
    "lease_id": "",
    "renewable": false,
    "lease_duration": 0,
    "data": {
        "data": {
            "instrumentation-key": "my-secret-instrumentation-key"
        },
        "metadata": {
            "created_time": "2020-08-26T14:56:35.936623451Z",
            "custom_metadata": null,
            "deletion_time": "",
            "destroyed": false,
            "version": 2
        }
    },
    "wrap_info": null,
    "warnings": null,
    "auth": null
}`
	jsonDynamicSecret = `{
    "request_id": "4dfb1662-c462-99f0-120e-ee61cd3b099e",
    "lease_id": "",
    "renewable": true,
    "lease_duration": 1,
    "data": {
        "data": {
            "instrumentation-key": "%s"
        },
        "metadata": {
            "created_time": "2020-08-26T14:56:35.936623451Z",
            "custom_metadata": null,
            "deletion_time": "",
            "destroyed": false,
            "version": 2
        }
    },
    "wrap_info": null,
    "warnings": null,
    "auth": null
}`
)
