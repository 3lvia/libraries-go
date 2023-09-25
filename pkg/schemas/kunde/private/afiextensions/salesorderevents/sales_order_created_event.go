// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100366_4.avsc
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

type SalesOrderCreatedEvent struct {
	ActorId *UnionNullLong `json:"ActorId"`

	CreatedTime *UnionNullLong `json:"CreatedTime"`

	DeliveryPointId *UnionNullInt `json:"DeliveryPointId"`

	EventId string `json:"EventId"`

	EventType string `json:"EventType"`

	ExternalCustomerGuid *UnionNullString `json:"ExternalCustomerGuid"`

	ExternalCustomerReference *UnionNullString `json:"ExternalCustomerReference"`

	Id string `json:"Id"`

	IsOrganization *UnionNullBool `json:"IsOrganization"`

	ObjectId *UnionNullString `json:"ObjectId"`

	SalesOrder SalesOrder `json:"SalesOrder"`

	SalesOrderSerialNumber *UnionNullLong `json:"SalesOrderSerialNumber"`
}

const SalesOrderCreatedEventAvroCRC64Fingerprint = "f\xcb\xc0\x8b\xa0\xf9yS"

func NewSalesOrderCreatedEvent() SalesOrderCreatedEvent {
	r := SalesOrderCreatedEvent{}
	r.ActorId = nil
	r.CreatedTime = nil
	r.DeliveryPointId = nil
	r.ExternalCustomerGuid = nil
	r.ExternalCustomerReference = nil
	r.IsOrganization = nil
	r.ObjectId = nil
	r.SalesOrder = NewSalesOrder()

	r.SalesOrderSerialNumber = nil
	return r
}

func DeserializeSalesOrderCreatedEvent(r io.Reader) (SalesOrderCreatedEvent, error) {
	t := NewSalesOrderCreatedEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSalesOrderCreatedEventFromSchema(r io.Reader, schema string) (SalesOrderCreatedEvent, error) {
	t := NewSalesOrderCreatedEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSalesOrderCreatedEvent(r SalesOrderCreatedEvent, w io.Writer) error {
	var err error
	err = writeUnionNullLong(r.ActorId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.CreatedTime, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.DeliveryPointId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EventId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EventType, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalCustomerGuid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalCustomerReference, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.IsOrganization, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ObjectId, w)
	if err != nil {
		return err
	}
	err = writeSalesOrder(r.SalesOrder, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.SalesOrderSerialNumber, w)
	if err != nil {
		return err
	}
	return err
}

func (r SalesOrderCreatedEvent) Serialize(w io.Writer) error {
	return writeSalesOrderCreatedEvent(r, w)
}

func (r SalesOrderCreatedEvent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"CreatedTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"DeliveryPointId\",\"type\":[\"null\",\"int\"]},{\"name\":\"EventId\",\"type\":\"string\"},{\"name\":\"EventType\",\"type\":\"string\"},{\"default\":null,\"name\":\"ExternalCustomerGuid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExternalCustomerReference\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"default\":null,\"name\":\"IsOrganization\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"ObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"SalesOrder\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Balance\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Amount\",\"type\":{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}},{\"default\":null,\"name\":\"CurrencyCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"Money\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"BilledToDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BillingGroupId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BillingGroupTermId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BillingStatus\",\"type\":[\"null\",{\"name\":\"BillingStatus\",\"symbols\":[\"Open\",\"BlockedForBilling\",\"ManualFinalBilling\",\"ManualFinalBillingAndBlockedForBilling\",\"ManualFinalBillingAndBlockedForBatchBilling\",\"ManualFinalBillingAndBlockedForAllBilling\",\"ManualBlockingBatchBilling\",\"BlockedForAllBilling\",\"BlockedByTakeoverWithoutSupplierChange\",\"BlockedByOutsideThresholdValue\",\"BlockedByOutsideThresholdValueVolumePercent\",\"ApprovedOutsideThresholdValueForBilling\",\"ApprovedOutsideThresholdValueForBillingToManualFinalBilling\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"BillingType\",\"type\":[\"null\",{\"name\":\"BillingType\",\"symbols\":[\"SettlementOrOnAccount\",\"Settlement\",\"EnergySettlement\",\"OnAccountInvoicing\",\"AnnualSettlementAfterEnergySettlement\",\"SettlementOfAdditionalOrders\"],\"type\":\"enum\"}]},{\"name\":\"ClosedForInvoicingChanges\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"CreatedDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Subject\",\"type\":[\"null\",\"string\"]}],\"name\":\"ClosedForInvoicingChange\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"ConsumptionCode\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]}],\"name\":\"ConsumptionCode\",\"type\":\"record\"}]},{\"name\":\"Counters\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"DeliveryPointCounterId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Directionn\",\"type\":[\"null\",{\"name\":\"Direction\",\"symbols\":[\"Production\",\"Consumption\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"FromDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastReadingDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumberOfDecimals\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"NumberOfDigits\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"PendingApproval\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Product\",\"type\":[\"null\",{\"name\":\"ProductType\",\"symbols\":[\"ActiveEnergy\",\"ActivePower\",\"ReactiveEnergy\",\"ReactivePower\",\"DistrictHeatingEnergy\",\"DistrictHeatingPower\",\"DistrictHeatingWaterVolume\",\"DistrictHeatingWaterDelivered\",\"NaturalGasEnergy\",\"NaturalGasVolume\",\"NaturalGasNormalizedVolume\",\"CapacitiveReactiveEnergy\",\"CapacitiveReactivePower\",\"InductiveReactiveEnergy\",\"InductiveReactivePower\",\"ActualProduction\",\"PlannedProduction\",\"ReviewedProductionInElSpot\",\"ActualConsumption\",\"ReviewedConsumptionInElSpot\",\"ConcessionPower\",\"ReplacementPower\",\"IndustryContracts\",\"RKImbalanceAsBalanceResponsibleForGridOwner\",\"EnergyPart\",\"MarginalLossRates\",\"NaturalGasControlVolume\",\"NaturalGasCorrectedVolume\",\"ActiveEnergyPPC\",\"ActiveEnergyFPC\",\"DistrictHeatingSteppedMeasured\",\"DistrictHeatingTemperatureInlet\",\"DistrictHeatingTemperatureOutlet\",\"DistrictHeatingVolBasedAvgTempIn\",\"DistrictHeatingVolBasedAvgTempOut\",\"DistrictCoolingEnergy\",\"DistrictCoolingWaterVolume\",\"DistrictCoolingTemperatureInlet\",\"DistrictCoolingTemperatureOutlet\",\"DistrictCoolingVolBasedAvgTempIn\",\"DistrictCoolingVolBasedAvgTempOut\",\"InvoiceShareElectricity\",\"InvoiceShareCooling\",\"InvoiceShareGasEnergy\",\"InvoiceShareGasVolume\",\"DistrictHeatingVolumeBasedAverageTemperatureInlet\",\"DistrictHeatingVolumeBasedAverageTemperatureOutlet\",\"DistrictCoolingVolumeBasedAverageTemperatureInlet\",\"DistrictCoolingVolumeBasedAverageTemperatureOutlet\",\"ActiveEnergyControl\",\"CustomProduct\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"ReadingType\",\"type\":[\"null\",{\"name\":\"ReadingType\",\"symbols\":[\"Hourly\",\"Manual\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"ToDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unit\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"YearlyConsumption\",\"type\":[\"null\",\"int\"]}],\"name\":\"Counter\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"DeliveryPointId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DeliveryToDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FreeFieldsId\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"FreePowerReflected\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"FromDate\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IndustryCode\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SsbStandard\",\"type\":[\"null\",\"int\"]}],\"name\":\"IndustryCode\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"InvoiceAgreementId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ItemSalesId\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"LocationAddress\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AddressFreeForm\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"City\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CountryCode\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"HouseNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MunicipalityId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropertyUnitNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"StreetName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ZipCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"LocationAddress\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"MeterId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MonitoringStatus\",\"type\":[\"null\",{\"name\":\"MonitoringStatus\",\"symbols\":[\"OnTheWatchList\",\"ApprovedForBilling\",\"NotOnTheWatchList\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"NotificationChannel\",\"type\":[\"null\",{\"name\":\"CommunicationChannelType\",\"symbols\":[\"Casehandling\",\"Chat\",\"ISCustomer\",\"EMail\",\"Facebook\",\"Fax\",\"PersonalAttendance\",\"Letter\",\"Sms\",\"Phone\",\"Web\",\"ISChange\",\"Elhub\",\"Power\",\"Net\",\"DSF\",\"ISSafe\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"NotificationEffectLimit\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"NotificationInterval\",\"type\":[\"null\",{\"name\":\"NotificationInterval\",\"symbols\":[\"Never\",\"Daily\",\"Weekly\",\"Monthly\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"OperatingUnitId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"OperatingUnitSubType\",\"type\":[\"null\",{\"name\":\"SalesOrderOperatingUnitSubType\",\"symbols\":[\"Distribution\",\"Sales\",\"SalesAndDistribution\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"OrderSupplier\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"FromDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"HourlyRead\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"SettlementMethod\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SupplierId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ToDate\",\"type\":[\"null\",\"string\"]}],\"name\":\"OrderSupplier\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"ProductionPercentage\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"ReadyToBeInvoiced\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"SerialNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ShortTermSubscription\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Status\",\"type\":[\"null\",{\"name\":\"SalesOrderStatus\",\"symbols\":[\"All\",\"Active\",\"Anulled\",\"NeedsPriceChange\",\"HandOverInProgress\",\"AboutToBeClosed\",\"Closed\",\"InCloseProcess\",\"AnullmentInitialised\",\"ActivationInitialised\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"SundryExpenceSalesId\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"SupplementaryProductSalesId\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"TariffSalesId\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"TermNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"TermYear\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"ThresholdValueSum\",\"type\":[\"null\",\"Afiextensions.Schemas.SalesOrderEvents.V1.Money\"]},{\"default\":null,\"name\":\"ThresholdValueUpperLimit\",\"type\":[\"null\",\"Afiextensions.Schemas.SalesOrderEvents.V1.Money\"]},{\"default\":null,\"name\":\"ToDate\",\"type\":[\"null\",\"string\"]}],\"name\":\"SalesOrder\",\"type\":\"record\"}},{\"default\":null,\"name\":\"SalesOrderSerialNumber\",\"type\":[\"null\",\"long\"]}],\"name\":\"Afiextensions.Schemas.SalesOrderEvents.V1.SalesOrderCreatedEvent\",\"type\":\"record\"}"
}

func (r SalesOrderCreatedEvent) SchemaName() string {
	return "Afiextensions.Schemas.SalesOrderEvents.V1.SalesOrderCreatedEvent"
}

func (_ SalesOrderCreatedEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SalesOrderCreatedEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.ActorId = NewUnionNullLong()

		return r.ActorId
	case 1:
		r.CreatedTime = NewUnionNullLong()

		return r.CreatedTime
	case 2:
		r.DeliveryPointId = NewUnionNullInt()

		return r.DeliveryPointId
	case 3:
		w := types.String{Target: &r.EventId}

		return w

	case 4:
		w := types.String{Target: &r.EventType}

		return w

	case 5:
		r.ExternalCustomerGuid = NewUnionNullString()

		return r.ExternalCustomerGuid
	case 6:
		r.ExternalCustomerReference = NewUnionNullString()

		return r.ExternalCustomerReference
	case 7:
		w := types.String{Target: &r.Id}

		return w

	case 8:
		r.IsOrganization = NewUnionNullBool()

		return r.IsOrganization
	case 9:
		r.ObjectId = NewUnionNullString()

		return r.ObjectId
	case 10:
		r.SalesOrder = NewSalesOrder()

		w := types.Record{Target: &r.SalesOrder}

		return w

	case 11:
		r.SalesOrderSerialNumber = NewUnionNullLong()

		return r.SalesOrderSerialNumber
	}
	panic("Unknown field index")
}

func (r *SalesOrderCreatedEvent) SetDefault(i int) {
	switch i {
	case 0:
		r.ActorId = nil
		return
	case 1:
		r.CreatedTime = nil
		return
	case 2:
		r.DeliveryPointId = nil
		return
	case 5:
		r.ExternalCustomerGuid = nil
		return
	case 6:
		r.ExternalCustomerReference = nil
		return
	case 8:
		r.IsOrganization = nil
		return
	case 9:
		r.ObjectId = nil
		return
	case 11:
		r.SalesOrderSerialNumber = nil
		return
	}
	panic("Unknown field index")
}

func (r *SalesOrderCreatedEvent) NullField(i int) {
	switch i {
	case 0:
		r.ActorId = nil
		return
	case 1:
		r.CreatedTime = nil
		return
	case 2:
		r.DeliveryPointId = nil
		return
	case 5:
		r.ExternalCustomerGuid = nil
		return
	case 6:
		r.ExternalCustomerReference = nil
		return
	case 8:
		r.IsOrganization = nil
		return
	case 9:
		r.ObjectId = nil
		return
	case 11:
		r.SalesOrderSerialNumber = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ SalesOrderCreatedEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ SalesOrderCreatedEvent) Finalize()                        {}

func (_ SalesOrderCreatedEvent) AvroCRC64Fingerprint() []byte {
	return []byte(SalesOrderCreatedEventAvroCRC64Fingerprint)
}

func (r SalesOrderCreatedEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ActorId"], err = json.Marshal(r.ActorId)
	if err != nil {
		return nil, err
	}
	output["CreatedTime"], err = json.Marshal(r.CreatedTime)
	if err != nil {
		return nil, err
	}
	output["DeliveryPointId"], err = json.Marshal(r.DeliveryPointId)
	if err != nil {
		return nil, err
	}
	output["EventId"], err = json.Marshal(r.EventId)
	if err != nil {
		return nil, err
	}
	output["EventType"], err = json.Marshal(r.EventType)
	if err != nil {
		return nil, err
	}
	output["ExternalCustomerGuid"], err = json.Marshal(r.ExternalCustomerGuid)
	if err != nil {
		return nil, err
	}
	output["ExternalCustomerReference"], err = json.Marshal(r.ExternalCustomerReference)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["IsOrganization"], err = json.Marshal(r.IsOrganization)
	if err != nil {
		return nil, err
	}
	output["ObjectId"], err = json.Marshal(r.ObjectId)
	if err != nil {
		return nil, err
	}
	output["SalesOrder"], err = json.Marshal(r.SalesOrder)
	if err != nil {
		return nil, err
	}
	output["SalesOrderSerialNumber"], err = json.Marshal(r.SalesOrderSerialNumber)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SalesOrderCreatedEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ActorId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ActorId); err != nil {
			return err
		}
	} else {
		r.ActorId = NewUnionNullLong()

		r.ActorId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CreatedTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreatedTime); err != nil {
			return err
		}
	} else {
		r.CreatedTime = NewUnionNullLong()

		r.CreatedTime = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DeliveryPointId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DeliveryPointId); err != nil {
			return err
		}
	} else {
		r.DeliveryPointId = NewUnionNullInt()

		r.DeliveryPointId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EventId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EventType"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventType); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EventType")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ExternalCustomerGuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalCustomerGuid); err != nil {
			return err
		}
	} else {
		r.ExternalCustomerGuid = NewUnionNullString()

		r.ExternalCustomerGuid = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ExternalCustomerReference"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalCustomerReference); err != nil {
			return err
		}
	} else {
		r.ExternalCustomerReference = NewUnionNullString()

		r.ExternalCustomerReference = nil
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
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IsOrganization"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsOrganization); err != nil {
			return err
		}
	} else {
		r.IsOrganization = NewUnionNullBool()

		r.IsOrganization = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ObjectId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ObjectId); err != nil {
			return err
		}
	} else {
		r.ObjectId = NewUnionNullString()

		r.ObjectId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SalesOrder"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SalesOrder); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SalesOrder")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SalesOrderSerialNumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SalesOrderSerialNumber); err != nil {
			return err
		}
	} else {
		r.SalesOrderSerialNumber = NewUnionNullLong()

		r.SalesOrderSerialNumber = nil
	}
	return nil
}
