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

func NewEventContentWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEventContent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EventContentReader struct {
	r io.Reader
	p *vm.Program
}

func NewEventContentReader(r io.Reader) (*EventContentReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEventContent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EventContentReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r EventContentReader) Read() (EventContent, error) {
	t := NewEventContent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
