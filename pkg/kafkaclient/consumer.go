package kafkaclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"net"
	"time"
)

type consumer interface {
	start(ctx context.Context, output chan<- *StreamingMessage)
}

func newConsumerFromFetcher(fetcher StreamingMessageFetcher, format mschema.Type, entityCreator EntityCreatorFunc) consumer {
	return &consumerFranz{
		fetcher:       fetcher,
		entityCreator: entityCreator,
		format:        format,
	}
}

func newConsumer(
	system,
	topic,
	application,
	broker,
	userName,
	password string,
	//format mschema.Type,
	entityCreator EntityCreatorFunc,
	registry mschema.Registry) (consumer, error) {
	seeds := []string{broker}

	consumerGroup := fmt.Sprintf("%s-%s.%s", topic, system, application)
	// TODO: clientID should be unique for each instance of the consumer
	clientID := "libraries-test"

	tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),

		// SASL Options
		kgo.SASL(plain.Auth{
			User: userName,
			Pass: password,
		}.AsMechanism()),

		// Configure TLS. Uses SystemCertPool for RootCAs by default.
		kgo.Dialer(tlsDialer.DialContext),

		kgo.ConsumerGroup(consumerGroup),
		kgo.ConsumeTopics(topic),
		kgo.ClientID(clientID),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	f := &kafkaMessageFetcher{client: client, registry: registry}
	r := &consumerFranz{fetcher: f, entityCreator: entityCreator}
	return r, nil
}

type consumerFranz struct {
	fetcher       StreamingMessageFetcher
	entityCreator EntityCreatorFunc
	format        mschema.Type
}

func (c *consumerFranz) start(ctx context.Context, output chan<- *StreamingMessage) {
	defer c.fetcher.Close()

	for {
		iter, err := c.fetcher.PollFetches(ctx, c.format, c.entityCreator)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// We can iterate through a record iterator...
		for !iter.Done() {
			record := iter.Next(ctx)
			if record != nil {
				// record may be nil if a message on the Kafka topic was marked as fake.
				output <- record
			}
		}
	}
}
