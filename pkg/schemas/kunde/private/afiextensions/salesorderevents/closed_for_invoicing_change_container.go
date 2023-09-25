// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100366_4.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewClosedForInvoicingChangeWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewClosedForInvoicingChange()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ClosedForInvoicingChangeReader struct {
	r io.Reader
	p *vm.Program
}

func NewClosedForInvoicingChangeReader(r io.Reader) (*ClosedForInvoicingChangeReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewClosedForInvoicingChange()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ClosedForInvoicingChangeReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ClosedForInvoicingChangeReader) Read() (ClosedForInvoicingChange, error) {
	t := NewClosedForInvoicingChange()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
