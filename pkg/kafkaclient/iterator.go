package kafkaclient

import (
	"context"
	"encoding/binary"
	"errors"
	"strings"

	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	ErrNoDescriptor = errors.New("could not get the descriptor")
)

func newIterator(iter *kgo.FetchesRecordIter, creator EntityCreatorFunc, format mschema.Type, registry mschema.Registry) StreamingMessageIterator {
	return &messageIterator{
		iter:     iter,
		creator:  creator,
		format: format,
		registry: registry,
	}
}

type messageIterator struct {
	iter     *kgo.FetchesRecordIter
	creator  EntityCreatorFunc
	format   mschema.Type
	registry mschema.Registry
}

func (i *messageIterator) Done() bool {
	return i.iter.Done()
}

func (i *messageIterator) Next(ctx context.Context) *StreamingMessage {
	record := i.iter.Next()

	headers := map[string]string{}
	for _, header := range record.Headers {
		headers[header.Key] = string(header.Value)
	}

	if f, ok := headers["IsFake"]; ok && strings.ToLower(f) == "true" {
		// i.tw.IncFakeMessages(ctx, 1)
		return nil
	}

	b := record.Value

	schemaID := 0
	var v any
	var err error

	switch i.format {
	case mschema.AVRO:
		id := b[1:5]

		schemaID = int(binary.BigEndian.Uint32(id))

		var d mschema.Descriptor
		d, err = i.registry.GetByID(ctx, schemaID)

		if err != nil {
			err = errors.Join(ErrNoDescriptor, err)
		} else {
			v, err = i.creator(b[5:], d)
		}
	default:
		var d mschema.Descriptor
		v, err = i.creator(b, d)
	}

	return &StreamingMessage{
		Key:      record.Key,
		SchemaID: schemaID,
		Value:    v,
		Headers:  headers,
		Error:    err,
	}
}
