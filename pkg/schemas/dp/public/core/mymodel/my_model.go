// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100154_2.avsc
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

type MyModel struct {
	Code *UnionNullInt `json:"Code"`

	CreatedDate string `json:"CreatedDate"`

	Id *UnionNullString `json:"Id"`

	Name *UnionNullString `json:"Name"`
}

const MyModelAvroCRC64Fingerprint = "R\xb7y\xa8|_\xa9]"

func NewMyModel() MyModel {
	r := MyModel{}
	r.Code = nil
	r.Id = nil
	r.Name = nil
	return r
}

func DeserializeMyModel(r io.Reader) (MyModel, error) {
	t := NewMyModel()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMyModelFromSchema(r io.Reader, schema string) (MyModel, error) {
	t := NewMyModel()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMyModel(r MyModel, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.Code, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CreatedDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r MyModel) Serialize(w io.Writer) error {
	return writeMyModel(r, w)
}

func (r MyModel) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Code\",\"type\":[\"null\",\"int\"]},{\"name\":\"CreatedDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]}],\"name\":\"DemoPublisher.MyModel\",\"type\":\"record\"}"
}

func (r MyModel) SchemaName() string {
	return "DemoPublisher.MyModel"
}

func (_ MyModel) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MyModel) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MyModel) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MyModel) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MyModel) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MyModel) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MyModel) SetString(v string)   { panic("Unsupported operation") }
func (_ MyModel) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MyModel) Get(i int) types.Field {
	switch i {
	case 0:
		r.Code = NewUnionNullInt()

		return r.Code
	case 1:
		w := types.String{Target: &r.CreatedDate}

		return w

	case 2:
		r.Id = NewUnionNullString()

		return r.Id
	case 3:
		r.Name = NewUnionNullString()

		return r.Name
	}
	panic("Unknown field index")
}

func (r *MyModel) SetDefault(i int) {
	switch i {
	case 0:
		r.Code = nil
		return
	case 2:
		r.Id = nil
		return
	case 3:
		r.Name = nil
		return
	}
	panic("Unknown field index")
}

func (r *MyModel) NullField(i int) {
	switch i {
	case 0:
		r.Code = nil
		return
	case 2:
		r.Id = nil
		return
	case 3:
		r.Name = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MyModel) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MyModel) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MyModel) HintSize(int)                     { panic("Unsupported operation") }
func (_ MyModel) Finalize()                        {}

func (_ MyModel) AvroCRC64Fingerprint() []byte {
	return []byte(MyModelAvroCRC64Fingerprint)
}

func (r MyModel) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Code"], err = json.Marshal(r.Code)
	if err != nil {
		return nil, err
	}
	output["CreatedDate"], err = json.Marshal(r.CreatedDate)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MyModel) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Code"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Code); err != nil {
			return err
		}
	} else {
		r.Code = NewUnionNullInt()

		r.Code = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CreatedDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreatedDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CreatedDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullString()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		r.Name = NewUnionNullString()

		r.Name = nil
	}
	return nil
}
