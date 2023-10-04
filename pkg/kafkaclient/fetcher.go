package kafkaclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
)

func newFetcher(client *kgo.Client, registry mschema.Registry, offsetSender chan<- OffsetInfo) StreamingMessageFetcher {
	return &kafkaMessageFetcher{
		client:       client,
		registry:     registry,
		offsetSender: offsetSender,
	}
}

type kafkaMessageFetcher struct {
	client       *kgo.Client
	registry     mschema.Registry
	offsetSender chan<- OffsetInfo
}

func (f *kafkaMessageFetcher) Close() {
	f.client.Close()
}

func (f *kafkaMessageFetcher) PollFetches(ctx context.Context, format mschema.Type, creator EntityCreatorFunc) (StreamingMessageIterator, error) {
	fetches := f.client.PollFetches(ctx)
	if errs := fetches.Errors(); len(errs) > 0 {
		// All errors are retried internally when fetching, but non-retriable errors are
		// returned from polls so that users can notice and take action.
		msg := fmt.Sprint(errs)
		return nil, errors.New(msg)
	}

	if f.offsetSender != nil {
		fetches.EachTopic(func(ft kgo.FetchTopic) {
			for _, partition := range ft.Partitions {
				offset := OffsetInfo{
					Topic:     ft.Topic,
					Partition: partition.Partition,
					Offset:    partition.HighWatermark,
				}
				f.offsetSender <- offset
			}
		})
	}

	it := newIterator(fetches.RecordIter(), creator, f.registry)

	return it, nil
}
