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

type UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnum int

const (
	UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnumArrayUnionNullWorkTaskNamesDto UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnum = 1
)

type UnionNullArrayUnionNullWorkTaskNamesDto struct {
	Null                           *types.NullVal
	ArrayUnionNullWorkTaskNamesDto []*UnionNullWorkTaskNamesDto
	UnionType                      UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnum
}

func writeUnionNullArrayUnionNullWorkTaskNamesDto(r *UnionNullArrayUnionNullWorkTaskNamesDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnumArrayUnionNullWorkTaskNamesDto:
		return writeArrayUnionNullWorkTaskNamesDto(r.ArrayUnionNullWorkTaskNamesDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullWorkTaskNamesDto")
}

func NewUnionNullArrayUnionNullWorkTaskNamesDto() *UnionNullArrayUnionNullWorkTaskNamesDto {
	return &UnionNullArrayUnionNullWorkTaskNamesDto{}
}

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullWorkTaskNamesDto(r, w)
}

func DeserializeUnionNullArrayUnionNullWorkTaskNamesDto(r io.Reader) (*UnionNullArrayUnionNullWorkTaskNamesDto, error) {
	t := NewUnionNullArrayUnionNullWorkTaskNamesDto()
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

func DeserializeUnionNullArrayUnionNullWorkTaskNamesDtoFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullWorkTaskNamesDto, error) {
	t := NewUnionNullArrayUnionNullWorkTaskNamesDto()
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

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"WorkTaskNamesDto\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetBoolean(v bool)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetInt(v int32)     { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetBytes(v []byte)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullWorkTaskNamesDto = make([]*UnionNullWorkTaskNamesDto, 0)
		return &ArrayUnionNullWorkTaskNamesDtoWrapper{Target: (&r.ArrayUnionNullWorkTaskNamesDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullWorkTaskNamesDto) Finalize() {}

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullWorkTaskNamesDtoTypeEnumArrayUnionNullWorkTaskNamesDto:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullWorkTaskNamesDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullWorkTaskNamesDto")
}

func (r *UnionNullArrayUnionNullWorkTaskNamesDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullWorkTaskNamesDto)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullWorkTaskNamesDto")
}
