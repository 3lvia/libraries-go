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

type reader interface {
	start(ctx context.Context, output chan<- *StreamingMessage)
}

func newReaderFromFetcher(fetcher StreamingMessageFetcher, format mschema.Type, entityCreator EntityCreatorFunc) reader {
	return &readerFranz{
		fetcher:       fetcher,
		entityCreator: entityCreator,
		format:        format,
	}
}

func newReader(
	topic,
	broker,
	groupID,
	clientID,
	userName,
	password string,
	format mschema.Type,
	entityCreator EntityCreatorFunc,
	registry mschema.Registry) (reader, error) {
	seeds := []string{broker}

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

		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(topic),
		kgo.ClientID(clientID),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	f := &kafkaMessageFetcher{client: client, registry: registry}
	r := &readerFranz{fetcher: f, entityCreator: entityCreator, format: format}
	return r, nil
}

type readerFranz struct {
	fetcher       StreamingMessageFetcher
	entityCreator EntityCreatorFunc
	format        mschema.Type
}

func (r *readerFranz) start(ctx context.Context, output chan<- *StreamingMessage) {
	defer r.fetcher.Close()

	for {
		iter, err := r.fetcher.PollFetches(ctx, r.format, r.entityCreator)
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
