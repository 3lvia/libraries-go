package tlsclient

import (
	"crypto/tls"
	"net/http"

	"github.com/3lvia/elvid-go/internal/cert"
)

// New creates a new instance of http.Client that is configured to use tls.VersionTLS12 as minimum
// and a cert.Pool based on Let's Encrypt public root level certificate and the passed in certificates.
func New(certificates ...string) (*http.Client, error) {

	pool, err := cert.MakePool(certificates...)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool.Certs,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	return httpClient, nil
}