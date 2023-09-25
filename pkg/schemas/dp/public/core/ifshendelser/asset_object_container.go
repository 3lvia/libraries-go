// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100136_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewAssetObjectWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewAssetObject()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type AssetObjectReader struct {
	r io.Reader
	p *vm.Program
}

func NewAssetObjectReader(r io.Reader) (*AssetObjectReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewAssetObject()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &AssetObjectReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r AssetObjectReader) Read() (AssetObject, error) {
	t := NewAssetObject()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
