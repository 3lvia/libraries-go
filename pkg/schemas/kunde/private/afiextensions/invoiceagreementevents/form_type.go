// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100398_5.avsc
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

type FormType int32

const (
	FormTypeBankgiro   FormType = 0
	FormTypeAvtalegiro FormType = 1
)

func (e FormType) String() string {
	switch e {
	case FormTypeBankgiro:
		return "Bankgiro"
	case FormTypeAvtalegiro:
		return "Avtalegiro"
	}
	return "unknown"
}

func writeFormType(r FormType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewFormTypeValue(raw string) (r FormType, err error) {
	switch raw {
	case "Bankgiro":
		return FormTypeBankgiro, nil
	case "Avtalegiro":
		return FormTypeAvtalegiro, nil
	}

	return -1, fmt.Errorf("invalid value for FormType: '%s'", raw)

}

func (b FormType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *FormType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewFormTypeValue(stringVal)
	*b = val
	return err
}

type FormTypeWrapper struct {
	Target *FormType
}

func (b FormTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b FormTypeWrapper) SetInt(v int32) {
	*(b.Target) = FormType(v)
}

func (b FormTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b FormTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b FormTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b FormTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b FormTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b FormTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b FormTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b FormTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b FormTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b FormTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b FormTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b FormTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b FormTypeWrapper) Finalize() {}
