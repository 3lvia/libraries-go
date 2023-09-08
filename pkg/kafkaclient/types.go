package kafkaclient

import "context"

// EntityCreatorFunc is a function that creates an entity from a byte array. It is the responsibility
// of the consumer of this package to provide an implementation of this function.
type EntityCreatorFunc func(value []byte, schema string) any

// StreamingMessage represents a message that has been read from Kafka. Som pre-processing has been done
// on the message, such as decoding the message value from Avro.
type StreamingMessage struct {
	// Key is the key of the message as it was stored in Kafka.
	Key []byte

	// SchemaID is the ID of the schema used to encode the message.
	SchemaID int

	// Value is the value of the message as it was stored in Kafka. If a message
	// transformer has been configured, the value will be the transformed value.
	Value interface{}

	// Headers are the headers of the message as it was stored in Kafka.
	Headers map[string]string

	// Error is the error that occurred while processing the message.
	Error error
}

type StreamingMessageIterator interface {
	Done() bool
	Next(ctx context.Context) *StreamingMessage
}
