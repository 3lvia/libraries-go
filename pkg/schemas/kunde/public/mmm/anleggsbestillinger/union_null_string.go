// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100349_1.avsc
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

type UnionNullStringTypeEnum int

const (
	UnionNullStringTypeEnumString UnionNullStringTypeEnum = 1
)

type UnionNullString struct {
	Null      *types.NullVal
	String    string
	UnionType UnionNullStringTypeEnum
}

func writeUnionNullString(r *UnionNullString, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullStringTypeEnumString:
		return vm.WriteString(r.String, w)
	}
	return fmt.Errorf("invalid value for *UnionNullString")
}

func NewUnionNullString() *UnionNullString {
	return &UnionNullString{}
}

func (r *UnionNullString) Serialize(w io.Writer) error {
	return writeUnionNullString(r, w)
}

func DeserializeUnionNullString(r io.Reader) (*UnionNullString, error) {
	t := NewUnionNullString()
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

func DeserializeUnionNullStringFromSchema(r io.Reader, schema string) (*UnionNullString, error) {
	t := NewUnionNullString()
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

func (r *UnionNullString) Schema() string {
	return "[\"null\",\"string\"]"
}

func (_ *UnionNullString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullString) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullString) SetLong(v int64) {

	r.UnionType = (UnionNullStringTypeEnum)(v)
}

func (r *UnionNullString) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &types.String{Target: (&r.String)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullString) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullString) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullString) Finalize()                        {}

func (r *UnionNullString) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullStringTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullString")
}

func (r *UnionNullString) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.String)
	}
	return fmt.Errorf("invalid value for *UnionNullString")
}
