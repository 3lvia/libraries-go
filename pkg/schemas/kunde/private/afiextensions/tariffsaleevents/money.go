// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100346_2.avsc
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

var _ = fmt.Printf

type Money struct {
	Amount Bytes `json:"Amount"`

	CurrencyCode *UnionNullString `json:"CurrencyCode"`
}

const MoneyAvroCRC64Fingerprint = "[G?\x94ΐ\xcb\xcd"

func NewMoney() Money {
	r := Money{}
	r.CurrencyCode = nil
	return r
}

func DeserializeMoney(r io.Reader) (Money, error) {
	t := NewMoney()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMoneyFromSchema(r io.Reader, schema string) (Money, error) {
	t := NewMoney()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMoney(r Money, w io.Writer) error {
	var err error
	err = vm.WriteBytes(r.Amount, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CurrencyCode, w)
	if err != nil {
		return err
	}
	return err
}

func (r Money) Serialize(w io.Writer) error {
	return writeMoney(r, w)
}

func (r Money) Schema() string {
	return "{\"fields\":[{\"name\":\"Amount\",\"type\":{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}},{\"default\":null,\"name\":\"CurrencyCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"Afiextensions.Schemas.TariffEvents.V1.Money\",\"type\":\"record\"}"
}

func (r Money) SchemaName() string {
	return "Afiextensions.Schemas.TariffEvents.V1.Money"
}

func (_ Money) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Money) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Money) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Money) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Money) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Money) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Money) SetString(v string)   { panic("Unsupported operation") }
func (_ Money) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Money) Get(i int) types.Field {
	switch i {
	case 0:
		w := BytesWrapper{Target: &r.Amount}

		return w

	case 1:
		r.CurrencyCode = NewUnionNullString()

		return r.CurrencyCode
	}
	panic("Unknown field index")
}

func (r *Money) SetDefault(i int) {
	switch i {
	case 1:
		r.CurrencyCode = nil
		return
	}
	panic("Unknown field index")
}

func (r *Money) NullField(i int) {
	switch i {
	case 1:
		r.CurrencyCode = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Money) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Money) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Money) HintSize(int)                     { panic("Unsupported operation") }
func (_ Money) Finalize()                        {}

func (_ Money) AvroCRC64Fingerprint() []byte {
	return []byte(MoneyAvroCRC64Fingerprint)
}

func (r Money) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Amount"], err = json.Marshal(r.Amount)
	if err != nil {
		return nil, err
	}
	output["CurrencyCode"], err = json.Marshal(r.CurrencyCode)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Money) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Amount"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Amount); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Amount")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CurrencyCode"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CurrencyCode); err != nil {
			return err
		}
	} else {
		r.CurrencyCode = NewUnionNullString()

		r.CurrencyCode = nil
	}
	return nil
}
