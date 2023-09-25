// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100159_3.avsc
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

type CustomerChangedEvent struct {
	Content *UnionNullEventContent `json:"Content"`

	Header *UnionNullEventHeader `json:"Header"`
}

const CustomerChangedEventAvroCRC64Fingerprint = "u\xf1\xbd\xd7\x17\xb0\f\x14"

func NewCustomerChangedEvent() CustomerChangedEvent {
	r := CustomerChangedEvent{}
	r.Content = nil
	r.Header = nil
	return r
}

func DeserializeCustomerChangedEvent(r io.Reader) (CustomerChangedEvent, error) {
	t := NewCustomerChangedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCustomerChangedEventFromSchema(r io.Reader, schema string) (CustomerChangedEvent, error) {
	t := NewCustomerChangedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCustomerChangedEvent(r CustomerChangedEvent, w io.Writer) error {
	var err error
	err = writeUnionNullEventContent(r.Content, w)
	if err != nil {
		return err
	}
	err = writeUnionNullEventHeader(r.Header, w)
	if err != nil {
		return err
	}
	return err
}

func (r CustomerChangedEvent) Serialize(w io.Writer) error {
	return writeCustomerChangedEvent(r, w)
}

func (r CustomerChangedEvent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Content\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Customer\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Communications\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"CommucationType\",\"type\":{\"name\":\"CommunicationType\",\"symbols\":[\"Email\",\"Phone\"],\"type\":\"enum\"}},{\"name\":\"CommunicationId\",\"type\":\"int\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Communication\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Contracts\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"ContractId\",\"type\":\"int\"},{\"default\":null,\"name\":\"EndTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InvoiceAddress\",\"type\":[\"null\",{\"fields\":[{\"name\":\"AddressId\",\"type\":\"int\"},{\"default\":null,\"name\":\"AttentionOf\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CareOf\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"name\":\"FakedAddressId\",\"type\":\"int\"},{\"default\":null,\"name\":\"Floor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"name\":\"HouseNumber\",\"type\":\"int\"},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"OnBehalf\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostBox\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"MeteringPointId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MeterValueProfile\",\"type\":[\"null\",{\"fields\":[{\"name\":\"MeterValueProfileId\",\"type\":\"int\"},{\"default\":null,\"name\":\"Method\",\"type\":[\"null\",{\"name\":\"Method\",\"symbols\":[\"Automatic\",\"Manual\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"Resolution\",\"type\":[\"null\",{\"name\":\"MeterValueResolution\",\"symbols\":[\"Quarterly\",\"Hourly\"],\"type\":\"enum\"}]}],\"name\":\"MeterValueProfile\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"NetProduct\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"name\":\"NetProductId\",\"type\":\"int\"}],\"name\":\"NetProduct\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"PostalAddress\",\"type\":[\"null\",\"Skogshorn.Kunde.Schemas.Entities.Address\"]},{\"name\":\"SettlementMethod\",\"type\":{\"name\":\"SettlementMethod\",\"symbols\":[\"Hourly\",\"Profile\"],\"type\":\"enum\"}},{\"name\":\"StartTime\",\"type\":\"string\"}],\"name\":\"Contract\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"CustomerId\",\"type\":[\"null\",\"string\"]},{\"name\":\"CustomerType\",\"type\":{\"name\":\"CustomerType\",\"symbols\":[\"Person\",\"Company\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"FirstName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdentificationNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"name\":\"Updated\",\"type\":\"string\"}],\"name\":\"Customer\",\"namespace\":\"Skogshorn.Kunde.Schemas.Entities\",\"type\":\"record\"}]}],\"name\":\"EventContent\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Header\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EventId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EventName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Source\",\"type\":[\"null\",\"string\"]},{\"name\":\"TimeStamp\",\"type\":\"string\"}],\"name\":\"EventHeader\",\"type\":\"record\"}]}],\"name\":\"Skogshorn.Kunde.Schemas.Events.CustomerChangedEvent\",\"type\":\"record\"}"
}

func (r CustomerChangedEvent) SchemaName() string {
	return "Skogshorn.Kunde.Schemas.Events.CustomerChangedEvent"
}

func (_ CustomerChangedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ CustomerChangedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CustomerChangedEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Content = NewUnionNullEventContent()

		return r.Content
	case 1:
		r.Header = NewUnionNullEventHeader()

		return r.Header
	}
	panic("Unknown field index")
}

func (r *CustomerChangedEvent) SetDefault(i int) {
	switch i {
	case 0:
		r.Content = nil
		return
	case 1:
		r.Header = nil
		return
	}
	panic("Unknown field index")
}

func (r *CustomerChangedEvent) NullField(i int) {
	switch i {
	case 0:
		r.Content = nil
		return
	case 1:
		r.Header = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CustomerChangedEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CustomerChangedEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CustomerChangedEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ CustomerChangedEvent) Finalize()                        {}

func (_ CustomerChangedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(CustomerChangedEventAvroCRC64Fingerprint)
}

func (r CustomerChangedEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Content"], err = json.Marshal(r.Content)
	if err != nil {
		return nil, err
	}
	output["Header"], err = json.Marshal(r.Header)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CustomerChangedEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Content"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Content); err != nil {
			return err
		}
	} else {
		r.Content = NewUnionNullEventContent()

		r.Content = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Header"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Header); err != nil {
			return err
		}
	} else {
		r.Header = NewUnionNullEventHeader()

		r.Header = nil
	}
	return nil
}
