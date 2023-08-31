package assert

import (
	"errors"
	"testing"
)

// NoErr if you don't expect an error
func NoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

// Err if you expect an error
func Err(t *testing.T, err error, want error) {
	t.Helper()

	if !errors.Is(err, want) {
		t.Fatalf("Expected error: %v, got: %v", want, err)
	}
}