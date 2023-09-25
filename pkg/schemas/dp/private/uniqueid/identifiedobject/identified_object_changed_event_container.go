// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100254_5.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewIdentifiedObjectChangedEventWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewIdentifiedObjectChangedEvent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type IdentifiedObjectChangedEventReader struct {
	r io.Reader
	p *vm.Program
}

func NewIdentifiedObjectChangedEventReader(r io.Reader) (*IdentifiedObjectChangedEventReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewIdentifiedObjectChangedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &IdentifiedObjectChangedEventReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r IdentifiedObjectChangedEventReader) Read() (IdentifiedObjectChangedEvent, error) {
	t := NewIdentifiedObjectChangedEvent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
