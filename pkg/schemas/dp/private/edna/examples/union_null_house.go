// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100172_6.avsc
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

type UnionNullHouseTypeEnum int

const (
	UnionNullHouseTypeEnumHouse UnionNullHouseTypeEnum = 1
)

type UnionNullHouse struct {
	Null      *types.NullVal
	House     House
	UnionType UnionNullHouseTypeEnum
}

func writeUnionNullHouse(r *UnionNullHouse, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullHouseTypeEnumHouse:
		return writeHouse(r.House, w)
	}
	return fmt.Errorf("invalid value for *UnionNullHouse")
}

func NewUnionNullHouse() *UnionNullHouse {
	return &UnionNullHouse{}
}

func (r *UnionNullHouse) Serialize(w io.Writer) error {
	return writeUnionNullHouse(r, w)
}

func DeserializeUnionNullHouse(r io.Reader) (*UnionNullHouse, error) {
	t := NewUnionNullHouse()
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

func DeserializeUnionNullHouseFromSchema(r io.Reader, schema string) (*UnionNullHouse, error) {
	t := NewUnionNullHouse()
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

func (r *UnionNullHouse) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"name\":\"Buildingtype\",\"type\":{\"name\":\"Buildingtype\",\"symbols\":[\"House\",\"Apartment\",\"Cabin\",\"Other\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Color\",\"type\":[\"null\",\"string\"]}],\"name\":\"House\",\"type\":\"record\"}]"
}

func (_ *UnionNullHouse) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullHouse) SetLong(v int64) {

	r.UnionType = (UnionNullHouseTypeEnum)(v)
}

func (r *UnionNullHouse) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.House = NewHouse()
		return &types.Record{Target: (&r.House)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullHouse) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullHouse) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullHouse) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullHouse) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHouse) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullHouse) Finalize()                        {}

func (r *UnionNullHouse) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullHouseTypeEnumHouse:
		return json.Marshal(map[string]interface{}{"dp.demoapp.House": r.House})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullHouse")
}

func (r *UnionNullHouse) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["dp.demoapp.House"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.House)
	}
	return fmt.Errorf("invalid value for *UnionNullHouse")
}
