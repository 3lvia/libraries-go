// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100136_7.avsc
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

type UnionNullFunctionalObjectTypeEnum int

const (
	UnionNullFunctionalObjectTypeEnumFunctionalObject UnionNullFunctionalObjectTypeEnum = 1
)

type UnionNullFunctionalObject struct {
	Null             *types.NullVal
	FunctionalObject FunctionalObject
	UnionType        UnionNullFunctionalObjectTypeEnum
}

func writeUnionNullFunctionalObject(r *UnionNullFunctionalObject, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullFunctionalObjectTypeEnumFunctionalObject:
		return writeFunctionalObject(r.FunctionalObject, w)
	}
	return fmt.Errorf("invalid value for *UnionNullFunctionalObject")
}

func NewUnionNullFunctionalObject() *UnionNullFunctionalObject {
	return &UnionNullFunctionalObject{}
}

func (r *UnionNullFunctionalObject) Serialize(w io.Writer) error {
	return writeUnionNullFunctionalObject(r, w)
}

func DeserializeUnionNullFunctionalObject(r io.Reader) (*UnionNullFunctionalObject, error) {
	t := NewUnionNullFunctionalObject()
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

func DeserializeUnionNullFunctionalObjectFromSchema(r io.Reader, schema string) (*UnionNullFunctionalObject, error) {
	t := NewUnionNullFunctionalObject()
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

func (r *UnionNullFunctionalObject) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Grid\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Grid\",\"type\":\"record\"}]},{\"name\":\"IsFaked\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Location\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Block\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FloorNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Leasehold\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalRegion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Section\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Altitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Location\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"ObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plant\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Plant\",\"type\":\"record\"}]},{\"name\":\"Status\",\"type\":{\"name\":\"Status\",\"symbols\":[\"Planned\",\"Created\",\"Activated\",\"Deactivated\",\"Installed\",\"Uninstalled\",\"Deleted\"],\"type\":\"enum\"}},{\"name\":\"Type\",\"type\":{\"name\":\"FunctionalObjectType\",\"symbols\":[\"Maalepunkt\",\"Nettstasjon\",\"Transformator\",\"Trafokrets\",\"Sikring\",\"Relevern\",\"Samleskinne\",\"Transformatorstasjon\",\"Stikkledning\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"FunctionalObject\",\"type\":\"record\"}]"
}

func (_ *UnionNullFunctionalObject) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullFunctionalObject) SetLong(v int64) {

	r.UnionType = (UnionNullFunctionalObjectTypeEnum)(v)
}

func (r *UnionNullFunctionalObject) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.FunctionalObject = NewFunctionalObject()
		return &types.Record{Target: (&r.FunctionalObject)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullFunctionalObject) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullFunctionalObject) Finalize()                        {}

func (r *UnionNullFunctionalObject) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullFunctionalObjectTypeEnumFunctionalObject:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Generated.Schemas.Enitites.FunctionalObject": r.FunctionalObject})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullFunctionalObject")
}

func (r *UnionNullFunctionalObject) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Generated.Schemas.Enitites.FunctionalObject"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.FunctionalObject)
	}
	return fmt.Errorf("invalid value for *UnionNullFunctionalObject")
}
