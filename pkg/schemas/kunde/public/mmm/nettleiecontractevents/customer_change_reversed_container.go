// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100251_9.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewCustomerChangeReversedWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewCustomerChangeReversed()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type CustomerChangeReversedReader struct {
	r io.Reader
	p *vm.Program
}

func NewCustomerChangeReversedReader(r io.Reader) (*CustomerChangeReversedReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewCustomerChangeReversed()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &CustomerChangeReversedReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r CustomerChangeReversedReader) Read() (CustomerChangeReversed, error) {
	t := NewCustomerChangeReversed()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
