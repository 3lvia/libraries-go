// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100058_21.avsc
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

type UnionNullMethodTypeEnum int

const (
	UnionNullMethodTypeEnumMethod UnionNullMethodTypeEnum = 1
)

type UnionNullMethod struct {
	Null      *types.NullVal
	Method    Method
	UnionType UnionNullMethodTypeEnum
}

func writeUnionNullMethod(r *UnionNullMethod, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullMethodTypeEnumMethod:
		return writeMethod(r.Method, w)
	}
	return fmt.Errorf("invalid value for *UnionNullMethod")
}

func NewUnionNullMethod() *UnionNullMethod {
	return &UnionNullMethod{}
}

func (r *UnionNullMethod) Serialize(w io.Writer) error {
	return writeUnionNullMethod(r, w)
}

func DeserializeUnionNullMethod(r io.Reader) (*UnionNullMethod, error) {
	t := NewUnionNullMethod()
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

func DeserializeUnionNullMethodFromSchema(r io.Reader, schema string) (*UnionNullMethod, error) {
	t := NewUnionNullMethod()
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

func (r *UnionNullMethod) Schema() string {
	return "[\"null\",{\"name\":\"Method\",\"symbols\":[\"Automatic\",\"Manual\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullMethod) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullMethod) SetLong(v int64) {

	r.UnionType = (UnionNullMethodTypeEnum)(v)
}

func (r *UnionNullMethod) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &MethodWrapper{Target: (&r.Method)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullMethod) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullMethod) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullMethod) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullMethod) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMethod) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullMethod) Finalize()                        {}

func (r *UnionNullMethod) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullMethodTypeEnumMethod:
		return json.Marshal(map[string]interface{}{"Msim.Domain.Model.Primitives.Method": r.Method})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullMethod")
}

func (r *UnionNullMethod) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Msim.Domain.Model.Primitives.Method"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Method)
	}
	return fmt.Errorf("invalid value for *UnionNullMethod")
}
