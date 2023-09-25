// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100471_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewReadingWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewReading()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ReadingReader struct {
	r io.Reader
	p *vm.Program
}

func NewReadingReader(r io.Reader) (*ReadingReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewReading()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ReadingReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ReadingReader) Read() (Reading, error) {
	t := NewReading()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
