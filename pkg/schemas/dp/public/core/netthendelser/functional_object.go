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

var _ = fmt.Printf

type FunctionalObject struct {
	Created string `json:"Created"`

	FunctionalObjectId *UnionNullString `json:"FunctionalObjectId"`

	Grid *UnionNullGrid `json:"Grid"`

	IsFaked bool `json:"IsFaked"`

	Location *UnionNullLocation `json:"Location"`

	ObjectId *UnionNullString `json:"ObjectId"`

	ParentFunctionalObjectId *UnionNullString `json:"ParentFunctionalObjectId"`

	Plant *UnionNullPlant `json:"Plant"`

	Status Status `json:"Status"`

	Type FunctionalObjectType `json:"Type"`

	Updated *UnionNullString `json:"Updated"`
}

const FunctionalObjectAvroCRC64Fingerprint = "\x83\x81u\xa3\x19U\xf3\x89"

func NewFunctionalObject() FunctionalObject {
	r := FunctionalObject{}
	r.FunctionalObjectId = nil
	r.Grid = nil
	r.Location = nil
	r.ObjectId = nil
	r.ParentFunctionalObjectId = nil
	r.Plant = nil
	r.Updated = nil
	return r
}

func DeserializeFunctionalObject(r io.Reader) (FunctionalObject, error) {
	t := NewFunctionalObject()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFunctionalObjectFromSchema(r io.Reader, schema string) (FunctionalObject, error) {
	t := NewFunctionalObject()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFunctionalObject(r FunctionalObject, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Created, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FunctionalObjectId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullGrid(r.Grid, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.IsFaked, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLocation(r.Location, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ObjectId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ParentFunctionalObjectId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPlant(r.Plant, w)
	if err != nil {
		return err
	}
	err = writeStatus(r.Status, w)
	if err != nil {
		return err
	}
	err = writeFunctionalObjectType(r.Type, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Updated, w)
	if err != nil {
		return err
	}
	return err
}

func (r FunctionalObject) Serialize(w io.Writer) error {
	return writeFunctionalObject(r, w)
}

func (r FunctionalObject) Schema() string {
	return "{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Grid\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Grid\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"IsFaked\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Location\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Block\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FloorNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseLetter\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Leasehold\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Municipality\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostalRegion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Section\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Address\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Altitude\",\"type\":[\"null\",\"double\"]},{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"LocationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitude\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Location\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"ObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ParentFunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Plant\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Plant\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]},{\"name\":\"Status\",\"type\":{\"name\":\"Status\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Planned\",\"Created\",\"Activated\",\"Deactivated\",\"Installed\",\"Uninstalled\",\"Deleted\"],\"type\":\"enum\"}},{\"name\":\"Type\",\"type\":{\"name\":\"FunctionalObjectType\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"symbols\":[\"Maalepunkt\",\"Nettstasjon\",\"Transformator\",\"Trafokrets\",\"Sikring\",\"Relevern\",\"Samleskinne\",\"Transformatorstasjon\",\"Stikkledning\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Skogshorn.Nett.Schemas.Enitites.FunctionalObject\",\"type\":\"record\"}"
}

func (r FunctionalObject) SchemaName() string {
	return "Skogshorn.Nett.Schemas.Enitites.FunctionalObject"
}

func (_ FunctionalObject) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FunctionalObject) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FunctionalObject) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FunctionalObject) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FunctionalObject) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FunctionalObject) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FunctionalObject) SetString(v string)   { panic("Unsupported operation") }
func (_ FunctionalObject) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FunctionalObject) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Created}

		return w

	case 1:
		r.FunctionalObjectId = NewUnionNullString()

		return r.FunctionalObjectId
	case 2:
		r.Grid = NewUnionNullGrid()

		return r.Grid
	case 3:
		w := types.Boolean{Target: &r.IsFaked}

		return w

	case 4:
		r.Location = NewUnionNullLocation()

		return r.Location
	case 5:
		r.ObjectId = NewUnionNullString()

		return r.ObjectId
	case 6:
		r.ParentFunctionalObjectId = NewUnionNullString()

		return r.ParentFunctionalObjectId
	case 7:
		r.Plant = NewUnionNullPlant()

		return r.Plant
	case 8:
		w := StatusWrapper{Target: &r.Status}

		return w

	case 9:
		w := FunctionalObjectTypeWrapper{Target: &r.Type}

		return w

	case 10:
		r.Updated = NewUnionNullString()

		return r.Updated
	}
	panic("Unknown field index")
}

func (r *FunctionalObject) SetDefault(i int) {
	switch i {
	case 1:
		r.FunctionalObjectId = nil
		return
	case 2:
		r.Grid = nil
		return
	case 4:
		r.Location = nil
		return
	case 5:
		r.ObjectId = nil
		return
	case 6:
		r.ParentFunctionalObjectId = nil
		return
	case 7:
		r.Plant = nil
		return
	case 10:
		r.Updated = nil
		return
	}
	panic("Unknown field index")
}

func (r *FunctionalObject) NullField(i int) {
	switch i {
	case 1:
		r.FunctionalObjectId = nil
		return
	case 2:
		r.Grid = nil
		return
	case 4:
		r.Location = nil
		return
	case 5:
		r.ObjectId = nil
		return
	case 6:
		r.ParentFunctionalObjectId = nil
		return
	case 7:
		r.Plant = nil
		return
	case 10:
		r.Updated = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ FunctionalObject) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FunctionalObject) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FunctionalObject) HintSize(int)                     { panic("Unsupported operation") }
func (_ FunctionalObject) Finalize()                        {}

func (_ FunctionalObject) AvroCRC64Fingerprint() []byte {
	return []byte(FunctionalObjectAvroCRC64Fingerprint)
}

func (r FunctionalObject) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Created"], err = json.Marshal(r.Created)
	if err != nil {
		return nil, err
	}
	output["FunctionalObjectId"], err = json.Marshal(r.FunctionalObjectId)
	if err != nil {
		return nil, err
	}
	output["Grid"], err = json.Marshal(r.Grid)
	if err != nil {
		return nil, err
	}
	output["IsFaked"], err = json.Marshal(r.IsFaked)
	if err != nil {
		return nil, err
	}
	output["Location"], err = json.Marshal(r.Location)
	if err != nil {
		return nil, err
	}
	output["ObjectId"], err = json.Marshal(r.ObjectId)
	if err != nil {
		return nil, err
	}
	output["ParentFunctionalObjectId"], err = json.Marshal(r.ParentFunctionalObjectId)
	if err != nil {
		return nil, err
	}
	output["Plant"], err = json.Marshal(r.Plant)
	if err != nil {
		return nil, err
	}
	output["Status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	output["Type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["Updated"], err = json.Marshal(r.Updated)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FunctionalObject) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Created"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Created); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Created")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FunctionalObjectId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FunctionalObjectId); err != nil {
			return err
		}
	} else {
		r.FunctionalObjectId = NewUnionNullString()

		r.FunctionalObjectId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Grid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grid); err != nil {
			return err
		}
	} else {
		r.Grid = NewUnionNullGrid()

		r.Grid = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IsFaked"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsFaked); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IsFaked")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Location"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Location); err != nil {
			return err
		}
	} else {
		r.Location = NewUnionNullLocation()

		r.Location = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ObjectId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ObjectId); err != nil {
			return err
		}
	} else {
		r.ObjectId = NewUnionNullString()

		r.ObjectId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ParentFunctionalObjectId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ParentFunctionalObjectId); err != nil {
			return err
		}
	} else {
		r.ParentFunctionalObjectId = NewUnionNullString()

		r.ParentFunctionalObjectId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Plant"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Plant); err != nil {
			return err
		}
	} else {
		r.Plant = NewUnionNullPlant()

		r.Plant = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Status")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Updated"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Updated); err != nil {
			return err
		}
	} else {
		r.Updated = NewUnionNullString()

		r.Updated = nil
	}
	return nil
}
