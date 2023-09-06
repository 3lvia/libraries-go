package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate_github(t *testing.T) {
	ctx := context.Background()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, ghVaultResponse)
	}))
	defer testServer.Close()

	ghToken := "MY_GITHUB_TOKEN"
	tokenResponse, err := Authenticate(ctx, testServer.URL, MethodGitHub, WithGitHubToken(ghToken), WithClient(testServer.Client()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tokenResponse.ClientToken() != "xxx" {
		t.Errorf("unexpected token: %s", tokenResponse.ClientToken())
	}
}

func TestAuthenticate_k8s(t *testing.T) {
	ctx := context.Background()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, ghVaultResponse)
	}))
	defer testServer.Close()

	servicePath := "MY_SERVICE_PATH"
	role := "MY_ROLE"
	tokenResponse, err := Authenticate(ctx, testServer.URL, MethodK8s, WithK8s(servicePath, role), WithClient(testServer.Client()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if tokenResponse.ClientToken() != "xxx" {
		t.Errorf("unexpected token: %s", tokenResponse.ClientToken())
	}
}

const ghVaultResponse = `{
    "request_id": "d645ddd7-3b2e-f28b-0138-512d5ff301a4",
    "lease_id": "",
    "renewable": false,
    "lease_duration": 0,
    "data": null,
    "wrap_info": null,
    "warnings": null,
    "auth": {
        "client_token": "xxx",
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
			"lease_duration": 2764800,
        "renewable": true,
        "entity_id": "2957867d-00ea-981d-e2e9-dd41ab412212",
        "token_type": "service",
        "orphan": true,
        "mfa_requirement": null,
        "num_uses": 0
    }
}`
