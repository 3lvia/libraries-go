package kafkaclientfranz

import (
	"context"
	"encoding/binary"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
	"strings"
)

func newIterator(iter *kgo.FetchesRecordIter, creator EntityCreatorFunc, registry mschema.Registry) StreamingMessageIterator {
	return &messageIterator{
		iter:     iter,
		creator:  creator,
		registry: registry,
	}
}

type messageIterator struct {
	iter     *kgo.FetchesRecordIter
	creator  EntityCreatorFunc
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
		//i.tw.IncFakeMessages(ctx, 1)
		return nil
	}

	b := record.Value

	schemaID := 0
	var err error
	id := b[1:5]
	schemaID = int(binary.BigEndian.Uint32(id))

	d, err := i.registry.GetByID(ctx, schemaID)

	v, err := i.creator(b[5:], d)

	return &StreamingMessage{
		Key:      record.Key,
		SchemaID: schemaID,
		Value:    v,
		Headers:  headers,
		Error:    err,
	}
}
