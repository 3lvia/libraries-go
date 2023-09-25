// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100152_1.avsc
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

type DummyEventHeader struct {
	EventId *UnionNullString `json:"EventId"`

	EventName *UnionNullString `json:"EventName"`

	Source *UnionNullString `json:"Source"`

	TimeStamp string `json:"TimeStamp"`
}

const DummyEventHeaderAvroCRC64Fingerprint = "\xf6o\xf3\xc1\x125/\xf0"

func NewDummyEventHeader() DummyEventHeader {
	r := DummyEventHeader{}
	r.EventId = nil
	r.EventName = nil
	r.Source = nil
	return r
}

func DeserializeDummyEventHeader(r io.Reader) (DummyEventHeader, error) {
	t := NewDummyEventHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDummyEventHeaderFromSchema(r io.Reader, schema string) (DummyEventHeader, error) {
	t := NewDummyEventHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDummyEventHeader(r DummyEventHeader, w io.Writer) error {
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

func (r DummyEventHeader) Serialize(w io.Writer) error {
	return writeDummyEventHeader(r, w)
}

func (r DummyEventHeader) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"EventId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EventName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Source\",\"type\":[\"null\",\"string\"]},{\"name\":\"TimeStamp\",\"type\":\"string\"}],\"name\":\"Skogshorn.Drift.Common.Database.Entity.DummyEventHeader\",\"type\":\"record\"}"
}

func (r DummyEventHeader) SchemaName() string {
	return "Skogshorn.Drift.Common.Database.Entity.DummyEventHeader"
}

func (_ DummyEventHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DummyEventHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DummyEventHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DummyEventHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DummyEventHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DummyEventHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DummyEventHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ DummyEventHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DummyEventHeader) Get(i int) types.Field {
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

func (r *DummyEventHeader) SetDefault(i int) {
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

func (r *DummyEventHeader) NullField(i int) {
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

func (_ DummyEventHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DummyEventHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DummyEventHeader) HintSize(int)                     { panic("Unsupported operation") }
func (_ DummyEventHeader) Finalize()                        {}

func (_ DummyEventHeader) AvroCRC64Fingerprint() []byte {
	return []byte(DummyEventHeaderAvroCRC64Fingerprint)
}

func (r DummyEventHeader) MarshalJSON() ([]byte, error) {
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

func (r *DummyEventHeader) UnmarshalJSON(data []byte) error {
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
