// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100254_5.avsc
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

type UnionNullNameTypeEnum int

const (
	UnionNullNameTypeEnumName UnionNullNameTypeEnum = 1
)

type UnionNullName struct {
	Null      *types.NullVal
	Name      Name
	UnionType UnionNullNameTypeEnum
}

func writeUnionNullName(r *UnionNullName, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullNameTypeEnumName:
		return writeName(r.Name, w)
	}
	return fmt.Errorf("invalid value for *UnionNullName")
}

func NewUnionNullName() *UnionNullName {
	return &UnionNullName{}
}

func (r *UnionNullName) Serialize(w io.Writer) error {
	return writeUnionNullName(r, w)
}

func DeserializeUnionNullName(r io.Reader) (*UnionNullName, error) {
	t := NewUnionNullName()
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

func DeserializeUnionNullNameFromSchema(r io.Reader, schema string) (*UnionNullName, error) {
	t := NewUnionNullName()
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

func (r *UnionNullName) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"Modified\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NameType\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"name\":\"MRID\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"name\":\"NamingAuthority\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}}],\"name\":\"NameType\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Name\",\"type\":\"record\"}]"
}

func (_ *UnionNullName) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullName) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullName) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullName) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullName) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullName) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullName) SetLong(v int64) {

	r.UnionType = (UnionNullNameTypeEnum)(v)
}

func (r *UnionNullName) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Name = NewName()
		return &types.Record{Target: (&r.Name)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullName) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullName) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullName) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullName) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullName) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullName) Finalize()                        {}

func (r *UnionNullName) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullNameTypeEnumName:
		return json.Marshal(map[string]interface{}{"UniqueId.Infrastructure.Kafka.Model.Name": r.Name})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullName")
}

func (r *UnionNullName) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["UniqueId.Infrastructure.Kafka.Model.Name"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Name)
	}
	return fmt.Errorf("invalid value for *UnionNullName")
}
