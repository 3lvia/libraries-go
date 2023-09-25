// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100471_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewIdentifiedObjectWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewIdentifiedObject()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type IdentifiedObjectReader struct {
	r io.Reader
	p *vm.Program
}

func NewIdentifiedObjectReader(r io.Reader) (*IdentifiedObjectReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewIdentifiedObject()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &IdentifiedObjectReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r IdentifiedObjectReader) Read() (IdentifiedObject, error) {
	t := NewIdentifiedObject()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
