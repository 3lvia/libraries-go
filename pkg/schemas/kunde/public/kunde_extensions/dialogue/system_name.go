// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100238_6.avsc
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

type SystemName int32

const (
	SystemNameUnknown SystemName = 0
	SystemNameElwin   SystemName = 1
)

func (e SystemName) String() string {
	switch e {
	case SystemNameUnknown:
		return "Unknown"
	case SystemNameElwin:
		return "Elwin"
	}
	return "unknown"
}

func writeSystemName(r SystemName, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewSystemNameValue(raw string) (r SystemName, err error) {
	switch raw {
	case "Unknown":
		return SystemNameUnknown, nil
	case "Elwin":
		return SystemNameElwin, nil
	}

	return -1, fmt.Errorf("invalid value for SystemName: '%s'", raw)

}

func (b SystemName) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *SystemName) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewSystemNameValue(stringVal)
	*b = val
	return err
}

type SystemNameWrapper struct {
	Target *SystemName
}

func (b SystemNameWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b SystemNameWrapper) SetInt(v int32) {
	*(b.Target) = SystemName(v)
}

func (b SystemNameWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b SystemNameWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b SystemNameWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b SystemNameWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b SystemNameWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b SystemNameWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b SystemNameWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b SystemNameWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b SystemNameWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b SystemNameWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b SystemNameWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b SystemNameWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b SystemNameWrapper) Finalize() {}
