// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100444_4.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullProductTypeTypeEnum int

const (
	UnionNullProductTypeTypeEnumProductType UnionNullProductTypeTypeEnum = 1
)

type UnionNullProductType struct {
	Null        *types.NullVal
	ProductType ProductType
	UnionType   UnionNullProductTypeTypeEnum
}

func writeUnionNullProductType(r *UnionNullProductType, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullProductTypeTypeEnumProductType:
		return writeProductType(r.ProductType, w)
	}
	return fmt.Errorf("invalid value for *UnionNullProductType")
}

func NewUnionNullProductType() *UnionNullProductType {
	return &UnionNullProductType{}
}

func (r *UnionNullProductType) Serialize(w io.Writer) error {
	return writeUnionNullProductType(r, w)
}

func DeserializeUnionNullProductType(r io.Reader) (*UnionNullProductType, error) {
	t := NewUnionNullProductType()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullProductTypeFromSchema(r io.Reader, schema string) (*UnionNullProductType, error) {
	t := NewUnionNullProductType()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullProductType) Schema() string {
	return "[\"null\",{\"name\":\"ProductType\",\"symbols\":[\"ActiveEnergy\",\"ActivePower\",\"ReactiveEnergy\",\"ReactivePower\",\"DistrictHeatingEnergy\",\"DistrictHeatingPower\",\"DistrictHeatingWaterVolume\",\"DistrictHeatingWaterDelivered\",\"NaturalGasEnergy\",\"NaturalGasVolume\",\"NaturalGasNormalizedVolume\",\"CapacitiveReactiveEnergy\",\"CapacitiveReactivePower\",\"InductiveReactiveEnergy\",\"InductiveReactivePower\",\"ActualProduction\",\"PlannedProduction\",\"ReviewedProductionInElSpot\",\"ActualConsumption\",\"ReviewedConsumptionInElSpot\",\"ConcessionPower\",\"ReplacementPower\",\"IndustryContracts\",\"RKImbalanceAsBalanceResponsibleForGridOwner\",\"EnergyPart\",\"MarginalLossRates\",\"NaturalGasControlVolume\",\"NaturalGasCorrectedVolume\",\"ActiveEnergyPPC\",\"ActiveEnergyFPC\",\"DistrictHeatingSteppedMeasured\",\"DistrictHeatingTemperatureInlet\",\"DistrictHeatingTemperatureOutlet\",\"DistrictHeatingVolBasedAvgTempIn\",\"DistrictHeatingVolBasedAvgTempOut\",\"DistrictCoolingEnergy\",\"DistrictCoolingWaterVolume\",\"DistrictCoolingTemperatureInlet\",\"DistrictCoolingTemperatureOutlet\",\"DistrictCoolingVolBasedAvgTempIn\",\"DistrictCoolingVolBasedAvgTempOut\",\"InvoiceShareElectricity\",\"InvoiceShareCooling\",\"InvoiceShareGasEnergy\",\"InvoiceShareGasVolume\",\"DistrictHeatingVolumeBasedAverageTemperatureInlet\",\"DistrictHeatingVolumeBasedAverageTemperatureOutlet\",\"DistrictCoolingVolumeBasedAverageTemperatureInlet\",\"DistrictCoolingVolumeBasedAverageTemperatureOutlet\",\"ActiveEnergyControl\",\"CustomProduct\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullProductType) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullProductType) SetLong(v int64) {

	r.UnionType = (UnionNullProductTypeTypeEnum)(v)
}

func (r *UnionNullProductType) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &ProductTypeWrapper{Target: (&r.ProductType)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullProductType) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullProductType) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullProductType) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullProductType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullProductType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullProductType) Finalize()                        {}

func (r *UnionNullProductType) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullProductTypeTypeEnumProductType:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.TariffEvents.V1.ProductType": r.ProductType})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullProductType")
}

func (r *UnionNullProductType) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.TariffEvents.V1.ProductType"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ProductType)
	}
	return fmt.Errorf("invalid value for *UnionNullProductType")
}
