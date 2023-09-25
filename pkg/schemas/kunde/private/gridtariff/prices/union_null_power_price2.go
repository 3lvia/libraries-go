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

type UnionNullPowerPrice2TypeEnum int

const (
	UnionNullPowerPrice2TypeEnumPowerPrice2 UnionNullPowerPrice2TypeEnum = 1
)

type UnionNullPowerPrice2 struct {
	Null        *types.NullVal
	PowerPrice2 PowerPrice2
	UnionType   UnionNullPowerPrice2TypeEnum
}

func writeUnionNullPowerPrice2(r *UnionNullPowerPrice2, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPowerPrice2TypeEnumPowerPrice2:
		return writePowerPrice2(r.PowerPrice2, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPowerPrice2")
}

func NewUnionNullPowerPrice2() *UnionNullPowerPrice2 {
	return &UnionNullPowerPrice2{}
}

func (r *UnionNullPowerPrice2) Serialize(w io.Writer) error {
	return writeUnionNullPowerPrice2(r, w)
}

func DeserializeUnionNullPowerPrice2(r io.Reader) (*UnionNullPowerPrice2, error) {
	t := NewUnionNullPowerPrice2()
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

func DeserializeUnionNullPowerPrice2FromSchema(r io.Reader, schema string) (*UnionNullPowerPrice2, error) {
	t := NewUnionNullPowerPrice2()
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

func (r *UnionNullPowerPrice2) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"PowerPriceLevel\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Currency\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LevelInfo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MonetaryUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"name\":\"MonthlyActivePowerExTaxes\",\"type\":\"double\"},{\"default\":null,\"name\":\"MonthlyReactivePowerExTaxes\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"MonthlyUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdDown\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdUp\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValueMax\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueMin\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueUnitOfMeasure\",\"type\":[\"null\",\"string\"]}],\"name\":\"PowerPrice\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"PowerPrice2\",\"type\":\"record\"}]"
}

func (_ *UnionNullPowerPrice2) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPowerPrice2) SetLong(v int64) {

	r.UnionType = (UnionNullPowerPrice2TypeEnum)(v)
}

func (r *UnionNullPowerPrice2) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.PowerPrice2 = NewPowerPrice2()
		return &types.Record{Target: (&r.PowerPrice2)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPowerPrice2) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPowerPrice2) Finalize()                        {}

func (r *UnionNullPowerPrice2) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPowerPrice2TypeEnumPowerPrice2:
		return json.Marshal(map[string]interface{}{"GridTariff.Events.V1.PowerPrice2": r.PowerPrice2})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPowerPrice2")
}

func (r *UnionNullPowerPrice2) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["GridTariff.Events.V1.PowerPrice2"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.PowerPrice2)
	}
	return fmt.Errorf("invalid value for *UnionNullPowerPrice2")
}
