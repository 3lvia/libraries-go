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

type UnionNullWorkNamesDtoTypeEnum int

const (
	UnionNullWorkNamesDtoTypeEnumWorkNamesDto UnionNullWorkNamesDtoTypeEnum = 1
)

type UnionNullWorkNamesDto struct {
	Null         *types.NullVal
	WorkNamesDto WorkNamesDto
	UnionType    UnionNullWorkNamesDtoTypeEnum
}

func writeUnionNullWorkNamesDto(r *UnionNullWorkNamesDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullWorkNamesDtoTypeEnumWorkNamesDto:
		return writeWorkNamesDto(r.WorkNamesDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullWorkNamesDto")
}

func NewUnionNullWorkNamesDto() *UnionNullWorkNamesDto {
	return &UnionNullWorkNamesDto{}
}

func (r *UnionNullWorkNamesDto) Serialize(w io.Writer) error {
	return writeUnionNullWorkNamesDto(r, w)
}

func DeserializeUnionNullWorkNamesDto(r io.Reader) (*UnionNullWorkNamesDto, error) {
	t := NewUnionNullWorkNamesDto()
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

func DeserializeUnionNullWorkNamesDtoFromSchema(r io.Reader, schema string) (*UnionNullWorkNamesDto, error) {
	t := NewUnionNullWorkNamesDto()
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

func (r *UnionNullWorkNamesDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"WorkNamesDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullWorkNamesDto) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullWorkNamesDto) SetLong(v int64) {

	r.UnionType = (UnionNullWorkNamesDtoTypeEnum)(v)
}

func (r *UnionNullWorkNamesDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.WorkNamesDto = NewWorkNamesDto()
		return &types.Record{Target: (&r.WorkNamesDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullWorkNamesDto) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullWorkNamesDto) Finalize()                        {}

func (r *UnionNullWorkNamesDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullWorkNamesDtoTypeEnumWorkNamesDto:
		return json.Marshal(map[string]interface{}{"SesamResponseServices.SesamDomainObjects.WorkNamesDto": r.WorkNamesDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullWorkNamesDto")
}

func (r *UnionNullWorkNamesDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SesamResponseServices.SesamDomainObjects.WorkNamesDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.WorkNamesDto)
	}
	return fmt.Errorf("invalid value for *UnionNullWorkNamesDto")
}
