// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100062_2.avsc
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

type UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnum int

const (
	UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnumArrayUnionNullKsqlDataSourceSchema_CARS UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnum = 1
)

type UnionNullArrayUnionNullKsqlDataSourceSchema_CARS struct {
	Null                                    *types.NullVal
	ArrayUnionNullKsqlDataSourceSchema_CARS []*UnionNullKsqlDataSourceSchema_CARS
	UnionType                               UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnum
}

func writeUnionNullArrayUnionNullKsqlDataSourceSchema_CARS(r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnumArrayUnionNullKsqlDataSourceSchema_CARS:
		return writeArrayUnionNullKsqlDataSourceSchema_CARS(r.ArrayUnionNullKsqlDataSourceSchema_CARS, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS")
}

func NewUnionNullArrayUnionNullKsqlDataSourceSchema_CARS() *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS {
	return &UnionNullArrayUnionNullKsqlDataSourceSchema_CARS{}
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullKsqlDataSourceSchema_CARS(r, w)
}

func DeserializeUnionNullArrayUnionNullKsqlDataSourceSchema_CARS(r io.Reader) (*UnionNullArrayUnionNullKsqlDataSourceSchema_CARS, error) {
	t := NewUnionNullArrayUnionNullKsqlDataSourceSchema_CARS()
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

func DeserializeUnionNullArrayUnionNullKsqlDataSourceSchema_CARSFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullKsqlDataSourceSchema_CARS, error) {
	t := NewUnionNullArrayUnionNullKsqlDataSourceSchema_CARS()
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

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MODEL\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GEARBOX\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"COLOR\",\"type\":[\"null\",\"string\"]}],\"name\":\"KsqlDataSourceSchema_CARS\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullKsqlDataSourceSchema_CARS = make([]*UnionNullKsqlDataSourceSchema_CARS, 0)
		return &ArrayUnionNullKsqlDataSourceSchema_CARSWrapper{Target: (&r.ArrayUnionNullKsqlDataSourceSchema_CARS)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) NullField(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) HintSize(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) SetDefault(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) Finalize() {}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullKsqlDataSourceSchema_CARSTypeEnumArrayUnionNullKsqlDataSourceSchema_CARS:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullKsqlDataSourceSchema_CARS})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS")
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullKsqlDataSourceSchema_CARS)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_CARS")
}
