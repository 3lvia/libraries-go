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

type UtilityType int32

const (
	UtilityTypeElectricity UtilityType = 0
	UtilityTypeWater       UtilityType = 1
	UtilityTypeGas         UtilityType = 2
	UtilityTypeHeat        UtilityType = 3
)

func (e UtilityType) String() string {
	switch e {
	case UtilityTypeElectricity:
		return "Electricity"
	case UtilityTypeWater:
		return "Water"
	case UtilityTypeGas:
		return "Gas"
	case UtilityTypeHeat:
		return "Heat"
	}
	return "unknown"
}

func writeUtilityType(r UtilityType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewUtilityTypeValue(raw string) (r UtilityType, err error) {
	switch raw {
	case "Electricity":
		return UtilityTypeElectricity, nil
	case "Water":
		return UtilityTypeWater, nil
	case "Gas":
		return UtilityTypeGas, nil
	case "Heat":
		return UtilityTypeHeat, nil
	}

	return -1, fmt.Errorf("invalid value for UtilityType: '%s'", raw)

}

func (b UtilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *UtilityType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewUtilityTypeValue(stringVal)
	*b = val
	return err
}

type UtilityTypeWrapper struct {
	Target *UtilityType
}

func (b UtilityTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b UtilityTypeWrapper) SetInt(v int32) {
	*(b.Target) = UtilityType(v)
}

func (b UtilityTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b UtilityTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b UtilityTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b UtilityTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b UtilityTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b UtilityTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b UtilityTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b UtilityTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b UtilityTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b UtilityTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b UtilityTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b UtilityTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b UtilityTypeWrapper) Finalize() {}
