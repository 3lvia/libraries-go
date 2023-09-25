// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100423_3.avsc
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

type UnionNullMessageStatusTypeEnum int

const (
	UnionNullMessageStatusTypeEnumMessageStatus UnionNullMessageStatusTypeEnum = 1
)

type UnionNullMessageStatus struct {
	Null          *types.NullVal
	MessageStatus MessageStatus
	UnionType     UnionNullMessageStatusTypeEnum
}

func writeUnionNullMessageStatus(r *UnionNullMessageStatus, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullMessageStatusTypeEnumMessageStatus:
		return writeMessageStatus(r.MessageStatus, w)
	}
	return fmt.Errorf("invalid value for *UnionNullMessageStatus")
}

func NewUnionNullMessageStatus() *UnionNullMessageStatus {
	return &UnionNullMessageStatus{}
}

func (r *UnionNullMessageStatus) Serialize(w io.Writer) error {
	return writeUnionNullMessageStatus(r, w)
}

func DeserializeUnionNullMessageStatus(r io.Reader) (*UnionNullMessageStatus, error) {
	t := NewUnionNullMessageStatus()
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

func DeserializeUnionNullMessageStatusFromSchema(r io.Reader, schema string) (*UnionNullMessageStatus, error) {
	t := NewUnionNullMessageStatus()
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

func (r *UnionNullMessageStatus) Schema() string {
	return "[\"null\",{\"name\":\"MessageStatus\",\"namespace\":\"AdmsExtensionsLibMessagetracker.Enums\",\"symbols\":[\"CouldNotSendToSaf\",\"SentToSaf\",\"ReceivedOkFromSaf\",\"ReceivedErrorFromSaf\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullMessageStatus) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullMessageStatus) SetLong(v int64) {

	r.UnionType = (UnionNullMessageStatusTypeEnum)(v)
}

func (r *UnionNullMessageStatus) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &MessageStatusWrapper{Target: (&r.MessageStatus)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullMessageStatus) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullMessageStatus) Finalize()                        {}

func (r *UnionNullMessageStatus) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullMessageStatusTypeEnumMessageStatus:
		return json.Marshal(map[string]interface{}{"AdmsExtensionsLibMessagetracker.Enums.MessageStatus": r.MessageStatus})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullMessageStatus")
}

func (r *UnionNullMessageStatus) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["AdmsExtensionsLibMessagetracker.Enums.MessageStatus"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.MessageStatus)
	}
	return fmt.Errorf("invalid value for *UnionNullMessageStatus")
}
