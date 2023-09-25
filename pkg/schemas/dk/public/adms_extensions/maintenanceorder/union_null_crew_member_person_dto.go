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

type UnionNullCrewMemberPersonDtoTypeEnum int

const (
	UnionNullCrewMemberPersonDtoTypeEnumCrewMemberPersonDto UnionNullCrewMemberPersonDtoTypeEnum = 1
)

type UnionNullCrewMemberPersonDto struct {
	Null                *types.NullVal
	CrewMemberPersonDto CrewMemberPersonDto
	UnionType           UnionNullCrewMemberPersonDtoTypeEnum
}

func writeUnionNullCrewMemberPersonDto(r *UnionNullCrewMemberPersonDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCrewMemberPersonDtoTypeEnumCrewMemberPersonDto:
		return writeCrewMemberPersonDto(r.CrewMemberPersonDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCrewMemberPersonDto")
}

func NewUnionNullCrewMemberPersonDto() *UnionNullCrewMemberPersonDto {
	return &UnionNullCrewMemberPersonDto{}
}

func (r *UnionNullCrewMemberPersonDto) Serialize(w io.Writer) error {
	return writeUnionNullCrewMemberPersonDto(r, w)
}

func DeserializeUnionNullCrewMemberPersonDto(r io.Reader) (*UnionNullCrewMemberPersonDto, error) {
	t := NewUnionNullCrewMemberPersonDto()
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

func DeserializeUnionNullCrewMemberPersonDtoFromSchema(r io.Reader, schema string) (*UnionNullCrewMemberPersonDto, error) {
	t := NewUnionNullCrewMemberPersonDto()
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

func (r *UnionNullCrewMemberPersonDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"electronicAddress\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"email1\",\"type\":[\"null\",\"string\"]}],\"name\":\"ElectronicAddressDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"employeeId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"firstName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lastModifiedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lastModifiedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"lastName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"mRID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"mobilePhone\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localNumber\",\"type\":[\"null\",\"string\"]}],\"name\":\"TelephoneNumberDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"prefix\",\"type\":[\"null\",\"string\"]}],\"name\":\"CrewMemberPersonDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullCrewMemberPersonDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCrewMemberPersonDto) SetLong(v int64) {

	r.UnionType = (UnionNullCrewMemberPersonDtoTypeEnum)(v)
}

func (r *UnionNullCrewMemberPersonDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.CrewMemberPersonDto = NewCrewMemberPersonDto()
		return &types.Record{Target: (&r.CrewMemberPersonDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCrewMemberPersonDto) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullCrewMemberPersonDto) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullCrewMemberPersonDto) Finalize()                {}

func (r *UnionNullCrewMemberPersonDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCrewMemberPersonDtoTypeEnumCrewMemberPersonDto:
		return json.Marshal(map[string]interface{}{"SesamResponseServices.SesamDomainObjects.CrewMemberPersonDto": r.CrewMemberPersonDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCrewMemberPersonDto")
}

func (r *UnionNullCrewMemberPersonDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SesamResponseServices.SesamDomainObjects.CrewMemberPersonDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.CrewMemberPersonDto)
	}
	return fmt.Errorf("invalid value for *UnionNullCrewMemberPersonDto")
}
