// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100366_4.avsc
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

type MonitoringStatus int32

const (
	MonitoringStatusOnTheWatchList     MonitoringStatus = 0
	MonitoringStatusApprovedForBilling MonitoringStatus = 1
	MonitoringStatusNotOnTheWatchList  MonitoringStatus = 2
)

func (e MonitoringStatus) String() string {
	switch e {
	case MonitoringStatusOnTheWatchList:
		return "OnTheWatchList"
	case MonitoringStatusApprovedForBilling:
		return "ApprovedForBilling"
	case MonitoringStatusNotOnTheWatchList:
		return "NotOnTheWatchList"
	}
	return "unknown"
}

func writeMonitoringStatus(r MonitoringStatus, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewMonitoringStatusValue(raw string) (r MonitoringStatus, err error) {
	switch raw {
	case "OnTheWatchList":
		return MonitoringStatusOnTheWatchList, nil
	case "ApprovedForBilling":
		return MonitoringStatusApprovedForBilling, nil
	case "NotOnTheWatchList":
		return MonitoringStatusNotOnTheWatchList, nil
	}

	return -1, fmt.Errorf("invalid value for MonitoringStatus: '%s'", raw)

}

func (b MonitoringStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *MonitoringStatus) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewMonitoringStatusValue(stringVal)
	*b = val
	return err
}

type MonitoringStatusWrapper struct {
	Target *MonitoringStatus
}

func (b MonitoringStatusWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b MonitoringStatusWrapper) SetInt(v int32) {
	*(b.Target) = MonitoringStatus(v)
}

func (b MonitoringStatusWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b MonitoringStatusWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b MonitoringStatusWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b MonitoringStatusWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b MonitoringStatusWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b MonitoringStatusWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b MonitoringStatusWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b MonitoringStatusWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b MonitoringStatusWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b MonitoringStatusWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b MonitoringStatusWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b MonitoringStatusWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b MonitoringStatusWrapper) Finalize() {}
