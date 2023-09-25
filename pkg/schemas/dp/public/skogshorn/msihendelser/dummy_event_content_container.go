// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100124_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewDummyEventContentWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewDummyEventContent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DummyEventContentReader struct {
	r io.Reader
	p *vm.Program
}

func NewDummyEventContentReader(r io.Reader) (*DummyEventContentReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewDummyEventContent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DummyEventContentReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r DummyEventContentReader) Read() (DummyEventContent, error) {
	t := NewDummyEventContent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
