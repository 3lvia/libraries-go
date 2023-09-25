// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100242_2.avsc
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

type UnionNullTaxConfigPowerTypeEnum int

const (
	UnionNullTaxConfigPowerTypeEnumTaxConfigPower UnionNullTaxConfigPowerTypeEnum = 1
)

type UnionNullTaxConfigPower struct {
	Null           *types.NullVal
	TaxConfigPower TaxConfigPower
	UnionType      UnionNullTaxConfigPowerTypeEnum
}

func writeUnionNullTaxConfigPower(r *UnionNullTaxConfigPower, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullTaxConfigPowerTypeEnumTaxConfigPower:
		return writeTaxConfigPower(r.TaxConfigPower, w)
	}
	return fmt.Errorf("invalid value for *UnionNullTaxConfigPower")
}

func NewUnionNullTaxConfigPower() *UnionNullTaxConfigPower {
	return &UnionNullTaxConfigPower{}
}

func (r *UnionNullTaxConfigPower) Serialize(w io.Writer) error {
	return writeUnionNullTaxConfigPower(r, w)
}

func DeserializeUnionNullTaxConfigPower(r io.Reader) (*UnionNullTaxConfigPower, error) {
	t := NewUnionNullTaxConfigPower()
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

func DeserializeUnionNullTaxConfigPowerFromSchema(r io.Reader, schema string) (*UnionNullTaxConfigPower, error) {
	t := NewUnionNullTaxConfigPower()
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

func (r *UnionNullTaxConfigPower) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"TaxType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxTypeDescription\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxUom\",\"type\":[\"null\",\"string\"]},{\"name\":\"TaxValue\",\"type\":\"double\"}],\"name\":\"TaxConfigPower\",\"type\":\"record\"}]"
}

func (_ *UnionNullTaxConfigPower) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullTaxConfigPower) SetLong(v int64) {

	r.UnionType = (UnionNullTaxConfigPowerTypeEnum)(v)
}

func (r *UnionNullTaxConfigPower) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.TaxConfigPower = NewTaxConfigPower()
		return &types.Record{Target: (&r.TaxConfigPower)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullTaxConfigPower) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullTaxConfigPower) Finalize()                        {}

func (r *UnionNullTaxConfigPower) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullTaxConfigPowerTypeEnumTaxConfigPower:
		return json.Marshal(map[string]interface{}{"GridTariff.Events.V1.TaxConfigPower": r.TaxConfigPower})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullTaxConfigPower")
}

func (r *UnionNullTaxConfigPower) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["GridTariff.Events.V1.TaxConfigPower"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.TaxConfigPower)
	}
	return fmt.Errorf("invalid value for *UnionNullTaxConfigPower")
}
