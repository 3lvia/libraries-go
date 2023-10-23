package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

func authGitHub(ctx context.Context, vaultAddr, githubToken string, client *http.Client) (AuthenticationResponse, error) {
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	_, span := tracer.Start(ctx, "auth.authGitHub", trace.WithAttributes(attribute.String("vault_addr", vaultAddr)))
	defer span.End()

	path := "auth/github/login"
	requestBody, err := githubLogin(githubToken)
	if err != nil {
		return nil, err
	}

	req, err := authReq(vaultAddr, path, requestBody)
	if err != nil {
		return nil, fmt.Errorf("while building http request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("while sending http request: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("while reading response body: %w", err)
	}
	defer resp.Body.Close()

	var response authenticationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("while unmarshalling response body: %w", err)
	}

	return response, nil
}

// githubLogin handles converting the github token string to a bytes buffer
func githubLogin(login string) (*bytes.Buffer, error) {
	return loginBuffer(&gitToken{
		Token: login,
	})
}
