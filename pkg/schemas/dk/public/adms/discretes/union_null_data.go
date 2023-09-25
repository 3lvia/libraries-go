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

type UnionNullDataTypeEnum int

const (
	UnionNullDataTypeEnumData UnionNullDataTypeEnum = 1
)

type UnionNullData struct {
	Null      *types.NullVal
	Data      Data
	UnionType UnionNullDataTypeEnum
}

func writeUnionNullData(r *UnionNullData, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDataTypeEnumData:
		return writeData(r.Data, w)
	}
	return fmt.Errorf("invalid value for *UnionNullData")
}

func NewUnionNullData() *UnionNullData {
	return &UnionNullData{}
}

func (r *UnionNullData) Serialize(w io.Writer) error {
	return writeUnionNullData(r, w)
}

func DeserializeUnionNullData(r io.Reader) (*UnionNullData, error) {
	t := NewUnionNullData()
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

func DeserializeUnionNullDataFromSchema(r io.Reader, schema string) (*UnionNullData, error) {
	t := NewUnionNullData()
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

func (r *UnionNullData) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EndDeviceEvents\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"UsagePoint\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]}],\"name\":\"UsagePoint\",\"type\":\"record\"}]}],\"name\":\"EndDeviceEvent\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Operation\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"CorrelationId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MessageType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",\"string\"]}],\"name\":\"Operation\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Result\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ResultCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StatusCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"Result\",\"type\":\"record\"}]}],\"name\":\"Data\",\"type\":\"record\"}]"
}

func (_ *UnionNullData) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullData) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullData) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullData) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullData) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullData) SetLong(v int64) {

	r.UnionType = (UnionNullDataTypeEnum)(v)
}

func (r *UnionNullData) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Data = NewData()
		return &types.Record{Target: (&r.Data)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullData) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullData) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullData) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullData) Finalize()                        {}

func (r *UnionNullData) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDataTypeEnumData:
		return json.Marshal(map[string]interface{}{"SmokeTests.Kafka.Dto.TelemetryScada.Data": r.Data})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullData")
}

func (r *UnionNullData) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["SmokeTests.Kafka.Dto.TelemetryScada.Data"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Data)
	}
	return fmt.Errorf("invalid value for *UnionNullData")
}
