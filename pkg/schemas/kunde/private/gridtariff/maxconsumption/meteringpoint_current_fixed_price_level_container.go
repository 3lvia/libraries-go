// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100244_2.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewMeteringpointCurrentFixedPriceLevelWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewMeteringpointCurrentFixedPriceLevel()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type MeteringpointCurrentFixedPriceLevelReader struct {
	r io.Reader
	p *vm.Program
}

func NewMeteringpointCurrentFixedPriceLevelReader(r io.Reader) (*MeteringpointCurrentFixedPriceLevelReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewMeteringpointCurrentFixedPriceLevel()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &MeteringpointCurrentFixedPriceLevelReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r MeteringpointCurrentFixedPriceLevelReader) Read() (MeteringpointCurrentFixedPriceLevel, error) {
	t := NewMeteringpointCurrentFixedPriceLevel()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
