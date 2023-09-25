// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100375_2.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullTechnicalAssetDtoTypeEnum int

const (
	UnionNullTechnicalAssetDtoTypeEnumTechnicalAssetDto UnionNullTechnicalAssetDtoTypeEnum = 1
)

type UnionNullTechnicalAssetDto struct {
	Null              *types.NullVal
	TechnicalAssetDto TechnicalAssetDto
	UnionType         UnionNullTechnicalAssetDtoTypeEnum
}

func writeUnionNullTechnicalAssetDto(r *UnionNullTechnicalAssetDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullTechnicalAssetDtoTypeEnumTechnicalAssetDto:
		return writeTechnicalAssetDto(r.TechnicalAssetDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullTechnicalAssetDto")
}

func NewUnionNullTechnicalAssetDto() *UnionNullTechnicalAssetDto {
	return &UnionNullTechnicalAssetDto{}
}

func (r *UnionNullTechnicalAssetDto) Serialize(w io.Writer) error {
	return writeUnionNullTechnicalAssetDto(r, w)
}

func DeserializeUnionNullTechnicalAssetDto(r io.Reader) (*UnionNullTechnicalAssetDto, error) {
	t := NewUnionNullTechnicalAssetDto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullTechnicalAssetDtoFromSchema(r io.Reader, schema string) (*UnionNullTechnicalAssetDto, error) {
	t := NewUnionNullTechnicalAssetDto()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullTechnicalAssetDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"B1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B3\",\"type\":[\"null\",\"string\"]}],\"name\":\"TechnicalAssetDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullTechnicalAssetDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullTechnicalAssetDto) SetLong(v int64) {

	r.UnionType = (UnionNullTechnicalAssetDtoTypeEnum)(v)
}

func (r *UnionNullTechnicalAssetDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.TechnicalAssetDto = NewTechnicalAssetDto()
		return &types.Record{Target: (&r.TechnicalAssetDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullTechnicalAssetDto) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullTechnicalAssetDto) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullTechnicalAssetDto) Finalize()                {}

func (r *UnionNullTechnicalAssetDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullTechnicalAssetDtoTypeEnumTechnicalAssetDto:
		return json.Marshal(map[string]interface{}{"CustomerNotificationSafConsumer.TechnicalAssetDto": r.TechnicalAssetDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullTechnicalAssetDto")
}

func (r *UnionNullTechnicalAssetDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["CustomerNotificationSafConsumer.TechnicalAssetDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.TechnicalAssetDto)
	}
	return fmt.Errorf("invalid value for *UnionNullTechnicalAssetDto")
}
