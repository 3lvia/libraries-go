// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100450_1.avsc
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

type UnionNullSummaryTypeEnum int

const (
	UnionNullSummaryTypeEnumSummary UnionNullSummaryTypeEnum = 1
)

type UnionNullSummary struct {
	Null      *types.NullVal
	Summary   Summary
	UnionType UnionNullSummaryTypeEnum
}

func writeUnionNullSummary(r *UnionNullSummary, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullSummaryTypeEnumSummary:
		return writeSummary(r.Summary, w)
	}
	return fmt.Errorf("invalid value for *UnionNullSummary")
}

func NewUnionNullSummary() *UnionNullSummary {
	return &UnionNullSummary{}
}

func (r *UnionNullSummary) Serialize(w io.Writer) error {
	return writeUnionNullSummary(r, w)
}

func DeserializeUnionNullSummary(r io.Reader) (*UnionNullSummary, error) {
	t := NewUnionNullSummary()
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

func DeserializeUnionNullSummaryFromSchema(r io.Reader, schema string) (*UnionNullSummary, error) {
	t := NewUnionNullSummary()
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

func (r *UnionNullSummary) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"criticalCount\",\"type\":\"int\"},{\"name\":\"totalCount\",\"type\":\"int\"}],\"name\":\"Summary\",\"type\":\"record\"}]"
}

func (_ *UnionNullSummary) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullSummary) SetLong(v int64) {

	r.UnionType = (UnionNullSummaryTypeEnum)(v)
}

func (r *UnionNullSummary) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Summary = NewSummary()
		return &types.Record{Target: (&r.Summary)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullSummary) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullSummary) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullSummary) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullSummary) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullSummary) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullSummary) Finalize()                        {}

func (r *UnionNullSummary) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullSummaryTypeEnumSummary:
		return json.Marshal(map[string]interface{}{"OutageEventsLvisToElvia.Kafka.Dto.Summary": r.Summary})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullSummary")
}

func (r *UnionNullSummary) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["OutageEventsLvisToElvia.Kafka.Dto.Summary"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Summary)
	}
	return fmt.Errorf("invalid value for *UnionNullSummary")
}
