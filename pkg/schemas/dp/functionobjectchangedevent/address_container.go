// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100135_2.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewAddressWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewAddress()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type AddressReader struct {
	r io.Reader
	p *vm.Program
}

func NewAddressReader(r io.Reader) (*AddressReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewAddress()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &AddressReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r AddressReader) Read() (Address, error) {
	t := NewAddress()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
