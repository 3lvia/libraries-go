// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100147_24.avsc
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

type FunctionObjectChangedEvent struct {
	Content *UnionNullEventContent `json:"Content"`

	Header *UnionNullEventHeader `json:"Header"`
}

const FunctionObjectChangedEventAvroCRC64Fingerprint = "a\xde\xf8\xb8tYS\xb9"

func NewFunctionObjectChangedEvent() FunctionObjectChangedEvent {
	r := FunctionObjectChangedEvent{}
	r.Content = nil
	r.Header = nil
	return r
}

func DeserializeFunctionObjectChangedEvent(r io.Reader) (FunctionObjectChangedEvent, error) {
	t := NewFunctionObjectChangedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFunctionObjectChangedEventFromSchema(r io.Reader, schema string) (FunctionObjectChangedEvent, error) {
	t := NewFunctionObjectChangedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFunctionObjectChangedEvent(r FunctionObjectChangedEvent, w io.Writer) error {
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

func (r FunctionObjectChangedEvent) Serialize(w io.Writer) error {
	return writeFunctionObjectChangedEvent(r, w)
}

func (r FunctionObjectChangedEvent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Content\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AssetObject\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AssetObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ConnectedTo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"ParentAssetObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"Status\",\"type\":{\"name\":\"Status\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Planned\",\"Created\",\"Activated\",\"Deactivated\",\"Installed\",\"Uninstalled\",\"Deleted\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"TechnicalProperties\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TechnicalPropertyId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"TechnicalProperty\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"name\":\"Type\",\"type\":{\"name\":\"AssetObjectType\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Antenne\",\"Maaler\",\"Fake\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"AssetObject\",\"namespace\":\"Skogshorn.Nett.Schemas.Enitites\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"FunctionalObject\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Grid\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Grid\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"IsFaked\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Location\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Block\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FloorNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Leasehold\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalRegion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Section\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Altitude\",\"type\":[\"null\",\"double\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Location\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"ObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plant\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Plant\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"Status\",\"type\":\"Skogshorn.Nett.Schemas.Primitives.Status\"},{\"name\":\"Type\",\"type\":{\"name\":\"FunctionalObjectType\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Maalepunkt\",\"Nettstasjon\",\"Transformator\",\"Trafokrets\",\"Sikring\",\"Relevern\",\"Samleskinne\",\"Transformatorstasjon\",\"Stikkledning\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"FunctionalObject\",\"namespace\":\"Skogshorn.Nett.Schemas.Enitites\",\"type\":\"record\"}]}],\"name\":\"EventContent\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Header\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EventId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EventName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Source\",\"type\":[\"null\",\"string\"]},{\"name\":\"TimeStamp\",\"type\":\"string\"}],\"name\":\"EventHeader\",\"type\":\"record\"}]}],\"name\":\"Skogshorn.Nett.Schemas.Events.FunctionObjectChangedEvent\",\"type\":\"record\"}"
}

func (r FunctionObjectChangedEvent) SchemaName() string {
	return "Skogshorn.Nett.Schemas.Events.FunctionObjectChangedEvent"
}

func (_ FunctionObjectChangedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FunctionObjectChangedEvent) Get(i int) types.Field {
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

func (r *FunctionObjectChangedEvent) SetDefault(i int) {
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

func (r *FunctionObjectChangedEvent) NullField(i int) {
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

func (_ FunctionObjectChangedEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ FunctionObjectChangedEvent) Finalize()                        {}

func (_ FunctionObjectChangedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(FunctionObjectChangedEventAvroCRC64Fingerprint)
}

func (r FunctionObjectChangedEvent) MarshalJSON() ([]byte, error) {
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

func (r *FunctionObjectChangedEvent) UnmarshalJSON(data []byte) error {
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
