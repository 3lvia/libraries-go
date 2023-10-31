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
	topic string,
	creator EntityCreatorFunc,
	registry mschema.Registry,
	returnFakes bool,
	kafkaConsumer *kafka.Consumer,
	offsets []kafka.TopicPartition) (<-chan *StreamingMessage, func() error, error) {

	if offsets != nil && len(offsets) > 0 {
		if err := kafkaConsumer.Assign(offsets); err != nil {
			return nil, nil, err
		}
	} else if err := kafkaConsumer.SubscribeTopics([]string{topic}, nil); err != nil {
		return nil, nil, err
	}

	worker := &consumer{
		registry:      registry,
		kafkaConsumer: kafkaConsumer,
		creator:       creator,
		returnFakes:   returnFakes,
	}

	ch := make(chan *StreamingMessage)
	go worker.start(ctx, ch)

	return ch, kafkaConsumer.Close, nil
}

type consumer struct {
	registry      mschema.Registry
	kafkaConsumer *kafka.Consumer
	creator       EntityCreatorFunc
	returnFakes   bool
}

func (c *consumer) start(ctx context.Context, output chan<- *StreamingMessage) {
	for {
		msg, err := c.kafkaConsumer.ReadMessage(-1) // TODO: Make configurable?
		if err != nil {
			output <- &StreamingMessage{Error: err}
			continue
		}

		sm := streamingMessage(ctx, msg, c.registry, c.creator, c.returnFakes)

		if sm != nil {
			output <- sm
		}
	}
}

func streamingMessage(ctx context.Context, msg *kafka.Message, registry mschema.Registry, creator EntityCreatorFunc, returnFakes bool) *StreamingMessage {
	if msg == nil {
		return nil
	}

	headers := map[string]string{}
	for _, header := range msg.Headers {
		headers[header.Key] = string(header.Value)
	}

	isFake := false

	if f, ok := headers["IsFake"]; ok && strings.ToLower(f) == "true" {
		isFake = true
	}

	if isFake && !returnFakes {
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
			IsFake:    isFake,
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
			IsFake:    isFake,
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
		IsFake:    isFake,
	}
}
