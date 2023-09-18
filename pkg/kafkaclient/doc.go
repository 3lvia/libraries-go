// Package kafkaclient provides a way to communicate with Kafka "the Elvia way", meaning that all the mechanics of
// negotiating with the Kafka cluster, fetching schemas from the schema registry, and decoding messages are handled
// automatically. The only thing the user needs to do is to provide a function that can create the entities that the
// messages are decoded into.
//
// ```
// v, errChan, err := hashivault.New(hashivault.WithOIDC(), hashivault.WithVaultAddress("https://vault.dev-elvia.io"))
//
//	if err != nil {
//	  log.Fatal(err)
//	}
//
//	go func(ec <-chan error) {
//		for err := range ec {
//			log.Println(err)
//		}
//	}(errChan)
//
// system := "core"
// topic := "private.dp.edna.examples"
// application := "democonsumer-2"
//
// stream, err := StartConsumer(ctx, system, topic, application, WithSecretsManager(v))
//
// msg := <-stream // here you would have a for-loop handling the messages from the stream forever.
// ```
package kafkaclient
