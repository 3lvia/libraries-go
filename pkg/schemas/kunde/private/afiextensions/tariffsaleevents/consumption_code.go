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

type ConsumptionCode struct {
	Id *UnionNullString `json:"Id"`
}

const ConsumptionCodeAvroCRC64Fingerprint = "\xc7[\xf6\x9b߫\x85s"

func NewConsumptionCode() ConsumptionCode {
	r := ConsumptionCode{}
	r.Id = nil
	return r
}

func DeserializeConsumptionCode(r io.Reader) (ConsumptionCode, error) {
	t := NewConsumptionCode()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeConsumptionCodeFromSchema(r io.Reader, schema string) (ConsumptionCode, error) {
	t := NewConsumptionCode()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeConsumptionCode(r ConsumptionCode, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	return err
}

func (r ConsumptionCode) Serialize(w io.Writer) error {
	return writeConsumptionCode(r, w)
}

func (r ConsumptionCode) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]}],\"name\":\"Afiextensions.Schemas.TariffEvents.V1.ConsumptionCode\",\"type\":\"record\"}"
}

func (r ConsumptionCode) SchemaName() string {
	return "Afiextensions.Schemas.TariffEvents.V1.ConsumptionCode"
}

func (_ ConsumptionCode) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ConsumptionCode) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ConsumptionCode) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ConsumptionCode) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ConsumptionCode) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ConsumptionCode) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ConsumptionCode) SetString(v string)   { panic("Unsupported operation") }
func (_ ConsumptionCode) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ConsumptionCode) Get(i int) types.Field {
	switch i {
	case 0:
		r.Id = NewUnionNullString()

		return r.Id
	}
	panic("Unknown field index")
}

func (r *ConsumptionCode) SetDefault(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	}
	panic("Unknown field index")
}

func (r *ConsumptionCode) NullField(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ConsumptionCode) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ConsumptionCode) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ConsumptionCode) HintSize(int)                     { panic("Unsupported operation") }
func (_ ConsumptionCode) Finalize()                        {}

func (_ ConsumptionCode) AvroCRC64Fingerprint() []byte {
	return []byte(ConsumptionCodeAvroCRC64Fingerprint)
}

func (r ConsumptionCode) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ConsumptionCode) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullString()

		r.Id = nil
	}
	return nil
}
