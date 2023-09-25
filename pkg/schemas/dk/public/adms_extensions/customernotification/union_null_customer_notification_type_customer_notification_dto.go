// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100375_2.avsc
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

type UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnum int

const (
	UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnumCustomerNotificationTypeCustomerNotificationDto UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnum = 1
)

type UnionNullCustomerNotificationTypeCustomerNotificationDto struct {
	Null                                            *types.NullVal
	CustomerNotificationTypeCustomerNotificationDto CustomerNotificationTypeCustomerNotificationDto
	UnionType                                       UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnum
}

func writeUnionNullCustomerNotificationTypeCustomerNotificationDto(r *UnionNullCustomerNotificationTypeCustomerNotificationDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnumCustomerNotificationTypeCustomerNotificationDto:
		return writeCustomerNotificationTypeCustomerNotificationDto(r.CustomerNotificationTypeCustomerNotificationDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCustomerNotificationTypeCustomerNotificationDto")
}

func NewUnionNullCustomerNotificationTypeCustomerNotificationDto() *UnionNullCustomerNotificationTypeCustomerNotificationDto {
	return &UnionNullCustomerNotificationTypeCustomerNotificationDto{}
}

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) Serialize(w io.Writer) error {
	return writeUnionNullCustomerNotificationTypeCustomerNotificationDto(r, w)
}

func DeserializeUnionNullCustomerNotificationTypeCustomerNotificationDto(r io.Reader) (*UnionNullCustomerNotificationTypeCustomerNotificationDto, error) {
	t := NewUnionNullCustomerNotificationTypeCustomerNotificationDto()
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

func DeserializeUnionNullCustomerNotificationTypeCustomerNotificationDtoFromSchema(r io.Reader, schema string) (*UnionNullCustomerNotificationTypeCustomerNotificationDto, error) {
	t := NewUnionNullCustomerNotificationTypeCustomerNotificationDto()
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

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MessageText\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NotificationData\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"CancelMessage\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NotificationID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NotificationType\",\"type\":[\"null\",\"string\"]}],\"name\":\"NotificationDataDto\",\"type\":\"record\"}]},{\"name\":\"OutageEnd\",\"type\":\"string\"},{\"name\":\"OutageEndSpecified\",\"type\":\"boolean\"},{\"name\":\"OutageStart\",\"type\":\"string\"},{\"name\":\"OutageStartSpecified\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"RecordNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TechnicalAsset\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"B1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"B3\",\"type\":[\"null\",\"string\"]}],\"name\":\"TechnicalAssetDto\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"CustomerNotificationTypeCustomerNotificationDto\",\"type\":\"record\"}]"
}

func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetBoolean(v bool) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetInt(v int32) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetFloat(v float32) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetBytes(v []byte) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetString(v string) {
	panic("Unsupported operation")
}

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetLong(v int64) {

	r.UnionType = (UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnum)(v)
}

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.CustomerNotificationTypeCustomerNotificationDto = NewCustomerNotificationTypeCustomerNotificationDto()
		return &types.Record{Target: (&r.CustomerNotificationTypeCustomerNotificationDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) NullField(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) HintSize(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) SetDefault(i int) {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullCustomerNotificationTypeCustomerNotificationDto) Finalize() {}

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCustomerNotificationTypeCustomerNotificationDtoTypeEnumCustomerNotificationTypeCustomerNotificationDto:
		return json.Marshal(map[string]interface{}{"CustomerNotificationSafConsumer.CustomerNotificationTypeCustomerNotificationDto": r.CustomerNotificationTypeCustomerNotificationDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCustomerNotificationTypeCustomerNotificationDto")
}

func (r *UnionNullCustomerNotificationTypeCustomerNotificationDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["CustomerNotificationSafConsumer.CustomerNotificationTypeCustomerNotificationDto"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.CustomerNotificationTypeCustomerNotificationDto)
	}
	return fmt.Errorf("invalid value for *UnionNullCustomerNotificationTypeCustomerNotificationDto")
}
