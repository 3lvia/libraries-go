package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
)

const defaultTracerName = "go.opentelemetry.io/otel"

var tracerName string

func Authenticate(ctx context.Context, addr string, method Method, opts ...Option) (AuthenticationResponse, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	tracerName = collector.otelTracerName
	if tracerName == "" {
		tracerName = defaultTracerName
	}

	l := collector.l
	l.Printf("authenticating to %s using %s", addr, methodToString(method))

	tracer := otel.GetTracerProvider().Tracer(tracerName)
	spanCtx, span := tracer.Start(
		ctx,
		"auth.Authenticate",
		trace.WithAttributes(attribute.String("method", methodToString(method))))
	defer span.End()

	client := collector.client
	if client == nil {
		client = &http.Client{}
	}

	switch method {
	case MethodOICD:
		return authOICD(spanCtx, addr)
	case MethodGitHub:
		if collector.gitHubToken == "" {
			err := errors.New("no GitHub token provided")
			traceError(span, err)
			return nil, err
		}
		return authGitHub(spanCtx, addr, collector.gitHubToken, client)
	case MethodK8s:
		if collector.k8sServicePath == "" || collector.k8sRole == "" {
			err := errors.New("no k8s service path or role provided")
			traceError(span, err)
			return nil, err
		}
		l.Printf("using k8s service path %s and role %s", collector.k8sServicePath, collector.k8sRole)
		r, err := authK8s(spanCtx, addr, collector.k8sServicePath, collector.k8sRole, client)
		if err != nil {
			l.Printf("error authenticating to k8s: %s", err)
		} else {
			l.Printf("successfully authenticated to k8s, got client token of length %d", len(r.ClientToken()))
		}
		return r, err
	}

	err := fmt.Errorf("unknown authentication method: %s", methodToString(method))
	traceError(span, err)
	return nil, err
}

// authReq returns a http request for authenticating to Vault
func authReq(addr, path string, body *bytes.Buffer) (*http.Request, error) {
	url := makeURL(addr, path)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("while building http request: %w", err)
	}

	return req, nil
}

// makeURL returns a correctly formatted url for Vault http requests
func makeURL(address, path string) string {
	return address + "/v1/" + path
}

// loginBuffer converts a login token to a bytes buffer
func loginBuffer(lt interface{}) (*bytes.Buffer, error) {
	js, err := json.Marshal(lt)
	if err != nil {
		return nil, fmt.Errorf("while marshaling token: %w", err)
	}

	return bytes.NewBuffer(js), nil
}

// getJWT reads JSON web token from file at the service path
func getJWT(k8ServicePath string) (string, error) {
	b, err := os.ReadFile(k8ServicePath)
	if err == nil {
		//return "", fmt.Errorf("failed to read jwt token from %s: %w", k8ServicePath, err)
		return string(bytes.TrimSpace(b)), nil
	}

	b, err = os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return "", fmt.Errorf("failed to read jwt token from hard-coded path: %w", err)
	}

	return string(bytes.TrimSpace(b)), nil
}

func traceError(span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
	}
}
