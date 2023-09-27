package main

import (
	"bytes"
	"context"
	"github.com/3lvia/libraries-go/cmd/test-client/model"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/kafkaclient"
	"github.com/actgardner/gogen-avro/v10/soe"
	"github.com/linkedin/goavro/v2"
	"log"
	"os"
	"sync"
)

func main() {
	topic()
	//write()

	//fn := `/Users/staleheitmann/go/src/github.com/3lvia/libraries-go/cmd/test-client/message.byt`
	//if err := goAvro(fn); err != nil {
	//	log.Fatal(err)
	//}

	//if err := inMemory(); err != nil {
	//	log.Fatal(err)
	//}

	//files := []string{`/Users/staleheitmann/go/src/github.com/3lvia/libraries-go/cmd/test-client/message.byt`, `/Users/staleheitmann/go/src/github.com/3lvia/libraries-go/cmd/test-client/reg.byt`}
	//for _, file := range files {
	//	fmt.Println(file)
	//	err := readFile(file)
	//	if err != nil {
	//		fmt.Printf("Error reading file %s: %s\n", file, err)
	//	}
	//}
}

func goAvro(fn string) error {
	schemaBytes, err := os.ReadFile("./100172_6.avsc")
	if err != nil {
		return err
	}
	codec, err := goavro.NewCodec(string(schemaBytes))
	if err != nil {
		return err
	}

	messageBytes, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	obj, bb, err := codec.NativeFromBinary(messageBytes[5:])
	if err != nil {
		return err
	}
	_ = obj
	_ = bb

	return nil
}

func inMemory() error {
	p := model.NewPerson()
	p.Id = &model.UnionNullString{String: "123", UnionType: model.UnionNullStringTypeEnumString}
	p.Name = &model.UnionNullString{String: "Ståle", UnionType: model.UnionNullStringTypeEnumString}

	var buf bytes.Buffer
	if err := soe.WriteRecord(&buf, p); err != nil {
		return err
	}

	r := bytes.NewReader(buf.Bytes())
	_, err := model.DeserializePerson(r)
	if err != nil {
		return err
	}
	return nil
}

func write() {

	p := model.NewPerson()
	p.Id = &model.UnionNullString{String: "123", UnionType: model.UnionNullStringTypeEnumString}
	p.Name = &model.UnionNullString{String: "Ståle", UnionType: model.UnionNullStringTypeEnumString}

	fo, err := os.Create("soe.byt")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	if err := soe.WriteRecord(fo, p); err != nil {
		log.Fatal(err)
	}

	fr, err := os.Create("reg.byt")
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	if err := p.Serialize(fr); err != nil {
		log.Fatal(err)
	}

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

	//creator := func(data []byte, schemaID int) (any, error) {
	//	p, err := model.DeserializePerson(bytes.NewReader(data))
	//	if err != nil {
	//		return nil, err
	//	}
	//	return p, nil
	//}

	go func(ec <-chan error) {
		for err := range ec {
			log.Println(err)
		}
	}(errChan)

	system := "core"
	topic := "private.dp.edna.examples"
	application := "democonsumer-2"

	stream, err := kafkaclient.StartConsumer(ctx, system, topic, application, kafkaclient.WithSecretsManager(v)) // kafkaclient.WithEntityCreatorFunc(creator)

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

func readFile(fn string) error {
	b, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	r := bytes.NewReader(b[5:])
	_, err = model.DeserializePerson(r)
	if err != nil {
		return err
	}

	return nil
}
