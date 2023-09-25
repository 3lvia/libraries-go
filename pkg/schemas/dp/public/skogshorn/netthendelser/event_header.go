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

type EventHeader struct {
	EventId *UnionNullString `json:"EventId"`

	EventName *UnionNullString `json:"EventName"`

	Source *UnionNullString `json:"Source"`

	TimeStamp string `json:"TimeStamp"`
}

const EventHeaderAvroCRC64Fingerprint = "\xd4_/|{\x03\xddc"

func NewEventHeader() EventHeader {
	r := EventHeader{}
	r.EventId = nil
	r.EventName = nil
	r.Source = nil
	return r
}

func DeserializeEventHeader(r io.Reader) (EventHeader, error) {
	t := NewEventHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventHeaderFromSchema(r io.Reader, schema string) (EventHeader, error) {
	t := NewEventHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventHeader(r EventHeader, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.EventId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EventName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Source, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TimeStamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventHeader) Serialize(w io.Writer) error {
	return writeEventHeader(r, w)
}

func (r EventHeader) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"EventId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EventName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Source\",\"type\":[\"null\",\"string\"]},{\"name\":\"TimeStamp\",\"type\":\"string\"}],\"name\":\"Skogshorn.Nett.Schemas.Events.EventHeader\",\"type\":\"record\"}"
}

func (r EventHeader) SchemaName() string {
	return "Skogshorn.Nett.Schemas.Events.EventHeader"
}

func (_ EventHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ EventHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventHeader) Get(i int) types.Field {
	switch i {
	case 0:
		r.EventId = NewUnionNullString()

		return r.EventId
	case 1:
		r.EventName = NewUnionNullString()

		return r.EventName
	case 2:
		r.Source = NewUnionNullString()

		return r.Source
	case 3:
		w := types.String{Target: &r.TimeStamp}

		return w

	}
	panic("Unknown field index")
}

func (r *EventHeader) SetDefault(i int) {
	switch i {
	case 0:
		r.EventId = nil
		return
	case 1:
		r.EventName = nil
		return
	case 2:
		r.Source = nil
		return
	}
	panic("Unknown field index")
}

func (r *EventHeader) NullField(i int) {
	switch i {
	case 0:
		r.EventId = nil
		return
	case 1:
		r.EventName = nil
		return
	case 2:
		r.Source = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventHeader) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventHeader) Finalize()                        {}

func (_ EventHeader) AvroCRC64Fingerprint() []byte {
	return []byte(EventHeaderAvroCRC64Fingerprint)
}

func (r EventHeader) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["EventId"], err = json.Marshal(r.EventId)
	if err != nil {
		return nil, err
	}
	output["EventName"], err = json.Marshal(r.EventName)
	if err != nil {
		return nil, err
	}
	output["Source"], err = json.Marshal(r.Source)
	if err != nil {
		return nil, err
	}
	output["TimeStamp"], err = json.Marshal(r.TimeStamp)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventHeader) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["EventId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventId); err != nil {
			return err
		}
	} else {
		r.EventId = NewUnionNullString()

		r.EventId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventName); err != nil {
			return err
		}
	} else {
		r.EventName = NewUnionNullString()

		r.EventName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Source"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Source); err != nil {
			return err
		}
	} else {
		r.Source = NewUnionNullString()

		r.Source = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TimeStamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TimeStamp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TimeStamp")
	}
	return nil
}
