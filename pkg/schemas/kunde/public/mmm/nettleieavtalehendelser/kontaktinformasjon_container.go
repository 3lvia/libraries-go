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

func NewKontaktinformasjonWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewKontaktinformasjon()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type KontaktinformasjonReader struct {
	r io.Reader
	p *vm.Program
}

func NewKontaktinformasjonReader(r io.Reader) (*KontaktinformasjonReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewKontaktinformasjon()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &KontaktinformasjonReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r KontaktinformasjonReader) Read() (Kontaktinformasjon, error) {
	t := NewKontaktinformasjon()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
