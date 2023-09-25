// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100419_2.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewEndDeviceEventWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewEndDeviceEvent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type EndDeviceEventReader struct {
	r io.Reader
	p *vm.Program
}

func NewEndDeviceEventReader(r io.Reader) (*EndDeviceEventReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewEndDeviceEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &EndDeviceEventReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r EndDeviceEventReader) Read() (EndDeviceEvent, error) {
	t := NewEndDeviceEvent()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
