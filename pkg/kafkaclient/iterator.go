package kafkaclient

//
//func newIterator(iter *kgo.FetchesRecordIter, creator EntityCreatorFunc, format schemaclient.SchemaType) StreamingMessageIterator {
//	return &streamingMessageIterator{
//		iter:    iter,
//		creator: creator,
//		format:  format,
//	}
//}
//
//type streamingMessageIterator struct {
//	iter    *kgo.FetchesRecordIter
//	creator EntityCreatorFunc
//	//format  schemaclient.SchemaType
//	//tw      telemetry.Writer
//}
//
//func (i *streamingMessageIterator) Done() bool {
//	return i.iter.Done()
//}
//
//func (i *streamingMessageIterator) Next(ctx context.Context) *StreamingMessage {
//	record := i.iter.Next()
//
//	headers := map[string]string{}
//	for _, header := range record.Headers {
//		headers[header.Key] = string(header.Value)
//	}
//
//	if f, ok := headers["IsFake"]; ok && strings.ToLower(f) == "true" {
//		//i.tw.IncFakeMessages(ctx, 1)
//		return nil
//	}
//
//	b := record.Value
//	schemaID := 0
//	var d schema.Descriptor
//	s := ""
//	var err error
//	if i.format == messageFormatConfluentAvro {
//		id := b[1:5]
//		schemaID = int(binary.BigEndian.Uint32(id))
//
//		d, err = schema.GetById(schemaID)
//		if d != nil {
//			s = d.Schema()
//		}
//
//		b = b[5:]
//	}
//
//	return &StreamingMessage{
//		Key:      record.Key,
//		SchemaID: schemaID,
//		Value:    i.creator(b, s),
//		Headers:  headers,
//		Error:    err,
//	}
//}
