// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100159_3.avsc
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

type UnionNullArrayUnionNullCommunicationTypeEnum int

const (
	UnionNullArrayUnionNullCommunicationTypeEnumArrayUnionNullCommunication UnionNullArrayUnionNullCommunicationTypeEnum = 1
)

type UnionNullArrayUnionNullCommunication struct {
	Null                        *types.NullVal
	ArrayUnionNullCommunication []*UnionNullCommunication
	UnionType                   UnionNullArrayUnionNullCommunicationTypeEnum
}

func writeUnionNullArrayUnionNullCommunication(r *UnionNullArrayUnionNullCommunication, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullCommunicationTypeEnumArrayUnionNullCommunication:
		return writeArrayUnionNullCommunication(r.ArrayUnionNullCommunication, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullCommunication")
}

func NewUnionNullArrayUnionNullCommunication() *UnionNullArrayUnionNullCommunication {
	return &UnionNullArrayUnionNullCommunication{}
}

func (r *UnionNullArrayUnionNullCommunication) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullCommunication(r, w)
}

func DeserializeUnionNullArrayUnionNullCommunication(r io.Reader) (*UnionNullArrayUnionNullCommunication, error) {
	t := NewUnionNullArrayUnionNullCommunication()
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

func DeserializeUnionNullArrayUnionNullCommunicationFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullCommunication, error) {
	t := NewUnionNullArrayUnionNullCommunication()
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

func (r *UnionNullArrayUnionNullCommunication) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"CommucationType\",\"type\":{\"name\":\"CommunicationType\",\"symbols\":[\"Email\",\"Phone\"],\"type\":\"enum\"}},{\"name\":\"CommunicationId\",\"type\":\"int\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Communication\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullCommunication) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayUnionNullCommunication) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullCommunicationTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullCommunication) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullCommunication = make([]*UnionNullCommunication, 0)
		return &ArrayUnionNullCommunicationWrapper{Target: (&r.ArrayUnionNullCommunication)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullCommunication) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullCommunication) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullCommunication) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullCommunication) Finalize() {}

func (r *UnionNullArrayUnionNullCommunication) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullCommunicationTypeEnumArrayUnionNullCommunication:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullCommunication})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullCommunication")
}

func (r *UnionNullArrayUnionNullCommunication) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullCommunication)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullCommunication")
}
