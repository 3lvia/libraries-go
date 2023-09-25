// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100487_5.avsc
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

type UnionNullEquipmentDtoTypeEnum int

const (
	UnionNullEquipmentDtoTypeEnumEquipmentDto UnionNullEquipmentDtoTypeEnum = 1
)

type UnionNullEquipmentDto struct {
	Null         *types.NullVal
	EquipmentDto EquipmentDto
	UnionType    UnionNullEquipmentDtoTypeEnum
}

func writeUnionNullEquipmentDto(r *UnionNullEquipmentDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullEquipmentDtoTypeEnumEquipmentDto:
		return writeEquipmentDto(r.EquipmentDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullEquipmentDto")
}

func NewUnionNullEquipmentDto() *UnionNullEquipmentDto {
	return &UnionNullEquipmentDto{}
}

func (r *UnionNullEquipmentDto) Serialize(w io.Writer) error {
	return writeUnionNullEquipmentDto(r, w)
}

func DeserializeUnionNullEquipmentDto(r io.Reader) (*UnionNullEquipmentDto, error) {
	t := NewUnionNullEquipmentDto()
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

func DeserializeUnionNullEquipmentDtoFromSchema(r io.Reader, schema string) (*UnionNullEquipmentDto, error) {
	t := NewUnionNullEquipmentDto()
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

func (r *UnionNullEquipmentDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"B1TEXT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B2TEXT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B3TEXT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ELTEXT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FeederLV\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Feedername\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Names\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"EquipmentNamesDto\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"VoltageLevelNum\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"mRID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"type\",\"type\":[\"null\",\"string\"]}],\"name\":\"EquipmentDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullEquipmentDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullEquipmentDto) SetLong(v int64) {

	r.UnionType = (UnionNullEquipmentDtoTypeEnum)(v)
}

func (r *UnionNullEquipmentDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.EquipmentDto = NewEquipmentDto()
		return &types.Record{Target: (&r.EquipmentDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullEquipmentDto) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullEquipmentDto) Finalize()                        {}

func (r *UnionNullEquipmentDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullEquipmentDtoTypeEnumEquipmentDto:
		return json.Marshal(map[string]interface{}{"SesamResponseServices.SesamDomainObjects.EquipmentDto": r.EquipmentDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullEquipmentDto")
}

func (r *UnionNullEquipmentDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SesamResponseServices.SesamDomainObjects.EquipmentDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.EquipmentDto)
	}
	return fmt.Errorf("invalid value for *UnionNullEquipmentDto")
}
