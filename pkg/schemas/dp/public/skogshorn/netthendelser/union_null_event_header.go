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

type UnionNullEventHeaderTypeEnum int

const (
	UnionNullEventHeaderTypeEnumEventHeader UnionNullEventHeaderTypeEnum = 1
)

type UnionNullEventHeader struct {
	Null        *types.NullVal
	EventHeader EventHeader
	UnionType   UnionNullEventHeaderTypeEnum
}

func writeUnionNullEventHeader(r *UnionNullEventHeader, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullEventHeaderTypeEnumEventHeader:
		return writeEventHeader(r.EventHeader, w)
	}
	return fmt.Errorf("invalid value for *UnionNullEventHeader")
}

func NewUnionNullEventHeader() *UnionNullEventHeader {
	return &UnionNullEventHeader{}
}

func (r *UnionNullEventHeader) Serialize(w io.Writer) error {
	return writeUnionNullEventHeader(r, w)
}

func DeserializeUnionNullEventHeader(r io.Reader) (*UnionNullEventHeader, error) {
	t := NewUnionNullEventHeader()
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

func DeserializeUnionNullEventHeaderFromSchema(r io.Reader, schema string) (*UnionNullEventHeader, error) {
	t := NewUnionNullEventHeader()
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

func (r *UnionNullEventHeader) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EventId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EventName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Source\",\"type\":[\"null\",\"string\"]},{\"name\":\"TimeStamp\",\"type\":\"string\"}],\"name\":\"EventHeader\",\"type\":\"record\"}]"
}

func (_ *UnionNullEventHeader) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullEventHeader) SetLong(v int64) {

	r.UnionType = (UnionNullEventHeaderTypeEnum)(v)
}

func (r *UnionNullEventHeader) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.EventHeader = NewEventHeader()
		return &types.Record{Target: (&r.EventHeader)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullEventHeader) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullEventHeader) Finalize()                        {}

func (r *UnionNullEventHeader) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullEventHeaderTypeEnumEventHeader:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Schemas.Events.EventHeader": r.EventHeader})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullEventHeader")
}

func (r *UnionNullEventHeader) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Schemas.Events.EventHeader"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.EventHeader)
	}
	return fmt.Errorf("invalid value for *UnionNullEventHeader")
}
