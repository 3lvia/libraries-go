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

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NotificationInterval int32

const (
	NotificationIntervalNever   NotificationInterval = 0
	NotificationIntervalDaily   NotificationInterval = 1
	NotificationIntervalWeekly  NotificationInterval = 2
	NotificationIntervalMonthly NotificationInterval = 3
)

func (e NotificationInterval) String() string {
	switch e {
	case NotificationIntervalNever:
		return "Never"
	case NotificationIntervalDaily:
		return "Daily"
	case NotificationIntervalWeekly:
		return "Weekly"
	case NotificationIntervalMonthly:
		return "Monthly"
	}
	return "unknown"
}

func writeNotificationInterval(r NotificationInterval, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewNotificationIntervalValue(raw string) (r NotificationInterval, err error) {
	switch raw {
	case "Never":
		return NotificationIntervalNever, nil
	case "Daily":
		return NotificationIntervalDaily, nil
	case "Weekly":
		return NotificationIntervalWeekly, nil
	case "Monthly":
		return NotificationIntervalMonthly, nil
	}

	return -1, fmt.Errorf("invalid value for NotificationInterval: '%s'", raw)

}

func (b NotificationInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *NotificationInterval) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewNotificationIntervalValue(stringVal)
	*b = val
	return err
}

type NotificationIntervalWrapper struct {
	Target *NotificationInterval
}

func (b NotificationIntervalWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b NotificationIntervalWrapper) SetInt(v int32) {
	*(b.Target) = NotificationInterval(v)
}

func (b NotificationIntervalWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b NotificationIntervalWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b NotificationIntervalWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b NotificationIntervalWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b NotificationIntervalWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b NotificationIntervalWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b NotificationIntervalWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b NotificationIntervalWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b NotificationIntervalWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b NotificationIntervalWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b NotificationIntervalWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b NotificationIntervalWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b NotificationIntervalWrapper) Finalize() {}
