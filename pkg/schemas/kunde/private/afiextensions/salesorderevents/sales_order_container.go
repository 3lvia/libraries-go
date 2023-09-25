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

func NewSalesOrderWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewSalesOrder()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type SalesOrderReader struct {
	r io.Reader
	p *vm.Program
}

func NewSalesOrderReader(r io.Reader) (*SalesOrderReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewSalesOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &SalesOrderReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r SalesOrderReader) Read() (SalesOrder, error) {
	t := NewSalesOrder()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
