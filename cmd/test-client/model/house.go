// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100112_1.avsc
 */
package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type House struct {
	Address *UnionNullString `json:"Address"`

	Buildingtype Buildingtype `json:"Buildingtype"`

	Color *UnionNullString `json:"Color"`
}

const HouseAvroCRC64Fingerprint = "\x83\xe3\x8cz\xcc\xe8\xb8\xc7"

func NewHouse() House {
	r := House{}
	r.Address = nil
	r.Color = nil
	return r
}

func DeserializeHouse(r io.Reader) (House, error) {
	t := NewHouse()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeHouseFromSchema(r io.Reader, schema string) (House, error) {
	t := NewHouse()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeHouse(r House, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Address, w)
	if err != nil {
		return err
	}
	err = writeBuildingtype(r.Buildingtype, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Color, w)
	if err != nil {
		return err
	}
	return err
}

func (r House) Serialize(w io.Writer) error {
	return writeHouse(r, w)
}

func (r House) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"name\":\"Buildingtype\",\"type\":{\"name\":\"Buildingtype\",\"symbols\":[\"House\",\"Apartment\",\"Cabin\",\"Other\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Color\",\"type\":[\"null\",\"string\"]}],\"name\":\"dp.demoapp.House\",\"type\":\"record\"}"
}

func (r House) SchemaName() string {
	return "dp.demoapp.House"
}

func (_ House) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ House) SetInt(v int32)       { panic("Unsupported operation") }
func (_ House) SetLong(v int64)      { panic("Unsupported operation") }
func (_ House) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ House) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ House) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ House) SetString(v string)   { panic("Unsupported operation") }
func (_ House) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *House) Get(i int) types.Field {
	switch i {
	case 0:
		r.Address = NewUnionNullString()

		return r.Address
	case 1:
		w := BuildingtypeWrapper{Target: &r.Buildingtype}

		return w

	case 2:
		r.Color = NewUnionNullString()

		return r.Color
	}
	panic("Unknown field index")
}

func (r *House) SetDefault(i int) {
	switch i {
	case 0:
		r.Address = nil
		return
	case 2:
		r.Color = nil
		return
	}
	panic("Unknown field index")
}

func (r *House) NullField(i int) {
	switch i {
	case 0:
		r.Address = nil
		return
	case 2:
		r.Color = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ House) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ House) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ House) HintSize(int)                     { panic("Unsupported operation") }
func (_ House) Finalize()                        {}

func (_ House) AvroCRC64Fingerprint() []byte {
	return []byte(HouseAvroCRC64Fingerprint)
}

func (r House) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Address"], err = json.Marshal(r.Address)
	if err != nil {
		return nil, err
	}
	output["Buildingtype"], err = json.Marshal(r.Buildingtype)
	if err != nil {
		return nil, err
	}
	output["Color"], err = json.Marshal(r.Color)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *House) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Address"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Address); err != nil {
			return err
		}
	} else {
		r.Address = NewUnionNullString()

		r.Address = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Buildingtype"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Buildingtype); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Buildingtype")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Color"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Color); err != nil {
			return err
		}
	} else {
		r.Color = NewUnionNullString()

		r.Color = nil
	}
	return nil
}
