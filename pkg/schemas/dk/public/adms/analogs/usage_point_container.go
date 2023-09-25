// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100493_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewUsagePointWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewUsagePoint()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type UsagePointReader struct {
	r io.Reader
	p *vm.Program
}

func NewUsagePointReader(r io.Reader) (*UsagePointReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUsagePoint()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UsagePointReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r UsagePointReader) Read() (UsagePoint, error) {
	t := NewUsagePoint()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
