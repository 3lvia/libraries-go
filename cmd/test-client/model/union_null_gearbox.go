// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100112_1.avsc
 */
package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullGearboxTypeEnum int

const (
	UnionNullGearboxTypeEnumGearbox UnionNullGearboxTypeEnum = 1
)

type UnionNullGearbox struct {
	Null      *types.NullVal
	Gearbox   Gearbox
	UnionType UnionNullGearboxTypeEnum
}

func writeUnionNullGearbox(r *UnionNullGearbox, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullGearboxTypeEnumGearbox:
		return writeGearbox(r.Gearbox, w)
	}
	return fmt.Errorf("invalid value for *UnionNullGearbox")
}

func NewUnionNullGearbox() *UnionNullGearbox {
	return &UnionNullGearbox{}
}

func (r *UnionNullGearbox) Serialize(w io.Writer) error {
	return writeUnionNullGearbox(r, w)
}

func DeserializeUnionNullGearbox(r io.Reader) (*UnionNullGearbox, error) {
	t := NewUnionNullGearbox()
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

func DeserializeUnionNullGearboxFromSchema(r io.Reader, schema string) (*UnionNullGearbox, error) {
	t := NewUnionNullGearbox()
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

func (r *UnionNullGearbox) Schema() string {
	return "[\"null\",{\"name\":\"Gearbox\",\"symbols\":[\"Automatic\",\"Manual\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullGearbox) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullGearbox) SetLong(v int64) {

	r.UnionType = (UnionNullGearboxTypeEnum)(v)
}

func (r *UnionNullGearbox) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &GearboxWrapper{Target: (&r.Gearbox)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullGearbox) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullGearbox) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullGearbox) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullGearbox) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullGearbox) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullGearbox) Finalize()                        {}

func (r *UnionNullGearbox) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullGearboxTypeEnumGearbox:
		return json.Marshal(map[string]interface{}{"dp.demoapp.Gearbox": r.Gearbox})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullGearbox")
}

func (r *UnionNullGearbox) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["dp.demoapp.Gearbox"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Gearbox)
	}
	return fmt.Errorf("invalid value for *UnionNullGearbox")
}
