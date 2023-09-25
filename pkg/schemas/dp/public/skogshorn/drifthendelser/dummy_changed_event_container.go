// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100126_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewDummyChangedEventWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewDummyChangedEvent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DummyChangedEventReader struct {
	r io.Reader
	p *vm.Program
}

func NewDummyChangedEventReader(r io.Reader) (*DummyChangedEventReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewDummyChangedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DummyChangedEventReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r DummyChangedEventReader) Read() (DummyChangedEvent, error) {
	t := NewDummyChangedEvent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
