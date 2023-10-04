package kafkaclient

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
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
	application string,
	info kafkaClientInfo,
	format mschema.Type,
	entityCreator EntityCreatorFunc,
	registry mschema.Registry,
	offsetSender chan<- OffsetInfo) (consumer, error) {

	consumerGroup := fmt.Sprintf("%s-%s.%s", topic, system, application)
	client, err := fetchClient(topic, consumerGroup, info)
	if err != nil {
		return nil, err
	}

	//admClient := kadm.NewClient(client)
	//admClient.

	//client.SetOffsets()

	f := newFetcher(client, registry, offsetSender)
	r := &consumerFranz{fetcher: f, entityCreator: entityCreator, format: format}
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
