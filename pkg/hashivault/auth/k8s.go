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

func authK8s(ctx context.Context, vaultAddr, k8ServicePath, role string, client *http.Client) (AuthenticationResponse, error) {
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	_, span := tracer.Start(
		ctx,
		"auth.authGitHub",
		trace.WithAttributes(
			attribute.String("vault_addr", vaultAddr),
			attribute.String("k8s_service_path", k8ServicePath),
			attribute.String("k8s_role", role),
		))
	defer span.End()

	path := "auth/" + k8ServicePath + "/login"
	requestBody, err := k8sLogin(k8ServicePath, role)
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

// k8sLogin handles converting the service path and role to a bytes buffer
func k8sLogin(k8ServicePath string, role string) (*bytes.Buffer, error) {
	jwt, err := getJWT(k8ServicePath)
	if err != nil {
		return nil, err
	}

	return loginBuffer(&k8sToken{
		JWT:  jwt,
		Role: role,
	})
}
