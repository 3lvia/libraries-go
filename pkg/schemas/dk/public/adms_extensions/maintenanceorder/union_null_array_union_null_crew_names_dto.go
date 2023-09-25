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

type UnionNullArrayUnionNullCrewNamesDtoTypeEnum int

const (
	UnionNullArrayUnionNullCrewNamesDtoTypeEnumArrayUnionNullCrewNamesDto UnionNullArrayUnionNullCrewNamesDtoTypeEnum = 1
)

type UnionNullArrayUnionNullCrewNamesDto struct {
	Null                       *types.NullVal
	ArrayUnionNullCrewNamesDto []*UnionNullCrewNamesDto
	UnionType                  UnionNullArrayUnionNullCrewNamesDtoTypeEnum
}

func writeUnionNullArrayUnionNullCrewNamesDto(r *UnionNullArrayUnionNullCrewNamesDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullCrewNamesDtoTypeEnumArrayUnionNullCrewNamesDto:
		return writeArrayUnionNullCrewNamesDto(r.ArrayUnionNullCrewNamesDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullCrewNamesDto")
}

func NewUnionNullArrayUnionNullCrewNamesDto() *UnionNullArrayUnionNullCrewNamesDto {
	return &UnionNullArrayUnionNullCrewNamesDto{}
}

func (r *UnionNullArrayUnionNullCrewNamesDto) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullCrewNamesDto(r, w)
}

func DeserializeUnionNullArrayUnionNullCrewNamesDto(r io.Reader) (*UnionNullArrayUnionNullCrewNamesDto, error) {
	t := NewUnionNullArrayUnionNullCrewNamesDto()
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

func DeserializeUnionNullArrayUnionNullCrewNamesDtoFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullCrewNamesDto, error) {
	t := NewUnionNullArrayUnionNullCrewNamesDto()
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

func (r *UnionNullArrayUnionNullCrewNamesDto) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"NameType\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"NameTypeAuthority\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"CrewNamesNameTypeNameTypeAuthorityDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"CrewNamesNameTypeDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"CrewNamesDto\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullCrewNamesDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayUnionNullCrewNamesDto) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullCrewNamesDtoTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullCrewNamesDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullCrewNamesDto = make([]*UnionNullCrewNamesDto, 0)
		return &ArrayUnionNullCrewNamesDtoWrapper{Target: (&r.ArrayUnionNullCrewNamesDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullCrewNamesDto) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCrewNamesDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullCrewNamesDto) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullCrewNamesDto) Finalize() {}

func (r *UnionNullArrayUnionNullCrewNamesDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullCrewNamesDtoTypeEnumArrayUnionNullCrewNamesDto:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullCrewNamesDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullCrewNamesDto")
}

func (r *UnionNullArrayUnionNullCrewNamesDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullCrewNamesDto)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullCrewNamesDto")
}
