// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100412_1.avsc
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

type UnionNullCustomerTypeEnum int

const (
	UnionNullCustomerTypeEnumCustomer UnionNullCustomerTypeEnum = 1
)

type UnionNullCustomer struct {
	Null      *types.NullVal
	Customer  Customer
	UnionType UnionNullCustomerTypeEnum
}

func writeUnionNullCustomer(r *UnionNullCustomer, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCustomerTypeEnumCustomer:
		return writeCustomer(r.Customer, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCustomer")
}

func NewUnionNullCustomer() *UnionNullCustomer {
	return &UnionNullCustomer{}
}

func (r *UnionNullCustomer) Serialize(w io.Writer) error {
	return writeUnionNullCustomer(r, w)
}

func DeserializeUnionNullCustomer(r io.Reader) (*UnionNullCustomer, error) {
	t := NewUnionNullCustomer()
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

func DeserializeUnionNullCustomerFromSchema(r io.Reader, schema string) (*UnionNullCustomer, error) {
	t := NewUnionNullCustomer()
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

func (r *UnionNullCustomer) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Id\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"default\":null,\"name\":\"IdentificationNumber\",\"type\":[\"null\",\"string\"]}],\"name\":\"Customer\",\"type\":\"record\"}]"
}

func (_ *UnionNullCustomer) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCustomer) SetLong(v int64) {

	r.UnionType = (UnionNullCustomerTypeEnum)(v)
}

func (r *UnionNullCustomer) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Customer = NewCustomer()
		return &types.Record{Target: (&r.Customer)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCustomer) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullCustomer) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullCustomer) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullCustomer) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullCustomer) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullCustomer) Finalize()                        {}

func (r *UnionNullCustomer) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCustomerTypeEnumCustomer:
		return json.Marshal(map[string]interface{}{"Kunde.Schema.Messages.Customer": r.Customer})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCustomer")
}

func (r *UnionNullCustomer) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Kunde.Schema.Messages.Customer"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Customer)
	}
	return fmt.Errorf("invalid value for *UnionNullCustomer")
}
