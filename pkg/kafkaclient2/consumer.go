package kafkaclient2

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"log"
)

func consume(ctx context.Context, topic, consumerGroup string, partition int, secrets *secretConfigValues) {
	dialer := &kafka.Dialer{
		SASLMechanism: plain.Mechanism{
			Username: secrets.key,
			Password: secrets.secret,
		},
		DualStack: true,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{secrets.bootstrapServer},
		Dialer:  dialer,
		Topic:   topic,
		GroupID: consumerGroup,
	})

	for {
		m, err := reader.ReadMessage(ctx)
		_ = m
		if err != nil {
			log.Fatal(err)
		}
	}
}
