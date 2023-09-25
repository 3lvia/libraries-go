// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100395_1.avsc
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

type UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum int

const (
	UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInitialInvoiceDebtCollectionCreatedEvent UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum = 0

	UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionCreatedEvent UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum = 1

	UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionUpdatedEvent UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum = 2
)

type UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent struct {
	InitialInvoiceDebtCollectionCreatedEvent InitialInvoiceDebtCollectionCreatedEvent
	InvoiceDebtCollectionCreatedEvent        InvoiceDebtCollectionCreatedEvent
	InvoiceDebtCollectionUpdatedEvent        InvoiceDebtCollectionUpdatedEvent
	UnionType                                UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum
}

func writeUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent(r UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInitialInvoiceDebtCollectionCreatedEvent:
		return writeInitialInvoiceDebtCollectionCreatedEvent(r.InitialInvoiceDebtCollectionCreatedEvent, w)
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionCreatedEvent:
		return writeInvoiceDebtCollectionCreatedEvent(r.InvoiceDebtCollectionCreatedEvent, w)
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionUpdatedEvent:
		return writeInvoiceDebtCollectionUpdatedEvent(r.InvoiceDebtCollectionUpdatedEvent, w)
	}
	return fmt.Errorf("invalid value for UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent")
}

func NewUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent() UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent {
	return UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent{}
}

func (r UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) Serialize(w io.Writer) error {
	return writeUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent(r, w)
}

func DeserializeUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent(r io.Reader) (UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent, error) {
	t := NewUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventFromSchema(r io.Reader, schema string) (UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent, error) {
	t := NewUnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) Schema() string {
	return "[{\"fields\":[{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"ActorType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AddressLine\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AddressLine2\",\"type\":[\"null\",\"string\"]},{\"name\":\"Amount\",\"type\":\"double\"},{\"default\":null,\"name\":\"CaseNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"City\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CompanyName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ContactPersonName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Country\",\"type\":[\"null\",\"string\"]},{\"name\":\"DebtCollectionAgency\",\"type\":\"int\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DueDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"DunningFeeAmount\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGUID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FirstName\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"default\":null,\"name\":\"IdentityNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InterestAmount\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"InvoiceAgreementId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"InvoiceDebtCollectionEventType\",\"type\":[\"null\",\"string\"]},{\"name\":\"InvoiceId\",\"type\":\"long\"},{\"default\":null,\"name\":\"LastName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MeterId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MeterPointAddress\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MeterPointId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PhoneCell\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PhonePrivate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PhoneWork\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ProductLineStatus\",\"type\":[\"null\",{\"name\":\"SalesOrderStatus\",\"symbols\":[\"All\",\"Active\",\"Anulled\",\"NeedsPriceChange\",\"HandOverInProgress\",\"AboutToBeClosed\",\"Closed\",\"InCloseProcess\",\"AnullmentInitialised\",\"ActivationInitialised\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"SourceId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"SourceType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TransactionDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ZipCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"InitialInvoiceDebtCollectionCreatedEvent\",\"namespace\":\"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1\",\"type\":\"record\"},{\"fields\":[{\"name\":\"Amount\",\"type\":\"double\"},{\"default\":null,\"name\":\"CaseNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cid\",\"type\":[\"null\",\"string\"]},{\"name\":\"DebtCollectionAgency\",\"type\":\"int\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGUID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FirstDueDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Id\",\"type\":\"string\"},{\"default\":null,\"name\":\"InvoiceDebtCollectionEventType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceDebtCollectionReversedEventType\",\"type\":[\"null\",\"string\"]},{\"name\":\"InvoiceId\",\"type\":\"long\"},{\"default\":null,\"name\":\"NumberOfPayments\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"SourceId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"SourceType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TransactionDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"InvoiceDebtCollectionCreatedEvent\",\"namespace\":\"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1\",\"type\":\"record\"},{\"fields\":[{\"name\":\"Amount\",\"type\":\"double\"},{\"default\":null,\"name\":\"CaseNumber\",\"type\":[\"null\",\"string\"]},{\"name\":\"DebtCollectionAgency\",\"type\":\"int\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGUID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"default\":null,\"name\":\"InvoiceDebtCollectionEventType\",\"type\":[\"null\",\"string\"]},{\"name\":\"InvoiceId\",\"type\":\"long\"},{\"default\":null,\"name\":\"SourceId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"SourceType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TransactionDate\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"InvoiceDebtCollectionUpdatedEvent\",\"namespace\":\"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1\",\"type\":\"record\"}]"
}

func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetLong(v int64) {

	r.UnionType = (UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnum)(v)
}

func (r *UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) Get(i int) types.Field {

	switch i {
	case 0:
		r.InitialInvoiceDebtCollectionCreatedEvent = NewInitialInvoiceDebtCollectionCreatedEvent()
		return &types.Record{Target: (&r.InitialInvoiceDebtCollectionCreatedEvent)}
	case 1:
		r.InvoiceDebtCollectionCreatedEvent = NewInvoiceDebtCollectionCreatedEvent()
		return &types.Record{Target: (&r.InvoiceDebtCollectionCreatedEvent)}
	case 2:
		r.InvoiceDebtCollectionUpdatedEvent = NewInvoiceDebtCollectionUpdatedEvent()
		return &types.Record{Target: (&r.InvoiceDebtCollectionUpdatedEvent)}
	}
	panic("Unknown field index")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) NullField(i int) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) HintSize(i int) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) SetDefault(i int) {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) Finalize() {
}

func (r UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInitialInvoiceDebtCollectionCreatedEvent:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InitialInvoiceDebtCollectionCreatedEvent": r.InitialInvoiceDebtCollectionCreatedEvent})
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionCreatedEvent:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InvoiceDebtCollectionCreatedEvent": r.InvoiceDebtCollectionCreatedEvent})
	case UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEventTypeEnumInvoiceDebtCollectionUpdatedEvent:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InvoiceDebtCollectionUpdatedEvent": r.InvoiceDebtCollectionUpdatedEvent})
	}
	return nil, fmt.Errorf("invalid value for UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent")
}

func (r *UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InitialInvoiceDebtCollectionCreatedEvent"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.InitialInvoiceDebtCollectionCreatedEvent)
	}
	if value, ok := fields["Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InvoiceDebtCollectionCreatedEvent"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.InvoiceDebtCollectionCreatedEvent)
	}
	if value, ok := fields["Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.InvoiceDebtCollectionUpdatedEvent"]; ok {
		r.UnionType = 2
		return json.Unmarshal([]byte(value), &r.InvoiceDebtCollectionUpdatedEvent)
	}
	return fmt.Errorf("invalid value for UnionInitialInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionCreatedEventInvoiceDebtCollectionUpdatedEvent")
}
