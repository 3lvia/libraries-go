// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100238_6.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewSingleDialogueSchemaWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewSingleDialogueSchema()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type SingleDialogueSchemaReader struct {
	r io.Reader
	p *vm.Program
}

func NewSingleDialogueSchemaReader(r io.Reader) (*SingleDialogueSchemaReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewSingleDialogueSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &SingleDialogueSchemaReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r SingleDialogueSchemaReader) Read() (SingleDialogueSchema, error) {
	t := NewSingleDialogueSchema()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
