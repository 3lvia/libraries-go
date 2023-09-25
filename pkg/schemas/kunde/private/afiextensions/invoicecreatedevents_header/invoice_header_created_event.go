// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100397_2.avsc
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

type InvoiceHeaderCreatedEvent struct {
	AccountCode *UnionNullString `json:"AccountCode"`

	ActorId *UnionNullLong `json:"ActorId"`

	Cid string `json:"Cid"`

	CreatedTime *UnionNullLong `json:"CreatedTime"`

	DueDate *UnionNullLong `json:"DueDate"`

	EventType string `json:"EventType"`

	ExternalCustomerGuid *UnionNullString `json:"ExternalCustomerGuid"`

	ExternalCustomerReference *UnionNullString `json:"ExternalCustomerReference"`

	Id string `json:"Id"`

	Invoice InvoiceHeader `json:"Invoice"`

	InvoiceAgreementId *UnionNullLong `json:"InvoiceAgreementId"`

	InvoiceBalance *UnionNullBytes `json:"InvoiceBalance"`

	InvoiceNumber *UnionNullLong `json:"InvoiceNumber"`
}

const InvoiceHeaderCreatedEventAvroCRC64Fingerprint = "\xd6SچC\x06\x9d\x1a"

func NewInvoiceHeaderCreatedEvent() InvoiceHeaderCreatedEvent {
	r := InvoiceHeaderCreatedEvent{}
	r.AccountCode = nil
	r.ActorId = nil
	r.CreatedTime = nil
	r.DueDate = nil
	r.ExternalCustomerGuid = nil
	r.ExternalCustomerReference = nil
	r.Invoice = NewInvoiceHeader()

	r.InvoiceAgreementId = nil
	r.InvoiceBalance = nil
	r.InvoiceNumber = nil
	return r
}

func DeserializeInvoiceHeaderCreatedEvent(r io.Reader) (InvoiceHeaderCreatedEvent, error) {
	t := NewInvoiceHeaderCreatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInvoiceHeaderCreatedEventFromSchema(r io.Reader, schema string) (InvoiceHeaderCreatedEvent, error) {
	t := NewInvoiceHeaderCreatedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInvoiceHeaderCreatedEvent(r InvoiceHeaderCreatedEvent, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.AccountCode, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.ActorId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.CreatedTime, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.DueDate, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EventType, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalCustomerGuid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalCustomerReference, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeInvoiceHeader(r.Invoice, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.InvoiceAgreementId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBytes(r.InvoiceBalance, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.InvoiceNumber, w)
	if err != nil {
		return err
	}
	return err
}

func (r InvoiceHeaderCreatedEvent) Serialize(w io.Writer) error {
	return writeInvoiceHeaderCreatedEvent(r, w)
}

func (r InvoiceHeaderCreatedEvent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"AccountCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"long\"]},{\"name\":\"Cid\",\"type\":\"string\"},{\"default\":null,\"name\":\"CreatedTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"DueDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGuid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Invoice\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Amount\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Amount\",\"type\":{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}},{\"name\":\"CurrencyCode\",\"type\":\"string\"}],\"name\":\"Money\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Balance\",\"type\":[\"null\",\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.Money\"]},{\"name\":\"Cid\",\"type\":\"string\"},{\"default\":null,\"name\":\"CreditationAllowed\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DateCredited\",\"type\":[\"null\",\"string\"]},{\"name\":\"DueDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"GiroAmount\",\"type\":[\"null\",\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.Money\"]},{\"default\":null,\"name\":\"GrossPresentation\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"InvoiceDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"InvoiceNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MeterIds\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"MeterPointIds\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"PrintedDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"ProductionDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"StatusEnum\",\"type\":[\"null\",{\"name\":\"InvoiceStatusEnum\",\"symbols\":[\"Finished\",\"Unfinished\",\"WrittenOff\",\"NotApplicable\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"SumPayableInterestDue\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Term\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Year\",\"type\":[\"null\",\"string\"]}],\"name\":\"InvoiceHeader\",\"type\":\"record\"}},{\"default\":null,\"name\":\"InvoiceAgreementId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"InvoiceBalance\",\"type\":[\"null\",{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}]},{\"default\":null,\"name\":\"InvoiceNumber\",\"type\":[\"null\",\"long\"]}],\"name\":\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.InvoiceHeaderCreatedEvent\",\"type\":\"record\"}"
}

func (r InvoiceHeaderCreatedEvent) SchemaName() string {
	return "Afiextensions.Schemas.InvoiceCreatedEvents.V1.InvoiceHeaderCreatedEvent"
}

func (_ InvoiceHeaderCreatedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InvoiceHeaderCreatedEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.AccountCode = NewUnionNullString()

		return r.AccountCode
	case 1:
		r.ActorId = NewUnionNullLong()

		return r.ActorId
	case 2:
		w := types.String{Target: &r.Cid}

		return w

	case 3:
		r.CreatedTime = NewUnionNullLong()

		return r.CreatedTime
	case 4:
		r.DueDate = NewUnionNullLong()

		return r.DueDate
	case 5:
		w := types.String{Target: &r.EventType}

		return w

	case 6:
		r.ExternalCustomerGuid = NewUnionNullString()

		return r.ExternalCustomerGuid
	case 7:
		r.ExternalCustomerReference = NewUnionNullString()

		return r.ExternalCustomerReference
	case 8:
		w := types.String{Target: &r.Id}

		return w

	case 9:
		r.Invoice = NewInvoiceHeader()

		w := types.Record{Target: &r.Invoice}

		return w

	case 10:
		r.InvoiceAgreementId = NewUnionNullLong()

		return r.InvoiceAgreementId
	case 11:
		r.InvoiceBalance = NewUnionNullBytes()

		return r.InvoiceBalance
	case 12:
		r.InvoiceNumber = NewUnionNullLong()

		return r.InvoiceNumber
	}
	panic("Unknown field index")
}

func (r *InvoiceHeaderCreatedEvent) SetDefault(i int) {
	switch i {
	case 0:
		r.AccountCode = nil
		return
	case 1:
		r.ActorId = nil
		return
	case 3:
		r.CreatedTime = nil
		return
	case 4:
		r.DueDate = nil
		return
	case 6:
		r.ExternalCustomerGuid = nil
		return
	case 7:
		r.ExternalCustomerReference = nil
		return
	case 10:
		r.InvoiceAgreementId = nil
		return
	case 11:
		r.InvoiceBalance = nil
		return
	case 12:
		r.InvoiceNumber = nil
		return
	}
	panic("Unknown field index")
}

func (r *InvoiceHeaderCreatedEvent) NullField(i int) {
	switch i {
	case 0:
		r.AccountCode = nil
		return
	case 1:
		r.ActorId = nil
		return
	case 3:
		r.CreatedTime = nil
		return
	case 4:
		r.DueDate = nil
		return
	case 6:
		r.ExternalCustomerGuid = nil
		return
	case 7:
		r.ExternalCustomerReference = nil
		return
	case 10:
		r.InvoiceAgreementId = nil
		return
	case 11:
		r.InvoiceBalance = nil
		return
	case 12:
		r.InvoiceNumber = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ InvoiceHeaderCreatedEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ InvoiceHeaderCreatedEvent) Finalize()                        {}

func (_ InvoiceHeaderCreatedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(InvoiceHeaderCreatedEventAvroCRC64Fingerprint)
}

func (r InvoiceHeaderCreatedEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["AccountCode"], err = json.Marshal(r.AccountCode)
	if err != nil {
		return nil, err
	}
	output["ActorId"], err = json.Marshal(r.ActorId)
	if err != nil {
		return nil, err
	}
	output["Cid"], err = json.Marshal(r.Cid)
	if err != nil {
		return nil, err
	}
	output["CreatedTime"], err = json.Marshal(r.CreatedTime)
	if err != nil {
		return nil, err
	}
	output["DueDate"], err = json.Marshal(r.DueDate)
	if err != nil {
		return nil, err
	}
	output["EventType"], err = json.Marshal(r.EventType)
	if err != nil {
		return nil, err
	}
	output["ExternalCustomerGuid"], err = json.Marshal(r.ExternalCustomerGuid)
	if err != nil {
		return nil, err
	}
	output["ExternalCustomerReference"], err = json.Marshal(r.ExternalCustomerReference)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Invoice"], err = json.Marshal(r.Invoice)
	if err != nil {
		return nil, err
	}
	output["InvoiceAgreementId"], err = json.Marshal(r.InvoiceAgreementId)
	if err != nil {
		return nil, err
	}
	output["InvoiceBalance"], err = json.Marshal(r.InvoiceBalance)
	if err != nil {
		return nil, err
	}
	output["InvoiceNumber"], err = json.Marshal(r.InvoiceNumber)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InvoiceHeaderCreatedEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["AccountCode"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AccountCode); err != nil {
			return err
		}
	} else {
		r.AccountCode = NewUnionNullString()

		r.AccountCode = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ActorId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ActorId); err != nil {
			return err
		}
	} else {
		r.ActorId = NewUnionNullLong()

		r.ActorId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CreatedTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreatedTime); err != nil {
			return err
		}
	} else {
		r.CreatedTime = NewUnionNullLong()

		r.CreatedTime = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DueDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DueDate); err != nil {
			return err
		}
	} else {
		r.DueDate = NewUnionNullLong()

		r.DueDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventType"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventType); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EventType")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ExternalCustomerGuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalCustomerGuid); err != nil {
			return err
		}
	} else {
		r.ExternalCustomerGuid = NewUnionNullString()

		r.ExternalCustomerGuid = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ExternalCustomerReference"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalCustomerReference); err != nil {
			return err
		}
	} else {
		r.ExternalCustomerReference = NewUnionNullString()

		r.ExternalCustomerReference = nil
	}
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
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Invoice"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Invoice); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Invoice")
	}
	val = func() json.RawMessage {
		if v, ok := fields["InvoiceAgreementId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceAgreementId); err != nil {
			return err
		}
	} else {
		r.InvoiceAgreementId = NewUnionNullLong()

		r.InvoiceAgreementId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["InvoiceBalance"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceBalance); err != nil {
			return err
		}
	} else {
		r.InvoiceBalance = NewUnionNullBytes()

		r.InvoiceBalance = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["InvoiceNumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceNumber); err != nil {
			return err
		}
	} else {
		r.InvoiceNumber = NewUnionNullLong()

		r.InvoiceNumber = nil
	}
	return nil
}
