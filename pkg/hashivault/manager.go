package hashivault

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"net/http"
	"os"
)

func newManager(vaultAddress string, tokenGetter tokenGetterFunc, errChan chan<- error, l *log.Logger) *manager {
	return &manager{
		vaultAddress: vaultAddress,
		client:       &http.Client{},
		tokenGetter:  tokenGetter,
		errChan:      errChan,
		l:            l,
	}
}

type manager struct {
	vaultAddress string
	client       *http.Client
	tokenGetter  tokenGetterFunc
	errChan      chan<- error
	l            *log.Logger
}

func (m *manager) GetSecret(ctx context.Context, path string) (EvergreenSecretsFunc, error) {
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	spanCtx, span := tracer.Start(
		ctx,
		"hashivault.GetSecret",
		trace.WithAttributes(attribute.String("path", path)))
	defer span.End()

	m.l.Printf("getting secrets from %s", path)

	sec, err := get(spanCtx, path, m.vaultAddress, m.tokenGetter(), m.client, m.l)
	if err != nil {
		return nil, err
	}

	if !sec.Renewable {
		return sec.data, nil
	}

	es := newEvergreen(path, m.vaultAddress, sec, m.tokenGetter, m.client, m.errChan, m.l)
	return es.get, nil
}

func (m *manager) SetDefaultGoogleCredentials(ctx context.Context, path, key string) error {
	m.l.Print("setting default google credentials")

	s, err := m.GetSecret(ctx, path)
	if err != nil {
		return err
	}

	sm := s()
	if _, ok := sm[key]; !ok {
		return fmt.Errorf("key %s not found in secret", key)
	}
	var encoded string
	var ok bool
	if encoded, ok = sm[key].(string); !ok {
		return fmt.Errorf("key %s is not a string", key)
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	fn := "google-credentials.json"
	if err := os.WriteFile(fn, decoded, 0644); err != nil {
		return err
	}

	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", fn); err != nil {
		return err
	}

	m.l.Printf("set default google credentials to %s", fn)

	return nil
}

func get(ctx context.Context, path, vaultAddress, token string, client *http.Client, l *log.Logger) (*secret, error) {
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	_, span := tracer.Start(
		ctx,
		"hashivault.get",
		trace.WithAttributes(attribute.String("path", path), attribute.String("vaultAddress", vaultAddress)))
	defer span.End()

	url := makeURL(vaultAddress, path)
	l.Printf("getting secrets from %s", url)

	req, err := secretsReq(url, token)
	if err != nil {
		traceError(span, err, l)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		traceError(span, err, l)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		traceError(span, err, l)
		return nil, err
	}

	var sec secret
	err = json.Unmarshal(body, &sec)
	if err != nil {
		traceError(span, err, l)
		return nil, err
	}

	l.Printf("got secrets from %s", url)
	return &sec, nil
}

// makeURL returns a correctly formatted url for Vault http requests
func makeURL(address, path string) string {
	return address + "/v1/" + path
}

// secretsReq returns a http request for getting secrets from Vault
func secretsReq(url, auth string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("while building http request: %w", err)
	}

	req.Header.Set("X-Vault-Token", auth)

	return req, nil
}
