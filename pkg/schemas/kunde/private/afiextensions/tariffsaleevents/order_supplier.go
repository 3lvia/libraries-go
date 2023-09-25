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

type OrderSupplier struct {
	AgreementStatus *UnionNullString `json:"AgreementStatus"`

	BalanceResponsibleId *UnionNullInt `json:"BalanceResponsibleId"`

	ContactPerson *UnionNullString `json:"ContactPerson"`

	FromDate *UnionNullString `json:"FromDate"`

	HourlyRead bool `json:"HourlyRead"`

	LastRead *UnionNullString `json:"LastRead"`

	ReadingFrequency *UnionNullString `json:"ReadingFrequency"`

	SettlementMethod *UnionNullString `json:"SettlementMethod"`

	SupplierId *UnionNullInt `json:"SupplierId"`

	ToDate *UnionNullString `json:"ToDate"`
}

const OrderSupplierAvroCRC64Fingerprint = "\x91\xacs\x94\x10N\xfe\x92"

func NewOrderSupplier() OrderSupplier {
	r := OrderSupplier{}
	r.AgreementStatus = nil
	r.BalanceResponsibleId = nil
	r.ContactPerson = nil
	r.FromDate = nil
	r.LastRead = nil
	r.ReadingFrequency = nil
	r.SettlementMethod = nil
	r.SupplierId = nil
	r.ToDate = nil
	return r
}

func DeserializeOrderSupplier(r io.Reader) (OrderSupplier, error) {
	t := NewOrderSupplier()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrderSupplierFromSchema(r io.Reader, schema string) (OrderSupplier, error) {
	t := NewOrderSupplier()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrderSupplier(r OrderSupplier, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.AgreementStatus, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.BalanceResponsibleId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContactPerson, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FromDate, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.HourlyRead, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LastRead, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReadingFrequency, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SettlementMethod, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.SupplierId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ToDate, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrderSupplier) Serialize(w io.Writer) error {
	return writeOrderSupplier(r, w)
}

func (r OrderSupplier) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"AgreementStatus\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BalanceResponsibleId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ContactPerson\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FromDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"HourlyRead\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"LastRead\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ReadingFrequency\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SettlementMethod\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SupplierId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ToDate\",\"type\":[\"null\",\"string\"]}],\"name\":\"Afiextensions.Schemas.TariffEvents.V1.OrderSupplier\",\"type\":\"record\"}"
}

func (r OrderSupplier) SchemaName() string {
	return "Afiextensions.Schemas.TariffEvents.V1.OrderSupplier"
}

func (_ OrderSupplier) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrderSupplier) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrderSupplier) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrderSupplier) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrderSupplier) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrderSupplier) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrderSupplier) SetString(v string)   { panic("Unsupported operation") }
func (_ OrderSupplier) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrderSupplier) Get(i int) types.Field {
	switch i {
	case 0:
		r.AgreementStatus = NewUnionNullString()

		return r.AgreementStatus
	case 1:
		r.BalanceResponsibleId = NewUnionNullInt()

		return r.BalanceResponsibleId
	case 2:
		r.ContactPerson = NewUnionNullString()

		return r.ContactPerson
	case 3:
		r.FromDate = NewUnionNullString()

		return r.FromDate
	case 4:
		w := types.Boolean{Target: &r.HourlyRead}

		return w

	case 5:
		r.LastRead = NewUnionNullString()

		return r.LastRead
	case 6:
		r.ReadingFrequency = NewUnionNullString()

		return r.ReadingFrequency
	case 7:
		r.SettlementMethod = NewUnionNullString()

		return r.SettlementMethod
	case 8:
		r.SupplierId = NewUnionNullInt()

		return r.SupplierId
	case 9:
		r.ToDate = NewUnionNullString()

		return r.ToDate
	}
	panic("Unknown field index")
}

func (r *OrderSupplier) SetDefault(i int) {
	switch i {
	case 0:
		r.AgreementStatus = nil
		return
	case 1:
		r.BalanceResponsibleId = nil
		return
	case 2:
		r.ContactPerson = nil
		return
	case 3:
		r.FromDate = nil
		return
	case 5:
		r.LastRead = nil
		return
	case 6:
		r.ReadingFrequency = nil
		return
	case 7:
		r.SettlementMethod = nil
		return
	case 8:
		r.SupplierId = nil
		return
	case 9:
		r.ToDate = nil
		return
	}
	panic("Unknown field index")
}

func (r *OrderSupplier) NullField(i int) {
	switch i {
	case 0:
		r.AgreementStatus = nil
		return
	case 1:
		r.BalanceResponsibleId = nil
		return
	case 2:
		r.ContactPerson = nil
		return
	case 3:
		r.FromDate = nil
		return
	case 5:
		r.LastRead = nil
		return
	case 6:
		r.ReadingFrequency = nil
		return
	case 7:
		r.SettlementMethod = nil
		return
	case 8:
		r.SupplierId = nil
		return
	case 9:
		r.ToDate = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ OrderSupplier) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrderSupplier) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrderSupplier) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrderSupplier) Finalize()                        {}

func (_ OrderSupplier) AvroCRC64Fingerprint() []byte {
	return []byte(OrderSupplierAvroCRC64Fingerprint)
}

func (r OrderSupplier) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["AgreementStatus"], err = json.Marshal(r.AgreementStatus)
	if err != nil {
		return nil, err
	}
	output["BalanceResponsibleId"], err = json.Marshal(r.BalanceResponsibleId)
	if err != nil {
		return nil, err
	}
	output["ContactPerson"], err = json.Marshal(r.ContactPerson)
	if err != nil {
		return nil, err
	}
	output["FromDate"], err = json.Marshal(r.FromDate)
	if err != nil {
		return nil, err
	}
	output["HourlyRead"], err = json.Marshal(r.HourlyRead)
	if err != nil {
		return nil, err
	}
	output["LastRead"], err = json.Marshal(r.LastRead)
	if err != nil {
		return nil, err
	}
	output["ReadingFrequency"], err = json.Marshal(r.ReadingFrequency)
	if err != nil {
		return nil, err
	}
	output["SettlementMethod"], err = json.Marshal(r.SettlementMethod)
	if err != nil {
		return nil, err
	}
	output["SupplierId"], err = json.Marshal(r.SupplierId)
	if err != nil {
		return nil, err
	}
	output["ToDate"], err = json.Marshal(r.ToDate)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrderSupplier) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["AgreementStatus"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AgreementStatus); err != nil {
			return err
		}
	} else {
		r.AgreementStatus = NewUnionNullString()

		r.AgreementStatus = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["BalanceResponsibleId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BalanceResponsibleId); err != nil {
			return err
		}
	} else {
		r.BalanceResponsibleId = NewUnionNullInt()

		r.BalanceResponsibleId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContactPerson"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContactPerson); err != nil {
			return err
		}
	} else {
		r.ContactPerson = NewUnionNullString()

		r.ContactPerson = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FromDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FromDate); err != nil {
			return err
		}
	} else {
		r.FromDate = NewUnionNullString()

		r.FromDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["HourlyRead"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.HourlyRead); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for HourlyRead")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LastRead"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LastRead); err != nil {
			return err
		}
	} else {
		r.LastRead = NewUnionNullString()

		r.LastRead = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReadingFrequency"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReadingFrequency); err != nil {
			return err
		}
	} else {
		r.ReadingFrequency = NewUnionNullString()

		r.ReadingFrequency = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SettlementMethod"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SettlementMethod); err != nil {
			return err
		}
	} else {
		r.SettlementMethod = NewUnionNullString()

		r.SettlementMethod = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SupplierId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SupplierId); err != nil {
			return err
		}
	} else {
		r.SupplierId = NewUnionNullInt()

		r.SupplierId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ToDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ToDate); err != nil {
			return err
		}
	} else {
		r.ToDate = NewUnionNullString()

		r.ToDate = nil
	}
	return nil
}
