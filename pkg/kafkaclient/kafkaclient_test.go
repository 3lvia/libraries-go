package kafkaclient

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"os"
	"testing"
)

func Demo() {

}

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

	stream, err := StartConsumer(ctx, system, topic, application, WithSecretsManager(v))

	msg := <-stream
	_ = msg
}
