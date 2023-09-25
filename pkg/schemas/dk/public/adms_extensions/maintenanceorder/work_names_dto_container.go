// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100487_5.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewWorkNamesDtoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewWorkNamesDto()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type WorkNamesDtoReader struct {
	r io.Reader
	p *vm.Program
}

func NewWorkNamesDtoReader(r io.Reader) (*WorkNamesDtoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewWorkNamesDto()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &WorkNamesDtoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r WorkNamesDtoReader) Read() (WorkNamesDto, error) {
	t := NewWorkNamesDto()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
