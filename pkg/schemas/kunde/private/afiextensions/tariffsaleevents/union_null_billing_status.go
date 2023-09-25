// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100346_2.avsc
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

type UnionNullBillingStatusTypeEnum int

const (
	UnionNullBillingStatusTypeEnumBillingStatus UnionNullBillingStatusTypeEnum = 1
)

type UnionNullBillingStatus struct {
	Null          *types.NullVal
	BillingStatus BillingStatus
	UnionType     UnionNullBillingStatusTypeEnum
}

func writeUnionNullBillingStatus(r *UnionNullBillingStatus, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBillingStatusTypeEnumBillingStatus:
		return writeBillingStatus(r.BillingStatus, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBillingStatus")
}

func NewUnionNullBillingStatus() *UnionNullBillingStatus {
	return &UnionNullBillingStatus{}
}

func (r *UnionNullBillingStatus) Serialize(w io.Writer) error {
	return writeUnionNullBillingStatus(r, w)
}

func DeserializeUnionNullBillingStatus(r io.Reader) (*UnionNullBillingStatus, error) {
	t := NewUnionNullBillingStatus()
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

func DeserializeUnionNullBillingStatusFromSchema(r io.Reader, schema string) (*UnionNullBillingStatus, error) {
	t := NewUnionNullBillingStatus()
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

func (r *UnionNullBillingStatus) Schema() string {
	return "[\"null\",{\"name\":\"BillingStatus\",\"symbols\":[\"Open\",\"BlockedForBilling\",\"ManualFinalBilling\",\"ManualFinalBillingAndBlockedForBilling\",\"ManualFinalBillingAndBlockedForBatchBilling\",\"ManualFinalBillingAndBlockedForAllBilling\",\"ManualBlockingBatchBilling\",\"BlockedForAllBilling\",\"BlockedByTakeoverWithoutSupplierChange\",\"BlockedByOutsideThresholdValue\",\"BlockedByOutsideThresholdValueVolumePercent\",\"ApprovedOutsideThresholdValueForBilling\",\"ApprovedOutsideThresholdValueForBillingToManualFinalBilling\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullBillingStatus) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBillingStatus) SetLong(v int64) {

	r.UnionType = (UnionNullBillingStatusTypeEnum)(v)
}

func (r *UnionNullBillingStatus) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &BillingStatusWrapper{Target: (&r.BillingStatus)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBillingStatus) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBillingStatus) Finalize()                        {}

func (r *UnionNullBillingStatus) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBillingStatusTypeEnumBillingStatus:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.TariffEvents.V1.BillingStatus": r.BillingStatus})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBillingStatus")
}

func (r *UnionNullBillingStatus) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.TariffEvents.V1.BillingStatus"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BillingStatus)
	}
	return fmt.Errorf("invalid value for *UnionNullBillingStatus")
}
