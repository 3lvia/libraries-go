package kafkaclient

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"log"
	"sync"
	"testing"
	"time"
)

const (
	confluentKey    = "PVTYEHIRJ46T27H5"
	confluentSecret = "TrLS8343q9M+9wOUEPGPL671JeJgtkROALGqZ22cZ4pawzAzffk05aNARyMEb8Of"
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

	stream, closer, err := StartConsumer(
		ctx,
		system,
		topic,
		application,
		WithSecretsManager(v),
		WithAPIKey(confluentKey, confluentSecret),
		ReturnFakes(),
		AutoOffsetResetEarliest(),
		WithOffsetTime(time.Now().Add(-12*time.Hour)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	wg := sync.WaitGroup{}
	wg.Add(10)

	go func(ch <-chan *StreamingMessage, wg *sync.WaitGroup) {
		for {
			msg := <-ch
			x := msg
			fmt.Printf("%s (%s) fake: %v\n", x.String, x.Timestamp.Format("2006-01-02 15:04:05.000"), x.IsFake)

			wg.Done()
		}
	}(stream, &wg)

	wg.Wait()
}
