package schemaclient

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"os"
	"testing"
)

func Test_newRegistry(t *testing.T) {
	ctx := context.Background()

	vaultAddr := "https://vault.dev-elvia.io"
	if err := os.Setenv("VAULT_ADDR", vaultAddr); err != nil {
		t.Fatal(err)
	}

	v, errChan, err := hashivault.New(
		ctx,
		hashivault.WithOIDC(),
		hashivault.WithVaultAddress(vaultAddr),
	)
	if err != nil {
		log.Fatal(err)
	}

	go func(ec <-chan error) {
		for err := range ec {
			log.Println(err)
		}
	}(errChan)

	r, err := New(ctx, "edna", WithSecretsManager(v))
	if err != nil {
		t.Fatal(err)
	}

	topic := "private.dp.edna.examples"

	s, err := r.Get(topic)
	if err != nil {
		t.Fatal(err)
	}

	_ = s
}
