package kafkaclient

import (
	"context"
	"encoding/binary"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
)

func startConsumer(
	ctx context.Context,
	topic, consumerGroup string,
	creator EntityCreatorFunc,
	registry mschema.Registry,
	c kafka.ConfigMap) (<-chan *StreamingMessage, func() error, error) {
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
		kafkaConsumer: kafkaConsumer,
		creator:       creator,
	}

	ch := make(chan *StreamingMessage)
	go worker.consume(ctx, ch)

	return ch, kafkaConsumer.Close, nil
}

type consumer struct {
	registry      mschema.Registry
	kafkaConsumer *kafka.Consumer
	creator       EntityCreatorFunc
}

func (c *consumer) consume(ctx context.Context, output chan<- *StreamingMessage) {
	for {
		msg, err := c.kafkaConsumer.ReadMessage(-1) // TODO: Make configurable
		if err != nil {
			output <- &StreamingMessage{Error: err}
			continue
		}

		sm := streamingMessage(ctx, msg, c.registry, c.creator)
		if sm != nil {
			output <- sm
		}
	}
}

func streamingMessage(ctx context.Context, msg *kafka.Message, registry mschema.Registry, creator EntityCreatorFunc) *StreamingMessage {
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
	if err != nil {
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

	transformed, err := creator(b, d)
	if err != nil {
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

	return &StreamingMessage{
		Key:       msg.Key,
		SchemaID:  schemaID,
		Value:     transformed,
		Headers:   headers,
		Error:     err,
		Timestamp: msg.Timestamp,
		String:    msg.String(),
	}
}
