// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100443_1.avsc
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

type UnionNullArrayUnionNullInvoiceAccountTypeEnum int

const (
	UnionNullArrayUnionNullInvoiceAccountTypeEnumArrayUnionNullInvoiceAccount UnionNullArrayUnionNullInvoiceAccountTypeEnum = 1
)

type UnionNullArrayUnionNullInvoiceAccount struct {
	Null                         *types.NullVal
	ArrayUnionNullInvoiceAccount []*UnionNullInvoiceAccount
	UnionType                    UnionNullArrayUnionNullInvoiceAccountTypeEnum
}

func writeUnionNullArrayUnionNullInvoiceAccount(r *UnionNullArrayUnionNullInvoiceAccount, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullInvoiceAccountTypeEnumArrayUnionNullInvoiceAccount:
		return writeArrayUnionNullInvoiceAccount(r.ArrayUnionNullInvoiceAccount, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullInvoiceAccount")
}

func NewUnionNullArrayUnionNullInvoiceAccount() *UnionNullArrayUnionNullInvoiceAccount {
	return &UnionNullArrayUnionNullInvoiceAccount{}
}

func (r *UnionNullArrayUnionNullInvoiceAccount) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullInvoiceAccount(r, w)
}

func DeserializeUnionNullArrayUnionNullInvoiceAccount(r io.Reader) (*UnionNullArrayUnionNullInvoiceAccount, error) {
	t := NewUnionNullArrayUnionNullInvoiceAccount()
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

func DeserializeUnionNullArrayUnionNullInvoiceAccountFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullInvoiceAccount, error) {
	t := NewUnionNullArrayUnionNullInvoiceAccount()
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

func (r *UnionNullArrayUnionNullInvoiceAccount) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Account\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Amount\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Amount\",\"type\":{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}},{\"default\":null,\"name\":\"CurrencyCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"Money\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"OperatingUnit\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostingCode\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"PriceCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ProjectNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Responsibility\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"VatCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"VatRate\",\"type\":[\"null\",\"double\"]}],\"name\":\"InvoiceAccount\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullInvoiceAccount) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayUnionNullInvoiceAccount) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullInvoiceAccountTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullInvoiceAccount) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullInvoiceAccount = make([]*UnionNullInvoiceAccount, 0)
		return &ArrayUnionNullInvoiceAccountWrapper{Target: (&r.ArrayUnionNullInvoiceAccount)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullInvoiceAccount) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullInvoiceAccount) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullInvoiceAccount) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullInvoiceAccount) Finalize() {}

func (r *UnionNullArrayUnionNullInvoiceAccount) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullInvoiceAccountTypeEnumArrayUnionNullInvoiceAccount:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullInvoiceAccount})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullInvoiceAccount")
}

func (r *UnionNullArrayUnionNullInvoiceAccount) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullInvoiceAccount)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullInvoiceAccount")
}
