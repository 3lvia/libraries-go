// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100147_1.avsc
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

type UnionNullEventContentTypeEnum int

const (
	UnionNullEventContentTypeEnumEventContent UnionNullEventContentTypeEnum = 1
)

type UnionNullEventContent struct {
	Null         *types.NullVal
	EventContent EventContent
	UnionType    UnionNullEventContentTypeEnum
}

func writeUnionNullEventContent(r *UnionNullEventContent, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullEventContentTypeEnumEventContent:
		return writeEventContent(r.EventContent, w)
	}
	return fmt.Errorf("invalid value for *UnionNullEventContent")
}

func NewUnionNullEventContent() *UnionNullEventContent {
	return &UnionNullEventContent{}
}

func (r *UnionNullEventContent) Serialize(w io.Writer) error {
	return writeUnionNullEventContent(r, w)
}

func DeserializeUnionNullEventContent(r io.Reader) (*UnionNullEventContent, error) {
	t := NewUnionNullEventContent()
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

func DeserializeUnionNullEventContentFromSchema(r io.Reader, schema string) (*UnionNullEventContent, error) {
	t := NewUnionNullEventContent()
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

func (r *UnionNullEventContent) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AssetObject\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AssetObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ConnectedTo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"ParentAssetObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"Status\",\"type\":{\"name\":\"Status\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Planned\",\"Created\",\"Activated\",\"Deactivated\",\"Installed\",\"Uninstalled\",\"Deleted\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"TechnicalProperties\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TechnicalPropertyId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"TechnicalProperty\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"name\":\"Type\",\"type\":{\"name\":\"AssetObjectType\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Antenne\",\"Maaler\",\"Fake\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"AssetObject\",\"namespace\":\"Skogshorn.Nett.Schemas.Enitites\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"FunctionalObject\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Grid\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Grid\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"IsFaked\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Location\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Block\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FloorNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Leasehold\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalRegion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Section\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Altitude\",\"type\":[\"null\",\"double\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Location\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"ObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plant\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Plant\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"Status\",\"type\":\"Skogshorn.Nett.Schemas.Primitives.Status\"},{\"name\":\"Type\",\"type\":{\"name\":\"FunctionalObjectType\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Maalepunkt\",\"Nettstasjon\",\"Transformator\",\"Trafokrets\",\"Sikring\",\"Relevern\",\"Samleskinne\",\"Transformatorstasjon\",\"Stikkledning\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"FunctionalObject\",\"namespace\":\"Skogshorn.Nett.Schemas.Enitites\",\"type\":\"record\"}]}],\"name\":\"EventContent\",\"type\":\"record\"}]"
}

func (_ *UnionNullEventContent) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullEventContent) SetLong(v int64) {

	r.UnionType = (UnionNullEventContentTypeEnum)(v)
}

func (r *UnionNullEventContent) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.EventContent = NewEventContent()
		return &types.Record{Target: (&r.EventContent)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullEventContent) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullEventContent) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullEventContent) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullEventContent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullEventContent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullEventContent) Finalize()                        {}

func (r *UnionNullEventContent) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullEventContentTypeEnumEventContent:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Schemas.Events.EventContent": r.EventContent})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullEventContent")
}

func (r *UnionNullEventContent) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Schemas.Events.EventContent"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.EventContent)
	}
	return fmt.Errorf("invalid value for *UnionNullEventContent")
}
