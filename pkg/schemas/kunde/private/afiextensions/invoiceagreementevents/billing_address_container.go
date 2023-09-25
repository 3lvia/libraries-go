// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100398_5.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewBillingAddressWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewBillingAddress()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type BillingAddressReader struct {
	r io.Reader
	p *vm.Program
}

func NewBillingAddressReader(r io.Reader) (*BillingAddressReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewBillingAddress()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &BillingAddressReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r BillingAddressReader) Read() (BillingAddress, error) {
	t := NewBillingAddress()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
