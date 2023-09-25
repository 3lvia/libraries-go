// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100441_1.avsc
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

type UnionNullContractTypeEnum int

const (
	UnionNullContractTypeEnumContract UnionNullContractTypeEnum = 1
)

type UnionNullContract struct {
	Null      *types.NullVal
	Contract  Contract
	UnionType UnionNullContractTypeEnum
}

func writeUnionNullContract(r *UnionNullContract, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullContractTypeEnumContract:
		return writeContract(r.Contract, w)
	}
	return fmt.Errorf("invalid value for *UnionNullContract")
}

func NewUnionNullContract() *UnionNullContract {
	return &UnionNullContract{}
}

func (r *UnionNullContract) Serialize(w io.Writer) error {
	return writeUnionNullContract(r, w)
}

func DeserializeUnionNullContract(r io.Reader) (*UnionNullContract, error) {
	t := NewUnionNullContract()
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

func DeserializeUnionNullContractFromSchema(r io.Reader, schema string) (*UnionNullContract, error) {
	t := NewUnionNullContract()
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

func (r *UnionNullContract) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EndTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Id\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"default\":null,\"name\":\"MeteringPointId\",\"type\":[\"null\",\"string\"]},{\"name\":\"StartTime\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Contract\",\"type\":\"record\"}]"
}

func (_ *UnionNullContract) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullContract) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullContract) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullContract) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullContract) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullContract) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullContract) SetLong(v int64) {

	r.UnionType = (UnionNullContractTypeEnum)(v)
}

func (r *UnionNullContract) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Contract = NewContract()
		return &types.Record{Target: (&r.Contract)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullContract) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullContract) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullContract) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullContract) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullContract) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullContract) Finalize()                        {}

func (r *UnionNullContract) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullContractTypeEnumContract:
		return json.Marshal(map[string]interface{}{"Kunde.Schema.Messages.Contract": r.Contract})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullContract")
}

func (r *UnionNullContract) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Kunde.Schema.Messages.Contract"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Contract)
	}
	return fmt.Errorf("invalid value for *UnionNullContract")
}
