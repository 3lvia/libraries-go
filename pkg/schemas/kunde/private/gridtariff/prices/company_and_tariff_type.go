// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100242_2.avsc
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

var _ = fmt.Printf

type CompanyAndTariffType struct {
	Company *UnionNullCompany `json:"Company"`

	TariffTypes *UnionNullArrayUnionNullTariffType `json:"TariffTypes"`
}

const CompanyAndTariffTypeAvroCRC64Fingerprint = "\x91;TJh\xc1\xaef"

func NewCompanyAndTariffType() CompanyAndTariffType {
	r := CompanyAndTariffType{}
	r.Company = nil
	r.TariffTypes = nil
	return r
}

func DeserializeCompanyAndTariffType(r io.Reader) (CompanyAndTariffType, error) {
	t := NewCompanyAndTariffType()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCompanyAndTariffTypeFromSchema(r io.Reader, schema string) (CompanyAndTariffType, error) {
	t := NewCompanyAndTariffType()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCompanyAndTariffType(r CompanyAndTariffType, w io.Writer) error {
	var err error
	err = writeUnionNullCompany(r.Company, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayUnionNullTariffType(r.TariffTypes, w)
	if err != nil {
		return err
	}
	return err
}

func (r CompanyAndTariffType) Serialize(w io.Writer) error {
	return writeCompanyAndTariffType(r, w)
}

func (r CompanyAndTariffType) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Company\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"CompanyName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CompanyOrgNo\",\"type\":[\"null\",\"string\"]}],\"name\":\"Company\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"TariffTypes\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"ConsumptionFlag\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FixedPriceConfiguration\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AdditionalProperties\",\"type\":[\"null\",{\"type\":\"map\",\"values\":[\"null\",{\"fields\":[],\"name\":\"Object\",\"namespace\":\"System\",\"type\":\"record\"}]}]},{\"default\":null,\"name\":\"AllDaysPerMonth\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Basis\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DaysPerMonth\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MaxhoursPerDay\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MaxhoursPerMonth\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Months\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MonthsOffset\",\"type\":[\"null\",\"int\"]}],\"name\":\"FixedPriceConfiguration\",\"type\":\"record\"}]},{\"name\":\"LastUpdated\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"PowerPriceConfiguration\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AdditionalProperties\",\"type\":[\"null\",{\"type\":\"map\",\"values\":[\"null\",\"System.Object\"]}]},{\"default\":null,\"name\":\"PowerFactorPercentage\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ReactivePowerPricing\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"PowerPriceConfiguration\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Product\",\"type\":[\"null\",\"string\"]},{\"name\":\"Resolution\",\"type\":\"int\"},{\"default\":null,\"name\":\"TariffKey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TariffPrices\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"FixedPrices\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"FixedPriceLevel\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Currency\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LevelInfo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MonetaryUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"name\":\"MonthlyFixedExTaxes\",\"type\":\"double\"},{\"default\":null,\"name\":\"MonthlyUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdDown\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdUp\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValueMax\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueMin\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueUnitOfMeasure\",\"type\":[\"null\",\"string\"]}],\"name\":\"FixedPriceLevel\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]}],\"name\":\"FixedPrices\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Seasons\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EnergyPrice\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Currency\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EnergyPriceLevel\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EnergyExTaxes\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"Hours\",\"type\":[\"null\",{\"items\":\"int\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Level\",\"type\":[\"null\",\"string\"]}],\"name\":\"EnergyPrice\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"MonetaryUnitOfMeasure\",\"type\":[\"null\",\"string\"]}],\"name\":\"EnergyPrice2\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Months\",\"type\":[\"null\",{\"items\":\"int\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PowerPrices\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"PowerPriceLevel\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Currency\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LevelInfo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MonetaryUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"name\":\"MonthlyActivePowerExTaxes\",\"type\":\"double\"},{\"default\":null,\"name\":\"MonthlyReactivePowerExTaxes\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"MonthlyUnitOfMeasure\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdDown\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NextIdUp\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValueMax\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueMin\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ValueUnitOfMeasure\",\"type\":[\"null\",\"string\"]}],\"name\":\"PowerPrice\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"PowerPrice2\",\"type\":\"record\"}]}],\"name\":\"Season\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"Taxes\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"EnergyPriceTaxes\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"TaxType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxTypeDescription\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxUom\",\"type\":[\"null\",\"string\"]},{\"name\":\"TaxValue\",\"type\":\"double\"}],\"name\":\"TaxConfigEnergy\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"FixedPriceTaxes\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"TaxType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxTypeDescription\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxUom\",\"type\":[\"null\",\"string\"]},{\"name\":\"TaxValue\",\"type\":\"double\"}],\"name\":\"TaxConfigFixed\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"PowerPriceTaxes\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"TaxType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxTypeDescription\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TaxUom\",\"type\":[\"null\",\"string\"]},{\"name\":\"TaxValue\",\"type\":\"double\"}],\"name\":\"TaxConfigPower\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"Taxes\",\"type\":\"record\"}]}],\"name\":\"TariffPrice\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"Title\",\"type\":[\"null\",\"string\"]},{\"name\":\"UsePowerPriceConfiguration\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"UsePublicHolidayOverride\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UseWeekendPriceOverride\",\"type\":[\"null\",\"string\"]}],\"name\":\"TariffType\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"GridTariff.Events.V1.CompanyAndTariffType\",\"type\":\"record\"}"
}

func (r CompanyAndTariffType) SchemaName() string {
	return "GridTariff.Events.V1.CompanyAndTariffType"
}

func (_ CompanyAndTariffType) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetString(v string)   { panic("Unsupported operation") }
func (_ CompanyAndTariffType) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CompanyAndTariffType) Get(i int) types.Field {
	switch i {
	case 0:
		r.Company = NewUnionNullCompany()

		return r.Company
	case 1:
		r.TariffTypes = NewUnionNullArrayUnionNullTariffType()

		return r.TariffTypes
	}
	panic("Unknown field index")
}

func (r *CompanyAndTariffType) SetDefault(i int) {
	switch i {
	case 0:
		r.Company = nil
		return
	case 1:
		r.TariffTypes = nil
		return
	}
	panic("Unknown field index")
}

func (r *CompanyAndTariffType) NullField(i int) {
	switch i {
	case 0:
		r.Company = nil
		return
	case 1:
		r.TariffTypes = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CompanyAndTariffType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CompanyAndTariffType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CompanyAndTariffType) HintSize(int)                     { panic("Unsupported operation") }
func (_ CompanyAndTariffType) Finalize()                        {}

func (_ CompanyAndTariffType) AvroCRC64Fingerprint() []byte {
	return []byte(CompanyAndTariffTypeAvroCRC64Fingerprint)
}

func (r CompanyAndTariffType) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Company"], err = json.Marshal(r.Company)
	if err != nil {
		return nil, err
	}
	output["TariffTypes"], err = json.Marshal(r.TariffTypes)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CompanyAndTariffType) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Company"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Company); err != nil {
			return err
		}
	} else {
		r.Company = NewUnionNullCompany()

		r.Company = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TariffTypes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TariffTypes); err != nil {
			return err
		}
	} else {
		r.TariffTypes = NewUnionNullArrayUnionNullTariffType()

		r.TariffTypes = nil
	}
	return nil
}
