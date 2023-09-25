// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100479_10.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewLeverandoerWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewLeverandoer()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type LeverandoerReader struct {
	r io.Reader
	p *vm.Program
}

func NewLeverandoerReader(r io.Reader) (*LeverandoerReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewLeverandoer()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &LeverandoerReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r LeverandoerReader) Read() (Leverandoer, error) {
	t := NewLeverandoer()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
