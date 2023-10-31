package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/linkedin/goavro/v2"
	"net/http"
	"time"
)

func StartConsumer(ctx context.Context, system, topic, consumerGroup string, opts ...Option) (<-chan *StreamingMessage, func() error, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	if collector.secrets == nil {
		return nil, nil, errors.New("no secrets manager provided")
	}

	client := collector.client
	if client == nil {
		client = &http.Client{}
	}

	secrets, err := getSecrets(ctx, system, collector.secrets, collector.key)
	if err != nil {
		return nil, nil, err
	}

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, nil, err
	}

	creator := creatorFunc(collector)

	c := config(secrets, collector, consumerGroup)

	kafkaConsumer, err := kafka.NewConsumer(&c)
	if err != nil {
		return nil, nil, err
	}

	offsets, err := getTimeOffsets(topic, kafkaConsumer, collector)

	ch, closer, err := startConsumer(ctx, topic, creator, registry, collector.returnFakes, kafkaConsumer, offsets)
	if err != nil {
		return nil, nil, err
	}
	return ch, closer, nil
}

func getTimeOffsets(topic string, c *kafka.Consumer, collector *optionsCollector) ([]kafka.TopicPartition, error) {
	if !collector.offsetTimeSet {

		return nil, nil
	}

	timestamp := collector.offsetTime.UnixNano() / int64(time.Millisecond)
	times := []kafka.TopicPartition{{Topic: &topic, Offset: kafka.Offset(timestamp)}}
	partitionOffsets, err := c.OffsetsForTimes(times, 5000)
	if err != nil {
		return nil, err
	}

	return partitionOffsets, nil
}

func kafkaConsumer(secrets *secretConfigValues, collector *optionsCollector, topic, consumerGroup string) (*kafka.Consumer, []kafka.TopicPartition, error) {
	c := config(secrets, collector, consumerGroup)

	kc, err := kafka.NewConsumer(&c)
	if err != nil {
		return nil, nil, err
	}

	var offsets []kafka.TopicPartition
	if collector.offsetTimeSet {
		timestamp := collector.offsetTime.UnixNano() / int64(time.Millisecond)
		times := []kafka.TopicPartition{{Topic: &topic, Offset: kafka.Offset(timestamp)}}
		partitionOffsets, err := kc.OffsetsForTimes(times, 5000)
		if err != nil {
			return nil, nil, err
		}
		offsets = partitionOffsets
	} else {
		adminClient, err := kafka.NewAdminClientFromConsumer(kc)
		if err != nil {
			return nil, nil, err
		}

		metadata, err := adminClient.GetMetadata(&topic, false, 5000)
		if err != nil {
			return nil, nil, err
		}

		for i := 0; i < len(metadata.Topics[topic].Partitions); i++ {
			tp := kafka.TopicPartition{
				Topic:     &topic,
				Partition: i,
				Offset:    kafka.OffsetEnd,
			}
			offsets = append(offsets, tp)
		}
	}

	//offsets, err := getTimeOffsets(secrets.topic, kafkaConsumer, collector)
	//if err != nil {
	//	return nil, nil, err
	//}

	return kc, offsets, nil
}

func config(secrets *secretConfigValues, collector *optionsCollector, consumerGroup string) kafka.ConfigMap {
	var offsetReset string
	switch collector.autoOffsetReset {
	case autoOffsetResetEarliest:
		offsetReset = "earliest"
	case autoOffsetResetLateset:
		offsetReset = "latest"
	default:
		offsetReset = "none"
	}

	return map[string]kafka.ConfigValue{
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"bootstrap.servers": secrets.bootstrapServer,
		"sasl.username":     secrets.key,
		"sasl.password":     secrets.secret,
		"auto.offset.reset": offsetReset,
		"group.id":          consumerGroup,
		//"acks":              "all",
	}
}

func creatorFunc(collector *optionsCollector) EntityCreatorFunc {
	creator := collector.creatorFunc
	if creator != nil {
		// If the client has provided a custom creator function, we use that, regardless of the format.
		return creator
	}
	if collector.formatSet && collector.format == mschema.AVRO {
		// Since collector.creatorFunc == nil, and the schema is AVRO, we use dynamic typing.
		return defaultAVROCreator
	}

	// The default creator function simply returns the raw byte slice from the Kafka message.
	return func(data []byte, d mschema.Descriptor) (any, error) {
		return data, nil
	}
}

func defaultAVROCreator(value []byte, d mschema.Descriptor) (any, error) {
	codec, err := goavro.NewCodec(d.Schema())
	if err != nil {
		return nil, err
	}
	obj, _, err := codec.NativeFromBinary(value)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
