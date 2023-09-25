// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100398_5.avsc
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

type UnionNullBillingAddressTypeEnum int

const (
	UnionNullBillingAddressTypeEnumBillingAddress UnionNullBillingAddressTypeEnum = 1
)

type UnionNullBillingAddress struct {
	Null           *types.NullVal
	BillingAddress BillingAddress
	UnionType      UnionNullBillingAddressTypeEnum
}

func writeUnionNullBillingAddress(r *UnionNullBillingAddress, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBillingAddressTypeEnumBillingAddress:
		return writeBillingAddress(r.BillingAddress, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBillingAddress")
}

func NewUnionNullBillingAddress() *UnionNullBillingAddress {
	return &UnionNullBillingAddress{}
}

func (r *UnionNullBillingAddress) Serialize(w io.Writer) error {
	return writeUnionNullBillingAddress(r, w)
}

func DeserializeUnionNullBillingAddress(r io.Reader) (*UnionNullBillingAddress, error) {
	t := NewUnionNullBillingAddress()
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

func DeserializeUnionNullBillingAddressFromSchema(r io.Reader, schema string) (*UnionNullBillingAddress, error) {
	t := NewUnionNullBillingAddress()
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

func (r *UnionNullBillingAddress) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Address2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AddressOrigin\",\"type\":[\"null\",{\"name\":\"AddressOriginType\",\"symbols\":[\"ActorAddress\",\"OtherPayer\",\"InvoiceAddress\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"City\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Country\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ZipCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"BillingAddress\",\"type\":\"record\"}]"
}

func (_ *UnionNullBillingAddress) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBillingAddress) SetLong(v int64) {

	r.UnionType = (UnionNullBillingAddressTypeEnum)(v)
}

func (r *UnionNullBillingAddress) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.BillingAddress = NewBillingAddress()
		return &types.Record{Target: (&r.BillingAddress)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBillingAddress) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBillingAddress) Finalize()                        {}

func (r *UnionNullBillingAddress) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBillingAddressTypeEnumBillingAddress:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.InvoiceAgreementEvents.V1.BillingAddress": r.BillingAddress})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBillingAddress")
}

func (r *UnionNullBillingAddress) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.InvoiceAgreementEvents.V1.BillingAddress"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BillingAddress)
	}
	return fmt.Errorf("invalid value for *UnionNullBillingAddress")
}
