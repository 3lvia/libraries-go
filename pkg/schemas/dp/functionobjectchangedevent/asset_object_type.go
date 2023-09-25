// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100135_2.avsc
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

type AssetObjectType int32

const (
	AssetObjectTypeAntenne AssetObjectType = 0
	AssetObjectTypeMaaler  AssetObjectType = 1
	AssetObjectTypeFake    AssetObjectType = 2
)

func (e AssetObjectType) String() string {
	switch e {
	case AssetObjectTypeAntenne:
		return "Antenne"
	case AssetObjectTypeMaaler:
		return "Maaler"
	case AssetObjectTypeFake:
		return "Fake"
	}
	return "unknown"
}

func writeAssetObjectType(r AssetObjectType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewAssetObjectTypeValue(raw string) (r AssetObjectType, err error) {
	switch raw {
	case "Antenne":
		return AssetObjectTypeAntenne, nil
	case "Maaler":
		return AssetObjectTypeMaaler, nil
	case "Fake":
		return AssetObjectTypeFake, nil
	}

	return -1, fmt.Errorf("invalid value for AssetObjectType: '%s'", raw)

}

func (b AssetObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *AssetObjectType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewAssetObjectTypeValue(stringVal)
	*b = val
	return err
}

type AssetObjectTypeWrapper struct {
	Target *AssetObjectType
}

func (b AssetObjectTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b AssetObjectTypeWrapper) SetInt(v int32) {
	*(b.Target) = AssetObjectType(v)
}

func (b AssetObjectTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b AssetObjectTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b AssetObjectTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b AssetObjectTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b AssetObjectTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b AssetObjectTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b AssetObjectTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b AssetObjectTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b AssetObjectTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b AssetObjectTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b AssetObjectTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b AssetObjectTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b AssetObjectTypeWrapper) Finalize() {}
