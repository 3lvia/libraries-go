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

var _ = fmt.Printf

type Counter struct {
	DeliveryPointCounterId *UnionNullInt `json:"DeliveryPointCounterId"`

	Directionn *UnionNullDirection `json:"Directionn"`

	FromDate *UnionNullString `json:"FromDate"`

	Id *UnionNullString `json:"Id"`

	LastReadingDate *UnionNullString `json:"LastReadingDate"`

	NumberOfDecimals *UnionNullInt `json:"NumberOfDecimals"`

	NumberOfDigits *UnionNullInt `json:"NumberOfDigits"`

	PendingApproval *UnionNullBool `json:"PendingApproval"`

	Product *UnionNullProductType `json:"Product"`

	ReadingType *UnionNullReadingType `json:"ReadingType"`

	ToDate *UnionNullString `json:"ToDate"`

	Unit *UnionNullString `json:"Unit"`

	YearlyConsumption *UnionNullInt `json:"YearlyConsumption"`
}

const CounterAvroCRC64Fingerprint = "\xe2\xaex\xe4\xac\x15r)"

func NewCounter() Counter {
	r := Counter{}
	r.DeliveryPointCounterId = nil
	r.Directionn = nil
	r.FromDate = nil
	r.Id = nil
	r.LastReadingDate = nil
	r.NumberOfDecimals = nil
	r.NumberOfDigits = nil
	r.PendingApproval = nil
	r.Product = nil
	r.ReadingType = nil
	r.ToDate = nil
	r.Unit = nil
	r.YearlyConsumption = nil
	return r
}

func DeserializeCounter(r io.Reader) (Counter, error) {
	t := NewCounter()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCounterFromSchema(r io.Reader, schema string) (Counter, error) {
	t := NewCounter()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCounter(r Counter, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.DeliveryPointCounterId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDirection(r.Directionn, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FromDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LastReadingDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.NumberOfDecimals, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.NumberOfDigits, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.PendingApproval, w)
	if err != nil {
		return err
	}
	err = writeUnionNullProductType(r.Product, w)
	if err != nil {
		return err
	}
	err = writeUnionNullReadingType(r.ReadingType, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ToDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Unit, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.YearlyConsumption, w)
	if err != nil {
		return err
	}
	return err
}

func (r Counter) Serialize(w io.Writer) error {
	return writeCounter(r, w)
}

func (r Counter) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"DeliveryPointCounterId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Directionn\",\"type\":[\"null\",{\"name\":\"Direction\",\"symbols\":[\"Production\",\"Consumption\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"FromDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastReadingDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumberOfDecimals\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"NumberOfDigits\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"PendingApproval\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Product\",\"type\":[\"null\",{\"name\":\"ProductType\",\"symbols\":[\"ActiveEnergy\",\"ActivePower\",\"ReactiveEnergy\",\"ReactivePower\",\"DistrictHeatingEnergy\",\"DistrictHeatingPower\",\"DistrictHeatingWaterVolume\",\"DistrictHeatingWaterDelivered\",\"NaturalGasEnergy\",\"NaturalGasVolume\",\"NaturalGasNormalizedVolume\",\"CapacitiveReactiveEnergy\",\"CapacitiveReactivePower\",\"InductiveReactiveEnergy\",\"InductiveReactivePower\",\"ActualProduction\",\"PlannedProduction\",\"ReviewedProductionInElSpot\",\"ActualConsumption\",\"ReviewedConsumptionInElSpot\",\"ConcessionPower\",\"ReplacementPower\",\"IndustryContracts\",\"RKImbalanceAsBalanceResponsibleForGridOwner\",\"EnergyPart\",\"MarginalLossRates\",\"NaturalGasControlVolume\",\"NaturalGasCorrectedVolume\",\"ActiveEnergyPPC\",\"ActiveEnergyFPC\",\"DistrictHeatingSteppedMeasured\",\"DistrictHeatingTemperatureInlet\",\"DistrictHeatingTemperatureOutlet\",\"DistrictHeatingVolBasedAvgTempIn\",\"DistrictHeatingVolBasedAvgTempOut\",\"DistrictCoolingEnergy\",\"DistrictCoolingWaterVolume\",\"DistrictCoolingTemperatureInlet\",\"DistrictCoolingTemperatureOutlet\",\"DistrictCoolingVolBasedAvgTempIn\",\"DistrictCoolingVolBasedAvgTempOut\",\"InvoiceShareElectricity\",\"InvoiceShareCooling\",\"InvoiceShareGasEnergy\",\"InvoiceShareGasVolume\",\"DistrictHeatingVolumeBasedAverageTemperatureInlet\",\"DistrictHeatingVolumeBasedAverageTemperatureOutlet\",\"DistrictCoolingVolumeBasedAverageTemperatureInlet\",\"DistrictCoolingVolumeBasedAverageTemperatureOutlet\",\"ActiveEnergyControl\",\"CustomProduct\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"ReadingType\",\"type\":[\"null\",{\"name\":\"ReadingType\",\"symbols\":[\"Hourly\",\"Manual\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"ToDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unit\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"YearlyConsumption\",\"type\":[\"null\",\"int\"]}],\"name\":\"Afiextensions.Schemas.TariffEvents.V1.Counter\",\"type\":\"record\"}"
}

func (r Counter) SchemaName() string {
	return "Afiextensions.Schemas.TariffEvents.V1.Counter"
}

func (_ Counter) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Counter) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Counter) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Counter) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Counter) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Counter) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Counter) SetString(v string)   { panic("Unsupported operation") }
func (_ Counter) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Counter) Get(i int) types.Field {
	switch i {
	case 0:
		r.DeliveryPointCounterId = NewUnionNullInt()

		return r.DeliveryPointCounterId
	case 1:
		r.Directionn = NewUnionNullDirection()

		return r.Directionn
	case 2:
		r.FromDate = NewUnionNullString()

		return r.FromDate
	case 3:
		r.Id = NewUnionNullString()

		return r.Id
	case 4:
		r.LastReadingDate = NewUnionNullString()

		return r.LastReadingDate
	case 5:
		r.NumberOfDecimals = NewUnionNullInt()

		return r.NumberOfDecimals
	case 6:
		r.NumberOfDigits = NewUnionNullInt()

		return r.NumberOfDigits
	case 7:
		r.PendingApproval = NewUnionNullBool()

		return r.PendingApproval
	case 8:
		r.Product = NewUnionNullProductType()

		return r.Product
	case 9:
		r.ReadingType = NewUnionNullReadingType()

		return r.ReadingType
	case 10:
		r.ToDate = NewUnionNullString()

		return r.ToDate
	case 11:
		r.Unit = NewUnionNullString()

		return r.Unit
	case 12:
		r.YearlyConsumption = NewUnionNullInt()

		return r.YearlyConsumption
	}
	panic("Unknown field index")
}

func (r *Counter) SetDefault(i int) {
	switch i {
	case 0:
		r.DeliveryPointCounterId = nil
		return
	case 1:
		r.Directionn = nil
		return
	case 2:
		r.FromDate = nil
		return
	case 3:
		r.Id = nil
		return
	case 4:
		r.LastReadingDate = nil
		return
	case 5:
		r.NumberOfDecimals = nil
		return
	case 6:
		r.NumberOfDigits = nil
		return
	case 7:
		r.PendingApproval = nil
		return
	case 8:
		r.Product = nil
		return
	case 9:
		r.ReadingType = nil
		return
	case 10:
		r.ToDate = nil
		return
	case 11:
		r.Unit = nil
		return
	case 12:
		r.YearlyConsumption = nil
		return
	}
	panic("Unknown field index")
}

func (r *Counter) NullField(i int) {
	switch i {
	case 0:
		r.DeliveryPointCounterId = nil
		return
	case 1:
		r.Directionn = nil
		return
	case 2:
		r.FromDate = nil
		return
	case 3:
		r.Id = nil
		return
	case 4:
		r.LastReadingDate = nil
		return
	case 5:
		r.NumberOfDecimals = nil
		return
	case 6:
		r.NumberOfDigits = nil
		return
	case 7:
		r.PendingApproval = nil
		return
	case 8:
		r.Product = nil
		return
	case 9:
		r.ReadingType = nil
		return
	case 10:
		r.ToDate = nil
		return
	case 11:
		r.Unit = nil
		return
	case 12:
		r.YearlyConsumption = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Counter) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Counter) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Counter) HintSize(int)                     { panic("Unsupported operation") }
func (_ Counter) Finalize()                        {}

func (_ Counter) AvroCRC64Fingerprint() []byte {
	return []byte(CounterAvroCRC64Fingerprint)
}

func (r Counter) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["DeliveryPointCounterId"], err = json.Marshal(r.DeliveryPointCounterId)
	if err != nil {
		return nil, err
	}
	output["Directionn"], err = json.Marshal(r.Directionn)
	if err != nil {
		return nil, err
	}
	output["FromDate"], err = json.Marshal(r.FromDate)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["LastReadingDate"], err = json.Marshal(r.LastReadingDate)
	if err != nil {
		return nil, err
	}
	output["NumberOfDecimals"], err = json.Marshal(r.NumberOfDecimals)
	if err != nil {
		return nil, err
	}
	output["NumberOfDigits"], err = json.Marshal(r.NumberOfDigits)
	if err != nil {
		return nil, err
	}
	output["PendingApproval"], err = json.Marshal(r.PendingApproval)
	if err != nil {
		return nil, err
	}
	output["Product"], err = json.Marshal(r.Product)
	if err != nil {
		return nil, err
	}
	output["ReadingType"], err = json.Marshal(r.ReadingType)
	if err != nil {
		return nil, err
	}
	output["ToDate"], err = json.Marshal(r.ToDate)
	if err != nil {
		return nil, err
	}
	output["Unit"], err = json.Marshal(r.Unit)
	if err != nil {
		return nil, err
	}
	output["YearlyConsumption"], err = json.Marshal(r.YearlyConsumption)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Counter) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["DeliveryPointCounterId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeliveryPointCounterId); err != nil {
			return err
		}
	} else {
		r.DeliveryPointCounterId = NewUnionNullInt()

		r.DeliveryPointCounterId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Directionn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Directionn); err != nil {
			return err
		}
	} else {
		r.Directionn = NewUnionNullDirection()

		r.Directionn = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FromDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FromDate); err != nil {
			return err
		}
	} else {
		r.FromDate = NewUnionNullString()

		r.FromDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullString()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LastReadingDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LastReadingDate); err != nil {
			return err
		}
	} else {
		r.LastReadingDate = NewUnionNullString()

		r.LastReadingDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumberOfDecimals"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumberOfDecimals); err != nil {
			return err
		}
	} else {
		r.NumberOfDecimals = NewUnionNullInt()

		r.NumberOfDecimals = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumberOfDigits"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumberOfDigits); err != nil {
			return err
		}
	} else {
		r.NumberOfDigits = NewUnionNullInt()

		r.NumberOfDigits = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PendingApproval"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PendingApproval); err != nil {
			return err
		}
	} else {
		r.PendingApproval = NewUnionNullBool()

		r.PendingApproval = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Product"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Product); err != nil {
			return err
		}
	} else {
		r.Product = NewUnionNullProductType()

		r.Product = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReadingType"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReadingType); err != nil {
			return err
		}
	} else {
		r.ReadingType = NewUnionNullReadingType()

		r.ReadingType = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ToDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ToDate); err != nil {
			return err
		}
	} else {
		r.ToDate = NewUnionNullString()

		r.ToDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Unit"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unit); err != nil {
			return err
		}
	} else {
		r.Unit = NewUnionNullString()

		r.Unit = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["YearlyConsumption"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.YearlyConsumption); err != nil {
			return err
		}
	} else {
		r.YearlyConsumption = NewUnionNullInt()

		r.YearlyConsumption = nil
	}
	return nil
}
