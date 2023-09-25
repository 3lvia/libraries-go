// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100444_4.avsc
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

type UnionNullBytesTypeEnum int

const (
	UnionNullBytesTypeEnumBytes UnionNullBytesTypeEnum = 1
)

type UnionNullBytes struct {
	Null      *types.NullVal
	Bytes     Bytes
	UnionType UnionNullBytesTypeEnum
}

func writeUnionNullBytes(r *UnionNullBytes, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBytesTypeEnumBytes:
		return vm.WriteBytes(r.Bytes, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBytes")
}

func NewUnionNullBytes() *UnionNullBytes {
	return &UnionNullBytes{}
}

func (r *UnionNullBytes) Serialize(w io.Writer) error {
	return writeUnionNullBytes(r, w)
}

func DeserializeUnionNullBytes(r io.Reader) (*UnionNullBytes, error) {
	t := NewUnionNullBytes()
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

func DeserializeUnionNullBytesFromSchema(r io.Reader, schema string) (*UnionNullBytes, error) {
	t := NewUnionNullBytes()
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

func (r *UnionNullBytes) Schema() string {
	return "[\"null\",{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}]"
}

func (_ *UnionNullBytes) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBytes) SetLong(v int64) {

	r.UnionType = (UnionNullBytesTypeEnum)(v)
}

func (r *UnionNullBytes) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &BytesWrapper{Target: (&r.Bytes)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBytes) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBytes) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBytes) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBytes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBytes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBytes) Finalize()                        {}

func (r *UnionNullBytes) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBytesTypeEnumBytes:
		return json.Marshal(map[string]interface{}{"bytes": r.Bytes})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBytes")
}

func (r *UnionNullBytes) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["bytes"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Bytes)
	}
	return fmt.Errorf("invalid value for *UnionNullBytes")
}
