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

func NewCrewNamesNameTypeDtoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewCrewNamesNameTypeDto()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type CrewNamesNameTypeDtoReader struct {
	r io.Reader
	p *vm.Program
}

func NewCrewNamesNameTypeDtoReader(r io.Reader) (*CrewNamesNameTypeDtoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewCrewNamesNameTypeDto()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &CrewNamesNameTypeDtoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r CrewNamesNameTypeDtoReader) Read() (CrewNamesNameTypeDto, error) {
	t := NewCrewNamesNameTypeDto()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
