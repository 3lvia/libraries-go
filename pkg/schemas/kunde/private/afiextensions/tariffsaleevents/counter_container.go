// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100346_2.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewCounterWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewCounter()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type CounterReader struct {
	r io.Reader
	p *vm.Program
}

func NewCounterReader(r io.Reader) (*CounterReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewCounter()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &CounterReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r CounterReader) Read() (Counter, error) {
	t := NewCounter()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
