// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100492_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewTelemetryDiscreteScadaDtoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewTelemetryDiscreteScadaDto()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type TelemetryDiscreteScadaDtoReader struct {
	r io.Reader
	p *vm.Program
}

func NewTelemetryDiscreteScadaDtoReader(r io.Reader) (*TelemetryDiscreteScadaDtoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewTelemetryDiscreteScadaDto()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &TelemetryDiscreteScadaDtoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r TelemetryDiscreteScadaDtoReader) Read() (TelemetryDiscreteScadaDto, error) {
	t := NewTelemetryDiscreteScadaDto()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
