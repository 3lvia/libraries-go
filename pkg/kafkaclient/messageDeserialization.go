package kafkaclient

import (
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/linkedin/goavro/v2"
)

func creatorFunc(format mschema.Type, collector *optionsCollector) EntityCreatorFunc {
	creator := collector.creatorFunc
	if creator != nil {
		return creator
	}
	if format == mschema.AVRO {
		// Since collector.creatorFunc == nil, and the schema is AVRO, we use dynamic typing.
		return defaultAVROCreator
	}

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
