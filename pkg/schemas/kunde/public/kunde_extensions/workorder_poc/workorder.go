// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100182_1.avsc
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

type Workorder struct {
}

const WorkorderAvroCRC64Fingerprint = "\x05ͨ\a\x97\xef\x97\xcd"

func NewWorkorder() Workorder {
	r := Workorder{}
	return r
}

func DeserializeWorkorder(r io.Reader) (Workorder, error) {
	t := NewWorkorder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeWorkorderFromSchema(r io.Reader, schema string) (Workorder, error) {
	t := NewWorkorder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeWorkorder(r Workorder, w io.Writer) error {
	var err error
	return err
}

func (r Workorder) Serialize(w io.Writer) error {
	return writeWorkorder(r, w)
}

func (r Workorder) Schema() string {
	return "{\"fields\":[],\"name\":\"Elwin.Workorders.Integration.Model.Workorder\",\"type\":\"record\"}"
}

func (r Workorder) SchemaName() string {
	return "Elwin.Workorders.Integration.Model.Workorder"
}

func (_ Workorder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Workorder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Workorder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Workorder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Workorder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Workorder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Workorder) SetString(v string)   { panic("Unsupported operation") }
func (_ Workorder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Workorder) Get(i int) types.Field {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Workorder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Workorder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Workorder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Workorder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Workorder) HintSize(int)                     { panic("Unsupported operation") }
func (_ Workorder) Finalize()                        {}

func (_ Workorder) AvroCRC64Fingerprint() []byte {
	return []byte(WorkorderAvroCRC64Fingerprint)
}

func (r Workorder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	return json.Marshal(output)
}

func (r *Workorder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	return nil
}
