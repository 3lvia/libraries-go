// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100480_13.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewLeverandoerbytteWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewLeverandoerbytte()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type LeverandoerbytteReader struct {
	r io.Reader
	p *vm.Program
}

func NewLeverandoerbytteReader(r io.Reader) (*LeverandoerbytteReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewLeverandoerbytte()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &LeverandoerbytteReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r LeverandoerbytteReader) Read() (Leverandoerbytte, error) {
	t := NewLeverandoerbytte()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
