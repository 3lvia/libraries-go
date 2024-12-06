package kafkaclient

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/3lvia/libraries-go/pkg/hashivault"
)

func TestStartConsumer(t *testing.T) {
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

	system := "core"
	topic := "private.dp.edna.examples"
	application := "democonsumer-2"

	opts := []Option{
		WithSecretsResolver(K8sSecrets{secrets: v}),
	}

	stream, err := StartConsumer(ctx, system, topic, application, opts...)
	if err != nil {
		log.Fatal(err)
	}

	msg := <-stream
	_ = msg
}