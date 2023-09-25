// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100480_13.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewInnflyttWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewInnflytt()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type InnflyttReader struct {
	r io.Reader
	p *vm.Program
}

func NewInnflyttReader(r io.Reader) (*InnflyttReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewInnflytt()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &InnflyttReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r InnflyttReader) Read() (Innflytt, error) {
	t := NewInnflytt()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
