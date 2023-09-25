// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100044_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewCarWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewCar()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type CarReader struct {
	r io.Reader
	p *vm.Program
}

func NewCarReader(r io.Reader) (*CarReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewCar()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &CarReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r CarReader) Read() (Car, error) {
	t := NewCar()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
