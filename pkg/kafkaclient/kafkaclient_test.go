package kafkaclient

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	ctx := context.Background()
	v, errChan, err := hashivault.New(
		ctx,
		hashivault.WithOIDC(),
		hashivault.WithVaultAddress("https://vault.dev-elvia.io"),
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
