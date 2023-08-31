package cert

import (
	"testing"

	"github.com/3lvia/elvid-go/testing/assert"
)

const (
	certFile = "../../testing/files/test_cert.cer"
	pemFile  = "../../testing/files/test_cert.pem"
)

func Test_MakePool(t *testing.T) {
	// Test with no file
	pool, err := MakePool()
	assert.NoErr(t, err)

	// Test with non-existing file
	pool, err = MakePool("non_existent")
	assert.Err(t, err, ErrLoadingCAfromFile)

	// Test with existing file in wrong format
	pool, err = MakePool(certFile)
	assert.Err(t, err, ErrFailedParsePEM)

	// Test with valid, outdated, certificate from golang.org
	pool, err = MakePool(pemFile)
	assert.NoErr(t, err)

	if pool == nil {
		t.Fatal("Pool is nil")
	}
}