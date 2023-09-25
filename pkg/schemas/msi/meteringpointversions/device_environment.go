// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100058_21.avsc
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

type DeviceEnvironment int32

const (
	DeviceEnvironmentUnknown         DeviceEnvironment = 0
	DeviceEnvironmentT03             DeviceEnvironment = 1
	DeviceEnvironmentT05             DeviceEnvironment = 2
	DeviceEnvironmentStagingSouth    DeviceEnvironment = 3
	DeviceEnvironmentProductionSouth DeviceEnvironment = 4
	DeviceEnvironmentTestNorth       DeviceEnvironment = 5
	DeviceEnvironmentStagingNorth    DeviceEnvironment = 6
	DeviceEnvironmentProductionNorth DeviceEnvironment = 7
)

func (e DeviceEnvironment) String() string {
	switch e {
	case DeviceEnvironmentUnknown:
		return "Unknown"
	case DeviceEnvironmentT03:
		return "T03"
	case DeviceEnvironmentT05:
		return "T05"
	case DeviceEnvironmentStagingSouth:
		return "StagingSouth"
	case DeviceEnvironmentProductionSouth:
		return "ProductionSouth"
	case DeviceEnvironmentTestNorth:
		return "TestNorth"
	case DeviceEnvironmentStagingNorth:
		return "StagingNorth"
	case DeviceEnvironmentProductionNorth:
		return "ProductionNorth"
	}
	return "unknown"
}

func writeDeviceEnvironment(r DeviceEnvironment, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewDeviceEnvironmentValue(raw string) (r DeviceEnvironment, err error) {
	switch raw {
	case "Unknown":
		return DeviceEnvironmentUnknown, nil
	case "T03":
		return DeviceEnvironmentT03, nil
	case "T05":
		return DeviceEnvironmentT05, nil
	case "StagingSouth":
		return DeviceEnvironmentStagingSouth, nil
	case "ProductionSouth":
		return DeviceEnvironmentProductionSouth, nil
	case "TestNorth":
		return DeviceEnvironmentTestNorth, nil
	case "StagingNorth":
		return DeviceEnvironmentStagingNorth, nil
	case "ProductionNorth":
		return DeviceEnvironmentProductionNorth, nil
	}

	return -1, fmt.Errorf("invalid value for DeviceEnvironment: '%s'", raw)

}

func (b DeviceEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *DeviceEnvironment) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewDeviceEnvironmentValue(stringVal)
	*b = val
	return err
}

type DeviceEnvironmentWrapper struct {
	Target *DeviceEnvironment
}

func (b DeviceEnvironmentWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b DeviceEnvironmentWrapper) SetInt(v int32) {
	*(b.Target) = DeviceEnvironment(v)
}

func (b DeviceEnvironmentWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b DeviceEnvironmentWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b DeviceEnvironmentWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b DeviceEnvironmentWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b DeviceEnvironmentWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b DeviceEnvironmentWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b DeviceEnvironmentWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b DeviceEnvironmentWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b DeviceEnvironmentWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b DeviceEnvironmentWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b DeviceEnvironmentWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b DeviceEnvironmentWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b DeviceEnvironmentWrapper) Finalize() {}
