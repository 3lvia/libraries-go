// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100058_21.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewUserWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewUser()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type UserReader struct {
	r io.Reader
	p *vm.Program
}

func NewUserReader(r io.Reader) (*UserReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUser()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UserReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r UserReader) Read() (User, error) {
	t := NewUser()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
