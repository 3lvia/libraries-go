package kafkaclient

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
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

	system := "edna"
	topic := "private.dp.edna.examples"
	groupID := "libraries-test"
	clientID := "libraries-test"

	stream, err := Start(ctx, system, topic, groupID, clientID, WithSecretsManager(v))

	msg := <-stream
	_ = msg
}
