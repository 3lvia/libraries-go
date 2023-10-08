package kafkaclient

import (
	"context"
	"encoding/binary"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
)

func startConsumer(ctx context.Context, topic, consumerGroup string, registry mschema.Registry, format mschema.Type, c kafka.ConfigMap) (<-chan *StreamingMessage, func() error, error) {
	c["group.id"] = consumerGroup
	c["auto.offset.reset"] = "earliest"
	kafkaConsumer, err := kafka.NewConsumer(&c)
	if err != nil {
		return nil, nil, err
	}

	if err = kafkaConsumer.SubscribeTopics([]string{topic}, nil); err != nil {
		return nil, nil, err
	}

	worker := &consumer{
		registry:      registry,
		format:        format,
		kafkaConsumer: kafkaConsumer,
	}

	ch := make(chan *StreamingMessage)
	go worker.consume(ctx, ch)

	return ch, kafkaConsumer.Close, nil
}

type consumer struct {
	registry      mschema.Registry
	format        mschema.Type
	kafkaConsumer *kafka.Consumer
}

func (c *consumer) consume(ctx context.Context, output chan<- *StreamingMessage) {
	for {
		msg, err := c.kafkaConsumer.ReadMessage(-1) // TODO: Make configurable
		if err != nil {
			output <- &StreamingMessage{Error: err}
			continue
		}

		sm := streamingMessage(ctx, msg, c.registry)
		if sm != nil {
			output <- sm
		}
	}
}

func streamingMessage(ctx context.Context, msg *kafka.Message, registry mschema.Registry) *StreamingMessage {
	if msg == nil {
		return nil
	}

	headers := map[string]string{}
	for _, header := range msg.Headers {
		headers[header.Key] = string(header.Value)
	}

	if f, ok := headers["IsFake"]; ok && strings.ToLower(f) == "true" {
		//i.tw.IncFakeMessages(ctx, 1)
		return nil
	}

	b := msg.Value

	schemaID := 0
	id := b[1:5]
	schemaID = int(binary.BigEndian.Uint32(id))

	d, err := registry.GetByID(ctx, schemaID)
	_ = d

	return &StreamingMessage{
		Key:       msg.Key,
		SchemaID:  schemaID,
		Value:     b,
		Headers:   headers,
		Error:     err,
		Timestamp: msg.Timestamp,
		String:    msg.String(),
	}
}
