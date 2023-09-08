package kafkaclient

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
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

	if err := Init(ctx, "edna", "private.dp.edna.examples", WithSecretsManager(v)); err != nil {
		t.Fatal(err)
	}
}