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

type UnionNullLocationTypeEnum int

const (
	UnionNullLocationTypeEnumLocation UnionNullLocationTypeEnum = 1
)

type UnionNullLocation struct {
	Null      *types.NullVal
	Location  Location
	UnionType UnionNullLocationTypeEnum
}

func writeUnionNullLocation(r *UnionNullLocation, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullLocationTypeEnumLocation:
		return writeLocation(r.Location, w)
	}
	return fmt.Errorf("invalid value for *UnionNullLocation")
}

func NewUnionNullLocation() *UnionNullLocation {
	return &UnionNullLocation{}
}

func (r *UnionNullLocation) Serialize(w io.Writer) error {
	return writeUnionNullLocation(r, w)
}

func DeserializeUnionNullLocation(r io.Reader) (*UnionNullLocation, error) {
	t := NewUnionNullLocation()
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

func DeserializeUnionNullLocationFromSchema(r io.Reader, schema string) (*UnionNullLocation, error) {
	t := NewUnionNullLocation()
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

func (r *UnionNullLocation) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Block\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FloorNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Leasehold\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalRegion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Section\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Altitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Location\",\"type\":\"record\"}]"
}

func (_ *UnionNullLocation) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullLocation) SetLong(v int64) {

	r.UnionType = (UnionNullLocationTypeEnum)(v)
}

func (r *UnionNullLocation) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Location = NewLocation()
		return &types.Record{Target: (&r.Location)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullLocation) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullLocation) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullLocation) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullLocation) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullLocation) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullLocation) Finalize()                        {}

func (r *UnionNullLocation) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullLocationTypeEnumLocation:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Generated.Schemas.Enitites.Location": r.Location})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullLocation")
}

func (r *UnionNullLocation) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Generated.Schemas.Enitites.Location"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Location)
	}
	return fmt.Errorf("invalid value for *UnionNullLocation")
}
