// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100471_3.avsc
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

type EndDeviceEventDetail struct {
	Name *UnionNullString `json:"name"`

	Value *UnionNullString `json:"value"`
}

const EndDeviceEventDetailAvroCRC64Fingerprint = "L\xb6\x93)\xa4G\x99\xda"

func NewEndDeviceEventDetail() EndDeviceEventDetail {
	r := EndDeviceEventDetail{}
	r.Name = nil
	r.Value = nil
	return r
}

func DeserializeEndDeviceEventDetail(r io.Reader) (EndDeviceEventDetail, error) {
	t := NewEndDeviceEventDetail()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEndDeviceEventDetailFromSchema(r io.Reader, schema string) (EndDeviceEventDetail, error) {
	t := NewEndDeviceEventDetail()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEndDeviceEventDetail(r EndDeviceEventDetail, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r EndDeviceEventDetail) Serialize(w io.Writer) error {
	return writeEndDeviceEventDetail(r, w)
}

func (r EndDeviceEventDetail) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"value\",\"type\":[\"null\",\"string\"]}],\"name\":\"MeteringeventsMsiToAdms.Eventhub.Aidon.EndDeviceEventDetail\",\"type\":\"record\"}"
}

func (r EndDeviceEventDetail) SchemaName() string {
	return "MeteringeventsMsiToAdms.Eventhub.Aidon.EndDeviceEventDetail"
}

func (_ EndDeviceEventDetail) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetString(v string)   { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EndDeviceEventDetail) Get(i int) types.Field {
	switch i {
	case 0:
		r.Name = NewUnionNullString()

		return r.Name
	case 1:
		r.Value = NewUnionNullString()

		return r.Value
	}
	panic("Unknown field index")
}

func (r *EndDeviceEventDetail) SetDefault(i int) {
	switch i {
	case 0:
		r.Name = nil
		return
	case 1:
		r.Value = nil
		return
	}
	panic("Unknown field index")
}

func (r *EndDeviceEventDetail) NullField(i int) {
	switch i {
	case 0:
		r.Name = nil
		return
	case 1:
		r.Value = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EndDeviceEventDetail) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) HintSize(int)                     { panic("Unsupported operation") }
func (_ EndDeviceEventDetail) Finalize()                        {}

func (_ EndDeviceEventDetail) AvroCRC64Fingerprint() []byte {
	return []byte(EndDeviceEventDetailAvroCRC64Fingerprint)
}

func (r EndDeviceEventDetail) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EndDeviceEventDetail) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
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
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		r.Value = NewUnionNullString()

		r.Value = nil
	}
	return nil
}
