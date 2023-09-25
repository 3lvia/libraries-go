// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100136_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewEventHeaderWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEventHeader()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EventHeaderReader struct {
	r io.Reader
	p *vm.Program
}

func NewEventHeaderReader(r io.Reader) (*EventHeaderReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEventHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EventHeaderReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r EventHeaderReader) Read() (EventHeader, error) {
	t := NewEventHeader()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
