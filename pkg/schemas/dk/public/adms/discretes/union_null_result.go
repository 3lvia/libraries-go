// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100492_3.avsc
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

type UnionNullResultTypeEnum int

const (
	UnionNullResultTypeEnumResult UnionNullResultTypeEnum = 1
)

type UnionNullResult struct {
	Null      *types.NullVal
	Result    Result
	UnionType UnionNullResultTypeEnum
}

func writeUnionNullResult(r *UnionNullResult, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullResultTypeEnumResult:
		return writeResult(r.Result, w)
	}
	return fmt.Errorf("invalid value for *UnionNullResult")
}

func NewUnionNullResult() *UnionNullResult {
	return &UnionNullResult{}
}

func (r *UnionNullResult) Serialize(w io.Writer) error {
	return writeUnionNullResult(r, w)
}

func DeserializeUnionNullResult(r io.Reader) (*UnionNullResult, error) {
	t := NewUnionNullResult()
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

func DeserializeUnionNullResultFromSchema(r io.Reader, schema string) (*UnionNullResult, error) {
	t := NewUnionNullResult()
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

func (r *UnionNullResult) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ResultCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StatusCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"Result\",\"type\":\"record\"}]"
}

func (_ *UnionNullResult) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullResult) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullResult) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullResult) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullResult) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullResult) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullResult) SetLong(v int64) {

	r.UnionType = (UnionNullResultTypeEnum)(v)
}

func (r *UnionNullResult) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Result = NewResult()
		return &types.Record{Target: (&r.Result)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullResult) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullResult) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullResult) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullResult) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullResult) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullResult) Finalize()                        {}

func (r *UnionNullResult) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullResultTypeEnumResult:
		return json.Marshal(map[string]interface{}{"SmokeTests.Kafka.Dto.TelemetryScada.Result": r.Result})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullResult")
}

func (r *UnionNullResult) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SmokeTests.Kafka.Dto.TelemetryScada.Result"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Result)
	}
	return fmt.Errorf("invalid value for *UnionNullResult")
}
