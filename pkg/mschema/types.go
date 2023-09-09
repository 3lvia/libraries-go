package mschema

import "context"

// Type represents the type of a schema.
type Type int

const (
	JSON Type = iota
	PROTOBUF
	AVRO
)

func TypeName(t Type) string {
	switch t {
	case JSON:
		return "JSON"
	case PROTOBUF:
		return "PROTOBUF"
	case AVRO:
		return "AVRO"
	default:
		return "AVRO"
	}
}

type Registry interface {
	GetByID(ctx context.Context, id int) (Descriptor, error)
	GetBySubject(ctx context.Context, subject string) (Descriptor, error)
}

// Descriptor is a structure that describes a message schema.
type Descriptor interface {
	// Type returns the type of the schema.
	Type() Type

	// ID returns the unique ID of the schema.
	ID() int

	// Schema returns the actual schema as a string.
	Schema() string

	// Subject returns the subject of the schema. For descriptors that originate from the Confluent schema registry,
	// this is based on the topic name.
	Subject() string
}

type descriptor struct {
	Subj      string `json:"subject"`
	Version   int    `json:"version"`
	I         int    `json:"id"`
	S         string `json:"schema"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	T         string `json:"type"`
}

func (d descriptor) Subject() string {
	return d.Subj
}

func (d descriptor) ID() int {
	return d.I
}

func (d descriptor) Schema() string {
	return d.S
}

func (d descriptor) Type() Type {
	switch d.T {
	case "JSON":
		return JSON
	case "PROTOBUF":
		return PROTOBUF
	case "AVRO":
		return AVRO
	default:
		return AVRO
	}

	return -1
}
