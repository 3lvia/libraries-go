package kafkaclient2

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"testing"
)

func TestStartConsumer(t *testing.T) {
	ctx := context.Background()

	opts := []hashivault.Option{
		hashivault.WithVaultAddress("https://vault.dev-elvia.io"),
		hashivault.WithOIDC(),
	}

	v, errChan, err := hashivault.New(ctx, opts...)
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

	_, err = StartConsumer(ctx, system, topic, application, WithSecretsManager(v))
	if err != nil {
		log.Fatal(err)
	}
}
