package kafkaclient

import (
	"context"
	"encoding/binary"
	"strings"

	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
)

func newIterator(iter *kgo.FetchesRecordIter, creator EntityCreatorFunc, format mschema.Type, registry mschema.Registry) StreamingMessageIterator {
	return &messageIterator{
		iter:     iter,
		creator:  creator,
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
	var d mschema.Descriptor

	var err error
	if i.format == mschema.AVRO {
		id := b[1:5]

		schemaID = int(binary.BigEndian.Uint32(id))
		d, err = i.registry.GetByID(ctx, schemaID)

		b = b[5:]
	}

	v, err := i.creator(b, d)

	return &StreamingMessage{
		Key:      record.Key,
		SchemaID: schemaID,
		Value:    v,
		Headers:  headers,
		Error:    err,
	}
}