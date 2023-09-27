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

type UnionNullCarTypeEnum int

const (
	UnionNullCarTypeEnumCar UnionNullCarTypeEnum = 1
)

type UnionNullCar struct {
	Null      *types.NullVal
	Car       Car
	UnionType UnionNullCarTypeEnum
}

func writeUnionNullCar(r *UnionNullCar, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCarTypeEnumCar:
		return writeCar(r.Car, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCar")
}

func NewUnionNullCar() *UnionNullCar {
	return &UnionNullCar{}
}

func (r *UnionNullCar) Serialize(w io.Writer) error {
	return writeUnionNullCar(r, w)
}

func DeserializeUnionNullCar(r io.Reader) (*UnionNullCar, error) {
	t := NewUnionNullCar()
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

func DeserializeUnionNullCarFromSchema(r io.Reader, schema string) (*UnionNullCar, error) {
	t := NewUnionNullCar()
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

func (r *UnionNullCar) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Color\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Gearbox\",\"type\":[\"null\",{\"name\":\"Gearbox\",\"symbols\":[\"Automatic\",\"Manual\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"model\",\"type\":[\"null\",\"string\"]}],\"name\":\"Car\",\"type\":\"record\"}]"
}

func (_ *UnionNullCar) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCar) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCar) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCar) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCar) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCar) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCar) SetLong(v int64) {

	r.UnionType = (UnionNullCarTypeEnum)(v)
}

func (r *UnionNullCar) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Car = NewCar()
		return &types.Record{Target: (&r.Car)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCar) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullCar) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullCar) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullCar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullCar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullCar) Finalize()                        {}

func (r *UnionNullCar) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCarTypeEnumCar:
		return json.Marshal(map[string]interface{}{"dp.demoapp.Car": r.Car})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCar")
}

func (r *UnionNullCar) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["dp.demoapp.Car"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Car)
	}
	return fmt.Errorf("invalid value for *UnionNullCar")
}
