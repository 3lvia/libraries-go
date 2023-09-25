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

type ServiceType int32

const (
	ServiceTypeConsumption ServiceType = 0
	ServiceTypeProduction  ServiceType = 1
	ServiceTypeCombination ServiceType = 2
	ServiceTypeExchange    ServiceType = 3
	ServiceTypeInternal    ServiceType = 4
)

func (e ServiceType) String() string {
	switch e {
	case ServiceTypeConsumption:
		return "Consumption"
	case ServiceTypeProduction:
		return "Production"
	case ServiceTypeCombination:
		return "Combination"
	case ServiceTypeExchange:
		return "Exchange"
	case ServiceTypeInternal:
		return "Internal"
	}
	return "unknown"
}

func writeServiceType(r ServiceType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewServiceTypeValue(raw string) (r ServiceType, err error) {
	switch raw {
	case "Consumption":
		return ServiceTypeConsumption, nil
	case "Production":
		return ServiceTypeProduction, nil
	case "Combination":
		return ServiceTypeCombination, nil
	case "Exchange":
		return ServiceTypeExchange, nil
	case "Internal":
		return ServiceTypeInternal, nil
	}

	return -1, fmt.Errorf("invalid value for ServiceType: '%s'", raw)

}

func (b ServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *ServiceType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewServiceTypeValue(stringVal)
	*b = val
	return err
}

type ServiceTypeWrapper struct {
	Target *ServiceType
}

func (b ServiceTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b ServiceTypeWrapper) SetInt(v int32) {
	*(b.Target) = ServiceType(v)
}

func (b ServiceTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b ServiceTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b ServiceTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b ServiceTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b ServiceTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b ServiceTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b ServiceTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b ServiceTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b ServiceTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b ServiceTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b ServiceTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b ServiceTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b ServiceTypeWrapper) Finalize() {}
