// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100250_7.avsc
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

type Direction int32

const (
	DirectionUnknown Direction = 0
	DirectionIn      Direction = 1
	DirectionOut     Direction = 2
)

func (e Direction) String() string {
	switch e {
	case DirectionUnknown:
		return "Unknown"
	case DirectionIn:
		return "In"
	case DirectionOut:
		return "Out"
	}
	return "unknown"
}

func writeDirection(r Direction, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewDirectionValue(raw string) (r Direction, err error) {
	switch raw {
	case "Unknown":
		return DirectionUnknown, nil
	case "In":
		return DirectionIn, nil
	case "Out":
		return DirectionOut, nil
	}

	return -1, fmt.Errorf("invalid value for Direction: '%s'", raw)

}

func (b Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Direction) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewDirectionValue(stringVal)
	*b = val
	return err
}

type DirectionWrapper struct {
	Target *Direction
}

func (b DirectionWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b DirectionWrapper) SetInt(v int32) {
	*(b.Target) = Direction(v)
}

func (b DirectionWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b DirectionWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b DirectionWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b DirectionWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b DirectionWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b DirectionWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b DirectionWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b DirectionWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b DirectionWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b DirectionWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b DirectionWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b DirectionWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b DirectionWrapper) Finalize() {}
