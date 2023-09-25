// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100486_7.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewSwitchDtoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewSwitchDto()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type SwitchDtoReader struct {
	r io.Reader
	p *vm.Program
}

func NewSwitchDtoReader(r io.Reader) (*SwitchDtoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewSwitchDto()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &SwitchDtoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r SwitchDtoReader) Read() (SwitchDto, error) {
	t := NewSwitchDto()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
