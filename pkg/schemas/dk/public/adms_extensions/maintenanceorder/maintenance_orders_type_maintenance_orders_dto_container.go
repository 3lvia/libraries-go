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

func NewMaintenanceOrdersTypeMaintenanceOrdersDtoWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewMaintenanceOrdersTypeMaintenanceOrdersDto()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type MaintenanceOrdersTypeMaintenanceOrdersDtoReader struct {
	r io.Reader
	p *vm.Program
}

func NewMaintenanceOrdersTypeMaintenanceOrdersDtoReader(r io.Reader) (*MaintenanceOrdersTypeMaintenanceOrdersDtoReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewMaintenanceOrdersTypeMaintenanceOrdersDto()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &MaintenanceOrdersTypeMaintenanceOrdersDtoReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r MaintenanceOrdersTypeMaintenanceOrdersDtoReader) Read() (MaintenanceOrdersTypeMaintenanceOrdersDto, error) {
	t := NewMaintenanceOrdersTypeMaintenanceOrdersDto()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
