// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100093_2.avsc
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

type JobEventStatus int32

const (
	JobEventStatusReceived JobEventStatus = 0
	JobEventStatusModified JobEventStatus = 1
	JobEventStatusStarted  JobEventStatus = 2
	JobEventStatusCanceled JobEventStatus = 3
	JobEventStatusFinished JobEventStatus = 4
)

func (e JobEventStatus) String() string {
	switch e {
	case JobEventStatusReceived:
		return "Received"
	case JobEventStatusModified:
		return "Modified"
	case JobEventStatusStarted:
		return "Started"
	case JobEventStatusCanceled:
		return "Canceled"
	case JobEventStatusFinished:
		return "Finished"
	}
	return "unknown"
}

func writeJobEventStatus(r JobEventStatus, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewJobEventStatusValue(raw string) (r JobEventStatus, err error) {
	switch raw {
	case "Received":
		return JobEventStatusReceived, nil
	case "Modified":
		return JobEventStatusModified, nil
	case "Started":
		return JobEventStatusStarted, nil
	case "Canceled":
		return JobEventStatusCanceled, nil
	case "Finished":
		return JobEventStatusFinished, nil
	}

	return -1, fmt.Errorf("invalid value for JobEventStatus: '%s'", raw)

}

func (b JobEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *JobEventStatus) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewJobEventStatusValue(stringVal)
	*b = val
	return err
}

type JobEventStatusWrapper struct {
	Target *JobEventStatus
}

func (b JobEventStatusWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b JobEventStatusWrapper) SetInt(v int32) {
	*(b.Target) = JobEventStatus(v)
}

func (b JobEventStatusWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b JobEventStatusWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b JobEventStatusWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b JobEventStatusWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b JobEventStatusWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b JobEventStatusWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b JobEventStatusWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b JobEventStatusWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b JobEventStatusWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b JobEventStatusWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b JobEventStatusWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b JobEventStatusWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b JobEventStatusWrapper) Finalize() {}
