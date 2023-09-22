//go:generate $GOPATH/bin/stringer -type=Type
package mschema

import (
	"context"
)

// Type specifies how the schema is encoded.
type Type int

const (
	JSON Type = iota
	PROTOBUF
	AVRO
)

// TypeName returns the name of the schema type.
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

// Registry is an interface that describes a schema registry. This is the main abstraction of this package.
type Registry interface {
	// GetByID returns a descriptor for the schema with the given ID.
	GetByID(ctx context.Context, id int) (Descriptor, error)

	// GetBySubject returns a descriptor for the latest version of the schema with the given subject.
	GetBySubject(ctx context.Context, subject string) (Descriptor, error)

	// List returns a list of all descriptors in the registry.
	List(ctx context.Context) ([]Descriptor, error)
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

	// Version returns the version of the schema.
	Version() int
}

type descriptor struct {
	Subj      string `json:"subject"`
	V         int    `json:"version"`
	I         int    `json:"id"`
	S         string `json:"schema"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	T         string `json:"schemaType"`
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

func (d descriptor) Version() int {
	return d.V
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
}
