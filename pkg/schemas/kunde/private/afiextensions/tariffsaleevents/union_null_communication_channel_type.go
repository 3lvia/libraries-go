// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100346_2.avsc
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

type UnionNullCommunicationChannelTypeTypeEnum int

const (
	UnionNullCommunicationChannelTypeTypeEnumCommunicationChannelType UnionNullCommunicationChannelTypeTypeEnum = 1
)

type UnionNullCommunicationChannelType struct {
	Null                     *types.NullVal
	CommunicationChannelType CommunicationChannelType
	UnionType                UnionNullCommunicationChannelTypeTypeEnum
}

func writeUnionNullCommunicationChannelType(r *UnionNullCommunicationChannelType, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCommunicationChannelTypeTypeEnumCommunicationChannelType:
		return writeCommunicationChannelType(r.CommunicationChannelType, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCommunicationChannelType")
}

func NewUnionNullCommunicationChannelType() *UnionNullCommunicationChannelType {
	return &UnionNullCommunicationChannelType{}
}

func (r *UnionNullCommunicationChannelType) Serialize(w io.Writer) error {
	return writeUnionNullCommunicationChannelType(r, w)
}

func DeserializeUnionNullCommunicationChannelType(r io.Reader) (*UnionNullCommunicationChannelType, error) {
	t := NewUnionNullCommunicationChannelType()
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

func DeserializeUnionNullCommunicationChannelTypeFromSchema(r io.Reader, schema string) (*UnionNullCommunicationChannelType, error) {
	t := NewUnionNullCommunicationChannelType()
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

func (r *UnionNullCommunicationChannelType) Schema() string {
	return "[\"null\",{\"name\":\"CommunicationChannelType\",\"symbols\":[\"Casehandling\",\"Chat\",\"ISCustomer\",\"EMail\",\"Facebook\",\"Fax\",\"PersonalAttendance\",\"Letter\",\"Sms\",\"Phone\",\"Web\",\"ISChange\",\"Elhub\",\"Power\",\"Net\",\"DSF\",\"ISSafe\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullCommunicationChannelType) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCommunicationChannelType) SetLong(v int64) {

	r.UnionType = (UnionNullCommunicationChannelTypeTypeEnum)(v)
}

func (r *UnionNullCommunicationChannelType) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &CommunicationChannelTypeWrapper{Target: (&r.CommunicationChannelType)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCommunicationChannelType) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullCommunicationChannelType) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullCommunicationChannelType) Finalize()                {}

func (r *UnionNullCommunicationChannelType) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCommunicationChannelTypeTypeEnumCommunicationChannelType:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.TariffEvents.V1.CommunicationChannelType": r.CommunicationChannelType})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCommunicationChannelType")
}

func (r *UnionNullCommunicationChannelType) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.TariffEvents.V1.CommunicationChannelType"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.CommunicationChannelType)
	}
	return fmt.Errorf("invalid value for *UnionNullCommunicationChannelType")
}
