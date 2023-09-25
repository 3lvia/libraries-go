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

func NewDataWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewData()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DataReader struct {
	r io.Reader
	p *vm.Program
}

func NewDataReader(r io.Reader) (*DataReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewData()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DataReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r DataReader) Read() (Data, error) {
	t := NewData()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
