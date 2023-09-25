// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100172_6.avsc
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

type Car struct {
	Color *UnionNullString `json:"Color"`

	Gearbox *UnionNullGearbox `json:"Gearbox"`

	Model *UnionNullString `json:"model"`

	Person *UnionNullPerson `json:"Person"`
}

const CarAvroCRC64Fingerprint = "\x10\x9e\xea\x18\x0f\xecv\x03"

func NewCar() Car {
	r := Car{}
	r.Color = nil
	r.Gearbox = nil
	r.Model = nil
	r.Person = nil
	return r
}

func DeserializeCar(r io.Reader) (Car, error) {
	t := NewCar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCarFromSchema(r io.Reader, schema string) (Car, error) {
	t := NewCar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCar(r Car, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Color, w)
	if err != nil {
		return err
	}
	err = writeUnionNullGearbox(r.Gearbox, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Model, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPerson(r.Person, w)
	if err != nil {
		return err
	}
	return err
}

func (r Car) Serialize(w io.Writer) error {
	return writeCar(r, w)
}

func (r Car) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Color\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Gearbox\",\"type\":[\"null\",{\"name\":\"Gearbox\",\"symbols\":[\"Automatic\",\"Manual\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"model\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Person\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Cars\",\"type\":[\"null\",{\"items\":[\"null\",\"dp.demoapp.Car\"],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Houses\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Address\",\"type\":[\"null\",\"string\"]},{\"name\":\"Buildingtype\",\"type\":{\"name\":\"Buildingtype\",\"symbols\":[\"House\",\"Apartment\",\"Cabin\",\"Other\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"Color\",\"type\":[\"null\",\"string\"]}],\"name\":\"House\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Person\",\"namespace\":\"dp.demoapp\",\"type\":\"record\"}]}],\"name\":\"dp.demoapp.Car\",\"type\":\"record\"}"
}

func (r Car) SchemaName() string {
	return "dp.demoapp.Car"
}

func (_ Car) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Car) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Car) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Car) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Car) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Car) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Car) SetString(v string)   { panic("Unsupported operation") }
func (_ Car) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Car) Get(i int) types.Field {
	switch i {
	case 0:
		r.Color = NewUnionNullString()

		return r.Color
	case 1:
		r.Gearbox = NewUnionNullGearbox()

		return r.Gearbox
	case 2:
		r.Model = NewUnionNullString()

		return r.Model
	case 3:
		r.Person = NewUnionNullPerson()

		return r.Person
	}
	panic("Unknown field index")
}

func (r *Car) SetDefault(i int) {
	switch i {
	case 0:
		r.Color = nil
		return
	case 1:
		r.Gearbox = nil
		return
	case 2:
		r.Model = nil
		return
	case 3:
		r.Person = nil
		return
	}
	panic("Unknown field index")
}

func (r *Car) NullField(i int) {
	switch i {
	case 0:
		r.Color = nil
		return
	case 1:
		r.Gearbox = nil
		return
	case 2:
		r.Model = nil
		return
	case 3:
		r.Person = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Car) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Car) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Car) HintSize(int)                     { panic("Unsupported operation") }
func (_ Car) Finalize()                        {}

func (_ Car) AvroCRC64Fingerprint() []byte {
	return []byte(CarAvroCRC64Fingerprint)
}

func (r Car) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Color"], err = json.Marshal(r.Color)
	if err != nil {
		return nil, err
	}
	output["Gearbox"], err = json.Marshal(r.Gearbox)
	if err != nil {
		return nil, err
	}
	output["model"], err = json.Marshal(r.Model)
	if err != nil {
		return nil, err
	}
	output["Person"], err = json.Marshal(r.Person)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Car) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
	val = func() json.RawMessage {
		if v, ok := fields["Gearbox"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Gearbox); err != nil {
			return err
		}
	} else {
		r.Gearbox = NewUnionNullGearbox()

		r.Gearbox = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["model"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Model); err != nil {
			return err
		}
	} else {
		r.Model = NewUnionNullString()

		r.Model = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Person"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Person); err != nil {
			return err
		}
	} else {
		r.Person = NewUnionNullPerson()

		r.Person = nil
	}
	return nil
}
