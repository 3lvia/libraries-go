// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100444_4.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewIndustryCodeWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewIndustryCode()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type IndustryCodeReader struct {
	r io.Reader
	p *vm.Program
}

func NewIndustryCodeReader(r io.Reader) (*IndustryCodeReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewIndustryCode()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &IndustryCodeReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r IndustryCodeReader) Read() (IndustryCode, error) {
	t := NewIndustryCode()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
