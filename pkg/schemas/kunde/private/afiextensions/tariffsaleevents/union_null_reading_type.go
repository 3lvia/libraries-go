// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100346_2.avsc
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

type UnionNullReadingTypeTypeEnum int

const (
	UnionNullReadingTypeTypeEnumReadingType UnionNullReadingTypeTypeEnum = 1
)

type UnionNullReadingType struct {
	Null        *types.NullVal
	ReadingType ReadingType
	UnionType   UnionNullReadingTypeTypeEnum
}

func writeUnionNullReadingType(r *UnionNullReadingType, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullReadingTypeTypeEnumReadingType:
		return writeReadingType(r.ReadingType, w)
	}
	return fmt.Errorf("invalid value for *UnionNullReadingType")
}

func NewUnionNullReadingType() *UnionNullReadingType {
	return &UnionNullReadingType{}
}

func (r *UnionNullReadingType) Serialize(w io.Writer) error {
	return writeUnionNullReadingType(r, w)
}

func DeserializeUnionNullReadingType(r io.Reader) (*UnionNullReadingType, error) {
	t := NewUnionNullReadingType()
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

func DeserializeUnionNullReadingTypeFromSchema(r io.Reader, schema string) (*UnionNullReadingType, error) {
	t := NewUnionNullReadingType()
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

func (r *UnionNullReadingType) Schema() string {
	return "[\"null\",{\"name\":\"ReadingType\",\"symbols\":[\"Hourly\",\"Manual\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullReadingType) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullReadingType) SetLong(v int64) {

	r.UnionType = (UnionNullReadingTypeTypeEnum)(v)
}

func (r *UnionNullReadingType) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &ReadingTypeWrapper{Target: (&r.ReadingType)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullReadingType) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullReadingType) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullReadingType) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullReadingType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullReadingType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullReadingType) Finalize()                        {}

func (r *UnionNullReadingType) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullReadingTypeTypeEnumReadingType:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.TariffEvents.V1.ReadingType": r.ReadingType})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullReadingType")
}

func (r *UnionNullReadingType) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.TariffEvents.V1.ReadingType"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ReadingType)
	}
	return fmt.Errorf("invalid value for *UnionNullReadingType")
}
