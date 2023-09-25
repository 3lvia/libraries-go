// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100126_1.avsc
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

type DummyEventContent struct {
	ObjectValue *UnionNullString `json:"ObjectValue"`
}

const DummyEventContentAvroCRC64Fingerprint = "\xea;\xce\x12p\x90\xad\x8c"

func NewDummyEventContent() DummyEventContent {
	r := DummyEventContent{}
	r.ObjectValue = nil
	return r
}

func DeserializeDummyEventContent(r io.Reader) (DummyEventContent, error) {
	t := NewDummyEventContent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDummyEventContentFromSchema(r io.Reader, schema string) (DummyEventContent, error) {
	t := NewDummyEventContent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDummyEventContent(r DummyEventContent, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.ObjectValue, w)
	if err != nil {
		return err
	}
	return err
}

func (r DummyEventContent) Serialize(w io.Writer) error {
	return writeDummyEventContent(r, w)
}

func (r DummyEventContent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"ObjectValue\",\"type\":[\"null\",\"string\"]}],\"name\":\"Skogshorn.Drift.Common.Database.Entity.DummyEventContent\",\"type\":\"record\"}"
}

func (r DummyEventContent) SchemaName() string {
	return "Skogshorn.Drift.Common.Database.Entity.DummyEventContent"
}

func (_ DummyEventContent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DummyEventContent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DummyEventContent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DummyEventContent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DummyEventContent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DummyEventContent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DummyEventContent) SetString(v string)   { panic("Unsupported operation") }
func (_ DummyEventContent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DummyEventContent) Get(i int) types.Field {
	switch i {
	case 0:
		r.ObjectValue = NewUnionNullString()

		return r.ObjectValue
	}
	panic("Unknown field index")
}

func (r *DummyEventContent) SetDefault(i int) {
	switch i {
	case 0:
		r.ObjectValue = nil
		return
	}
	panic("Unknown field index")
}

func (r *DummyEventContent) NullField(i int) {
	switch i {
	case 0:
		r.ObjectValue = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DummyEventContent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DummyEventContent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DummyEventContent) HintSize(int)                     { panic("Unsupported operation") }
func (_ DummyEventContent) Finalize()                        {}

func (_ DummyEventContent) AvroCRC64Fingerprint() []byte {
	return []byte(DummyEventContentAvroCRC64Fingerprint)
}

func (r DummyEventContent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ObjectValue"], err = json.Marshal(r.ObjectValue)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DummyEventContent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ObjectValue"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ObjectValue); err != nil {
			return err
		}
	} else {
		r.ObjectValue = NewUnionNullString()

		r.ObjectValue = nil
	}
	return nil
}
