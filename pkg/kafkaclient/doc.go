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
//
// CONSUMING MESSAGES
// The following configuration values are needed, and on the basis of these, the correct configuration will be fecthed
// from Vault, and the kafka client will be configured:
// * system. This is the name of the system which the the consuming application is a part of.
// * topic. This is the name of the topic that the application is consuming from.
// * application. This is the name of the application that is consuming messages.
package kafkaclient
