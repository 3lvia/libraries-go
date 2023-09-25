// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100136_1.avsc
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

type Plant struct {
	Created *UnionNullString `json:"Created"`

	FunctionalObjectId *UnionNullString `json:"FunctionalObjectId"`

	FuseSize int32 `json:"FuseSize"`

	NominalVoltage int32 `json:"NominalVoltage"`

	Phases int32 `json:"Phases"`

	PlantId *UnionNullString `json:"PlantId"`

	Updated *UnionNullString `json:"Updated"`
}

const PlantAvroCRC64Fingerprint = ";\x04\x0f\xb1\xfbo0\x92"

func NewPlant() Plant {
	r := Plant{}
	r.Created = nil
	r.FunctionalObjectId = nil
	r.PlantId = nil
	r.Updated = nil
	return r
}

func DeserializePlant(r io.Reader) (Plant, error) {
	t := NewPlant()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePlantFromSchema(r io.Reader, schema string) (Plant, error) {
	t := NewPlant()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePlant(r Plant, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Created, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FunctionalObjectId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.FuseSize, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NominalVoltage, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Phases, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PlantId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Updated, w)
	if err != nil {
		return err
	}
	return err
}

func (r Plant) Serialize(w io.Writer) error {
	return writePlant(r, w)
}

func (r Plant) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Created\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Skogshorn.Nett.Generated.Schemas.Enitites.Plant\",\"type\":\"record\"}"
}

func (r Plant) SchemaName() string {
	return "Skogshorn.Nett.Generated.Schemas.Enitites.Plant"
}

func (_ Plant) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Plant) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Plant) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Plant) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Plant) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Plant) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Plant) SetString(v string)   { panic("Unsupported operation") }
func (_ Plant) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Plant) Get(i int) types.Field {
	switch i {
	case 0:
		r.Created = NewUnionNullString()

		return r.Created
	case 1:
		r.FunctionalObjectId = NewUnionNullString()

		return r.FunctionalObjectId
	case 2:
		w := types.Int{Target: &r.FuseSize}

		return w

	case 3:
		w := types.Int{Target: &r.NominalVoltage}

		return w

	case 4:
		w := types.Int{Target: &r.Phases}

		return w

	case 5:
		r.PlantId = NewUnionNullString()

		return r.PlantId
	case 6:
		r.Updated = NewUnionNullString()

		return r.Updated
	}
	panic("Unknown field index")
}

func (r *Plant) SetDefault(i int) {
	switch i {
	case 0:
		r.Created = nil
		return
	case 1:
		r.FunctionalObjectId = nil
		return
	case 5:
		r.PlantId = nil
		return
	case 6:
		r.Updated = nil
		return
	}
	panic("Unknown field index")
}

func (r *Plant) NullField(i int) {
	switch i {
	case 0:
		r.Created = nil
		return
	case 1:
		r.FunctionalObjectId = nil
		return
	case 5:
		r.PlantId = nil
		return
	case 6:
		r.Updated = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Plant) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Plant) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Plant) HintSize(int)                     { panic("Unsupported operation") }
func (_ Plant) Finalize()                        {}

func (_ Plant) AvroCRC64Fingerprint() []byte {
	return []byte(PlantAvroCRC64Fingerprint)
}

func (r Plant) MarshalJSON() ([]byte, error) {
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
	output["FuseSize"], err = json.Marshal(r.FuseSize)
	if err != nil {
		return nil, err
	}
	output["NominalVoltage"], err = json.Marshal(r.NominalVoltage)
	if err != nil {
		return nil, err
	}
	output["Phases"], err = json.Marshal(r.Phases)
	if err != nil {
		return nil, err
	}
	output["PlantId"], err = json.Marshal(r.PlantId)
	if err != nil {
		return nil, err
	}
	output["Updated"], err = json.Marshal(r.Updated)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Plant) UnmarshalJSON(data []byte) error {
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
		r.Created = NewUnionNullString()

		r.Created = nil
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
		if v, ok := fields["FuseSize"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FuseSize); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FuseSize")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NominalVoltage"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NominalVoltage); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NominalVoltage")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Phases"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Phases); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Phases")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PlantId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PlantId); err != nil {
			return err
		}
	} else {
		r.PlantId = NewUnionNullString()

		r.PlantId = nil
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
