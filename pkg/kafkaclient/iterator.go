package kafkaclient

import (
	"context"
	"encoding/binary"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/twmb/franz-go/pkg/kgo"
	"strings"
)

func newIterator(iter *kgo.FetchesRecordIter, creator EntityCreatorFunc, format mschema.Type, registry mschema.Registry) StreamingMessageIterator {
	return &streamingMessageIterator{
		iter:     iter,
		creator:  creator,
		format:   format,
		registry: registry,
	}
}

type streamingMessageIterator struct {
	iter     *kgo.FetchesRecordIter
	creator  EntityCreatorFunc
	format   mschema.Type
	registry mschema.Registry
}

func (i *streamingMessageIterator) Done() bool {
	return i.iter.Done()
}

func (i *streamingMessageIterator) Next(ctx context.Context) *StreamingMessage {
	record := i.iter.Next()

	headers := map[string]string{}
	for _, header := range record.Headers {
		headers[header.Key] = string(header.Value)
	}

	if f, ok := headers["IsFake"]; ok && strings.ToLower(f) == "true" {
		//i.tw.IncFakeMessages(ctx, 1)
		v, err := i.creator(record.Value[5:], 1)
		_ = err
		_ = v
		return nil
	}

	b := record.Value
	schemaID := 0
	//var d mschema.Descriptor
	//s := ""
	var err error
	id := b[1:5]
	schemaID = int(binary.BigEndian.Uint32(id))

	//d, err = i.registry.GetByID(ctx, schemaID)
	//if d != nil {
	//	s = d.Schema()
	//}

	b = b[5:]

	v, err := i.creator(b, schemaID)

	return &StreamingMessage{
		Key:      record.Key,
		SchemaID: schemaID,
		Value:    v,
		Headers:  headers,
		Error:    err,
	}
}
