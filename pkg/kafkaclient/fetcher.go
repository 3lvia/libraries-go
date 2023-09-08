package kafkaclient

//type kafkaMessageFetcher struct {
//	client *kgo.Client
//}
//
//func (f *kafkaMessageFetcher) Close() {
//	f.client.Close()
//}
//
//func (f *kafkaMessageFetcher) PollFetches(ctx context.Context, format schemaclient.SchemaType, creator EntityCreatorFunc) (StreamingMessageIterator, error) {
//	fetches := f.client.PollFetches(ctx)
//	if errs := fetches.Errors(); len(errs) > 0 {
//		// All errors are retried internally when fetching, but non-retriable errors are
//		// returned from polls so that users can notice and take action.
//		msg := fmt.Sprint(errs)
//		return nil, errors.New(msg)
//	}
//
//	fetches.EachTopic(func(ft kgo.FetchTopic) {
//		for _, partition := range ft.Partitions {
//			fmt.Printf("topic %s partition %d high watermark %d\n", ft.Topic, partition.Partition, partition.HighWatermark)
//		}
//	})
//
//	it := newIterator(fetches.RecordIter(), creator, format, f.tw)
//	return it, nil
//}
