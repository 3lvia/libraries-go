package main

import (
	"bytes"
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/kafkaclient"
	model "github.com/3lvia/libraries-go/pkg/schemas/dp/private/edna/examples"
	"log"
	"os"
	"path"
	"sync"
)

func main() {
	topic()

}

func readFile() {
	b, err := os.ReadFile(`C:\Users\wr676\go\src\github.com\3lvia\libraries-go\cmd\test-client\message.byt`)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(b[5:])
	pr, err := model.NewPersonReader(r)
	if err != nil {
		log.Fatal(err)
	}

	_ = pr
}

func topic() {
	ctx := context.Background()

	vaultAddr := "https://vault.dev-elvia.io"
	if err := os.Setenv("VAULT_ADDR", vaultAddr); err != nil {
		log.Fatal(err)
	}

	v, errChan, err := hashivault.New(
		ctx,
		hashivault.WithOIDC(),
		hashivault.WithVaultAddress(vaultAddr),
	)
	if err != nil {
		log.Fatal(err)
	}

	creator := func(data []byte, schemaID int) (any, error) {
		f := path.Join(`C:\Users\wr676\go\src\github.com\3lvia\libraries-go\cmd\test-client`, "message.byt")
		os.Create(f)

		err := os.WriteFile(f, data, 0644)
		if err != nil {
			return nil, err
		}

		r := bytes.NewReader(data)
		pr, err := model.NewPersonReader(r)
		if err != nil {
			return nil, err
		}
		p, err := pr.Read()
		return p, err
	}

	go func(ec <-chan error) {
		for err := range ec {
			log.Println(err)
		}
	}(errChan)

	system := "core"
	topic := "private.dp.edna.examples"
	application := "democonsumer-2"

	stream, err := kafkaclient.StartConsumer(ctx, system, topic, application, kafkaclient.WithSecretsManager(v), kafkaclient.WithEntityCreatorFunc(creator))

	wg := &sync.WaitGroup{}
	wg.Add(100)

	var messages []*kafkaclient.StreamingMessage

	go func(ch <-chan *kafkaclient.StreamingMessage, w *sync.WaitGroup) {
		m := <-ch
		messages = append(messages, m)
		w.Done()
	}(stream, wg)

	wg.Wait()
	l := len(messages)
	_ = l
}
