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

type UnionNullStatusDtoTypeEnum int

const (
	UnionNullStatusDtoTypeEnumStatusDto UnionNullStatusDtoTypeEnum = 1
)

type UnionNullStatusDto struct {
	Null      *types.NullVal
	StatusDto StatusDto
	UnionType UnionNullStatusDtoTypeEnum
}

func writeUnionNullStatusDto(r *UnionNullStatusDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullStatusDtoTypeEnumStatusDto:
		return writeStatusDto(r.StatusDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullStatusDto")
}

func NewUnionNullStatusDto() *UnionNullStatusDto {
	return &UnionNullStatusDto{}
}

func (r *UnionNullStatusDto) Serialize(w io.Writer) error {
	return writeUnionNullStatusDto(r, w)
}

func DeserializeUnionNullStatusDto(r io.Reader) (*UnionNullStatusDto, error) {
	t := NewUnionNullStatusDto()
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

func DeserializeUnionNullStatusDtoFromSchema(r io.Reader, schema string) (*UnionNullStatusDto, error) {
	t := NewUnionNullStatusDto()
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

func (r *UnionNullStatusDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"dateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"value\",\"type\":[\"null\",\"string\"]}],\"name\":\"StatusDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullStatusDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullStatusDto) SetLong(v int64) {

	r.UnionType = (UnionNullStatusDtoTypeEnum)(v)
}

func (r *UnionNullStatusDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.StatusDto = NewStatusDto()
		return &types.Record{Target: (&r.StatusDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullStatusDto) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullStatusDto) Finalize()                        {}

func (r *UnionNullStatusDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullStatusDtoTypeEnumStatusDto:
		return json.Marshal(map[string]interface{}{"SesamResponseServices.SesamDomainObjects.StatusDto": r.StatusDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullStatusDto")
}

func (r *UnionNullStatusDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SesamResponseServices.SesamDomainObjects.StatusDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.StatusDto)
	}
	return fmt.Errorf("invalid value for *UnionNullStatusDto")
}
