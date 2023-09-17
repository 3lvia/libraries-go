//go:generate $GOPATH/bin/stringer -type=Type
package mschema

import (
	"context"
	"fmt"
	"strings"
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
	GetByID(ctx context.Context, id int) (Descriptor, error)
	GetBySubject(ctx context.Context, subject string) (Descriptor, error)
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

	// GenerationFolder returns the folder where types that are generated from this schema should be stored.
	GenerationFolder() (string, error)
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

func (d descriptor) GenerationFolder() (string, error) {
	arr := strings.Split(d.Subj, ".")
	if len(arr) != 4 {
		return "", fmt.Errorf("invalid subject: %s", d.Subj)
	}
	typ := arr[len(arr)-1]
	if strings.Contains(typ, "-value") {
		typ = strings.Replace(typ, "-value", "", 1)
	}
	return fmt.Sprintf("%s/%s/%s/%s/", arr[1], arr[2], typ, arr[0]), nil
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
