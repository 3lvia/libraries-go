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

var _ = fmt.Printf

type InvoiceAgreementUpdatedEvent struct {
	ActorId int64 `json:"ActorId"`

	EventId string `json:"EventId"`

	EventType string `json:"EventType"`

	ExternalCustomerGuid *UnionNullString `json:"ExternalCustomerGuid"`

	ExternalCustomerReference *UnionNullString `json:"ExternalCustomerReference"`

	Id string `json:"Id"`

	InvoiceAgreement InvoiceAgreement `json:"InvoiceAgreement"`

	InvoiceAgreementId int64 `json:"InvoiceAgreementId"`

	IsOrganization bool `json:"IsOrganization"`

	UpdatedTime *UnionNullLong `json:"UpdatedTime"`
}

const InvoiceAgreementUpdatedEventAvroCRC64Fingerprint = "\xe0\xac\xc6\x1a\x01济"

func NewInvoiceAgreementUpdatedEvent() InvoiceAgreementUpdatedEvent {
	r := InvoiceAgreementUpdatedEvent{}
	r.ExternalCustomerGuid = nil
	r.ExternalCustomerReference = nil
	r.InvoiceAgreement = NewInvoiceAgreement()

	r.UpdatedTime = nil
	return r
}

func DeserializeInvoiceAgreementUpdatedEvent(r io.Reader) (InvoiceAgreementUpdatedEvent, error) {
	t := NewInvoiceAgreementUpdatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInvoiceAgreementUpdatedEventFromSchema(r io.Reader, schema string) (InvoiceAgreementUpdatedEvent, error) {
	t := NewInvoiceAgreementUpdatedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInvoiceAgreementUpdatedEvent(r InvoiceAgreementUpdatedEvent, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.ActorId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EventId, w)
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
	err = writeInvoiceAgreement(r.InvoiceAgreement, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.InvoiceAgreementId, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.IsOrganization, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.UpdatedTime, w)
	if err != nil {
		return err
	}
	return err
}

func (r InvoiceAgreementUpdatedEvent) Serialize(w io.Writer) error {
	return writeInvoiceAgreementUpdatedEvent(r, w)
}

func (r InvoiceAgreementUpdatedEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"ActorId\",\"type\":\"long\"},{\"name\":\"EventId\",\"type\":\"string\"},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGuid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"InvoiceAgreement\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"AccountDistributionInformation\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AccountName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AccountNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AgreementDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AlternativeInvoicingAddressStatus\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BillingAddress\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Address2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AddressOrigin\",\"type\":[\"null\",{\"name\":\"AddressOriginType\",\"symbols\":[\"ActorAddress\",\"OtherPayer\",\"InvoiceAddress\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"City\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Country\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ZipCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"BillingAddress\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Cid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DaysUntilDueDate\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DueDate\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"DunningGroup\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"FormCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FormType\",\"type\":[\"null\",{\"name\":\"FormType\",\"symbols\":[\"Bankgiro\",\"Avtalegiro\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InterestGraceDays\",\"type\":[\"null\",\"int\"]},{\"name\":\"InterestRate\",\"type\":\"double\"},{\"default\":null,\"name\":\"InvoiceAgreementCategoryId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceIssuerSupplementaryProduct\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"InvoiceIssuerTariff\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceLabel\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceLabelFrom\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceLabelTo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SendingMethod\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",{\"name\":\"InvoiceAgreementType\",\"symbols\":[\"InvoiceAgreement\",\"InvoiceAgreementDeliverypoint\",\"PaymentDocumentAgreement\",\"AdvanceAccount\"],\"type\":\"enum\"}]}],\"name\":\"InvoiceAgreement\",\"type\":\"record\"}},{\"name\":\"InvoiceAgreementId\",\"type\":\"long\"},{\"name\":\"IsOrganization\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"UpdatedTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Afiextensions.Schemas.InvoiceAgreementEvents.V1.InvoiceAgreementUpdatedEvent\",\"type\":\"record\"}"
}

func (r InvoiceAgreementUpdatedEvent) SchemaName() string {
	return "Afiextensions.Schemas.InvoiceAgreementEvents.V1.InvoiceAgreementUpdatedEvent"
}

func (_ InvoiceAgreementUpdatedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InvoiceAgreementUpdatedEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.ActorId}

		return w

	case 1:
		w := types.String{Target: &r.EventId}

		return w

	case 2:
		w := types.String{Target: &r.EventType}

		return w

	case 3:
		r.ExternalCustomerGuid = NewUnionNullString()

		return r.ExternalCustomerGuid
	case 4:
		r.ExternalCustomerReference = NewUnionNullString()

		return r.ExternalCustomerReference
	case 5:
		w := types.String{Target: &r.Id}

		return w

	case 6:
		r.InvoiceAgreement = NewInvoiceAgreement()

		w := types.Record{Target: &r.InvoiceAgreement}

		return w

	case 7:
		w := types.Long{Target: &r.InvoiceAgreementId}

		return w

	case 8:
		w := types.Boolean{Target: &r.IsOrganization}

		return w

	case 9:
		r.UpdatedTime = NewUnionNullLong()

		return r.UpdatedTime
	}
	panic("Unknown field index")
}

func (r *InvoiceAgreementUpdatedEvent) SetDefault(i int) {
	switch i {
	case 3:
		r.ExternalCustomerGuid = nil
		return
	case 4:
		r.ExternalCustomerReference = nil
		return
	case 9:
		r.UpdatedTime = nil
		return
	}
	panic("Unknown field index")
}

func (r *InvoiceAgreementUpdatedEvent) NullField(i int) {
	switch i {
	case 3:
		r.ExternalCustomerGuid = nil
		return
	case 4:
		r.ExternalCustomerReference = nil
		return
	case 9:
		r.UpdatedTime = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ InvoiceAgreementUpdatedEvent) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ InvoiceAgreementUpdatedEvent) AppendArray() types.Field { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) HintSize(int)             { panic("Unsupported operation") }
func (_ InvoiceAgreementUpdatedEvent) Finalize()                {}

func (_ InvoiceAgreementUpdatedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(InvoiceAgreementUpdatedEventAvroCRC64Fingerprint)
}

func (r InvoiceAgreementUpdatedEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ActorId"], err = json.Marshal(r.ActorId)
	if err != nil {
		return nil, err
	}
	output["EventId"], err = json.Marshal(r.EventId)
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
	output["InvoiceAgreement"], err = json.Marshal(r.InvoiceAgreement)
	if err != nil {
		return nil, err
	}
	output["InvoiceAgreementId"], err = json.Marshal(r.InvoiceAgreementId)
	if err != nil {
		return nil, err
	}
	output["IsOrganization"], err = json.Marshal(r.IsOrganization)
	if err != nil {
		return nil, err
	}
	output["UpdatedTime"], err = json.Marshal(r.UpdatedTime)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InvoiceAgreementUpdatedEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		return fmt.Errorf("no value specified for ActorId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EventId")
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
		if v, ok := fields["InvoiceAgreement"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceAgreement); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for InvoiceAgreement")
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
		return fmt.Errorf("no value specified for InvoiceAgreementId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IsOrganization"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsOrganization); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IsOrganization")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UpdatedTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdatedTime); err != nil {
			return err
		}
	} else {
		r.UpdatedTime = NewUnionNullLong()

		r.UpdatedTime = nil
	}
	return nil
}
