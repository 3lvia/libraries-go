package kafkaclient

import (
	"github.com/3lvia/libraries-go/pkg/mschema"
	"time"
)

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

	// Timestamp is the timestamp of the message as it was stored in Kafka.
	Timestamp time.Time

	// String is a string representation of the message on the format "TOPIC[PARTITION]@OFFSET"
	String string

	// IsFake is true if the message is a fake message, i.e. a message that comes from the topic, but that is marked as
	// fake by the producer.
	IsFake bool
}

// EntityCreatorFunc is a function that creates an entity from a byte array. It is the responsibility
// of the consumer of this package to provide an implementation of this function.
type EntityCreatorFunc func(value []byte, d mschema.Descriptor) (any, error)

type apiKey struct {
	key, secret string
}
