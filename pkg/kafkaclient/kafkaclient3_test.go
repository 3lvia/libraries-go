package kafkaclient

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"sync"
	"testing"
)

const (
	confluentKey    = "XX"
	confluentSecret = "XX"
	boostrapServer  = "pkc-lq8gm.westeurope.azure.confluent.cloud:9092"
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
	application := "democonsumer-1"

	stream, closer, err := StartConsumer(ctx, system, topic, application, WithSecretsManager(v), WithAPIKey(confluentKey, confluentSecret))
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	wg := sync.WaitGroup{}
	wg.Add(100)

	go func(ch <-chan *StreamingMessage, wg *sync.WaitGroup) {
		msg := <-ch
		x := msg
		fmt.Printf("%s\n", x.String)
		wg.Done()
	}(stream, &wg)

	wg.Wait()
}
