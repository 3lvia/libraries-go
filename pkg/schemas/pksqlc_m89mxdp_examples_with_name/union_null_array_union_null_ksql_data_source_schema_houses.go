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

type UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnum int

const (
	UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnumArrayUnionNullKsqlDataSourceSchema_HOUSES UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnum = 1
)

type UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES struct {
	Null                                      *types.NullVal
	ArrayUnionNullKsqlDataSourceSchema_HOUSES []*UnionNullKsqlDataSourceSchema_HOUSES
	UnionType                                 UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnum
}

func writeUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES(r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnumArrayUnionNullKsqlDataSourceSchema_HOUSES:
		return writeArrayUnionNullKsqlDataSourceSchema_HOUSES(r.ArrayUnionNullKsqlDataSourceSchema_HOUSES, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES")
}

func NewUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES() *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES {
	return &UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES{}
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES(r, w)
}

func DeserializeUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES(r io.Reader) (*UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES, error) {
	t := NewUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES()
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

func DeserializeUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES, error) {
	t := NewUnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES()
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

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ADDRESS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BUILDINGTYPE\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"COLOR\",\"type\":[\"null\",\"string\"]}],\"name\":\"KsqlDataSourceSchema_HOUSES\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullKsqlDataSourceSchema_HOUSES = make([]*UnionNullKsqlDataSourceSchema_HOUSES, 0)
		return &ArrayUnionNullKsqlDataSourceSchema_HOUSESWrapper{Target: (&r.ArrayUnionNullKsqlDataSourceSchema_HOUSES)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) NullField(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) HintSize(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) SetDefault(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) Finalize() {}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSESTypeEnumArrayUnionNullKsqlDataSourceSchema_HOUSES:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullKsqlDataSourceSchema_HOUSES})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES")
}

func (r *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullKsqlDataSourceSchema_HOUSES)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullKsqlDataSourceSchema_HOUSES")
}
