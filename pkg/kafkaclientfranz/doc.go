// Package kafkaclientfranz provides a way to communicate with Kafka "the Elvia way", meaning that all the mechanics of
// negotiating with the Kafka cluster, fetching schemas from the schema registry, and decoding messages are handled
// automatically. The only thing the user needs to do is to provide a function that can create the entities that the
// messages are decoded into (if not opting for the defaults implemented internally).
//
// EntityCreatorFunc
// Instances of StreamingMessage that are 'streamed' out of this package on the channel, have a property 'Value' that
// contains the decoded message. How is the message decoded? This is done by the EntityCreatorFunc that can be
// provided by the client via an option. If no EntityCreatorFunc is provided, the method that is used to decode the
// message is determined by the options UseAVRO, UseJSON and UseProtobuf. Currently, only UseAVRO is supported, the
// other two methods have no effect (yet).
//
// If UseAVRO is used (and no EntityCreatorFunc is provided), the message is decoded using the schema that is provided
// with the message. The actual schema is then fetched from the schema registry via the mschema library. The
// functionality of the package 'github.com/linkedin/goavro/v2' is used to decode the message dynamically, returning
// a map[string]interface{} (which probably will be nested for each contained value).
//
// If no EntityCreatorFunc is provided, and UseAVRO is not used, the message is returned as is, i.e. the raw byte
// slice from the Kafka message.
//
// Example using no specified decoding, i.e. returning raw bytes:
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
// Example using specified EntityCreatorFunc. In this case, model objects that have been created using 'schema-cli' is
// used to decode the message.
// ```
//
//	creator := func(data []byte, d mschema.Descriptor) (any, error) {
//	   p, err := model.DeserializePerson(bytes.NewReader(data))
//	   if err != nil {
//	     return nil, err
//	   }
//	   return p, nil
//	}
//
// v, errChan, err := hashivault.New(hashivault.WithOIDC(), hashivault.WithVaultAddress("https://vault.dev-elvia.io"), kafkaclientfranz.WithEntityCreatorFunc(creator))
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
//	creator := func(data []byte, d mschema.Descriptor) (any, error) {
//	  p, err := model.DeserializePerson(bytes.NewReader(data))
//			if err != nil {
//				return nil, err
//			}
//			return p, nil
//		}
//
// stream, err := StartConsumer(
//
//	         ctx,
//	         system,
//	         topic,
//	         application,
//	         kafkaclientfranz.WithSecretsManager(v),
//	         kafkaclientfranz.WithEntityCreatorFunc(creator),
//	)
//
// msg := <-stream // here you would have a for-loop handling the messages from the stream forever.
// ```
package kafkaclientfranz
