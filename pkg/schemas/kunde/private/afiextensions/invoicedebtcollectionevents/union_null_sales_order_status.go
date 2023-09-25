// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100395_1.avsc
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

type UnionNullSalesOrderStatusTypeEnum int

const (
	UnionNullSalesOrderStatusTypeEnumSalesOrderStatus UnionNullSalesOrderStatusTypeEnum = 1
)

type UnionNullSalesOrderStatus struct {
	Null             *types.NullVal
	SalesOrderStatus SalesOrderStatus
	UnionType        UnionNullSalesOrderStatusTypeEnum
}

func writeUnionNullSalesOrderStatus(r *UnionNullSalesOrderStatus, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullSalesOrderStatusTypeEnumSalesOrderStatus:
		return writeSalesOrderStatus(r.SalesOrderStatus, w)
	}
	return fmt.Errorf("invalid value for *UnionNullSalesOrderStatus")
}

func NewUnionNullSalesOrderStatus() *UnionNullSalesOrderStatus {
	return &UnionNullSalesOrderStatus{}
}

func (r *UnionNullSalesOrderStatus) Serialize(w io.Writer) error {
	return writeUnionNullSalesOrderStatus(r, w)
}

func DeserializeUnionNullSalesOrderStatus(r io.Reader) (*UnionNullSalesOrderStatus, error) {
	t := NewUnionNullSalesOrderStatus()
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

func DeserializeUnionNullSalesOrderStatusFromSchema(r io.Reader, schema string) (*UnionNullSalesOrderStatus, error) {
	t := NewUnionNullSalesOrderStatus()
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

func (r *UnionNullSalesOrderStatus) Schema() string {
	return "[\"null\",{\"name\":\"SalesOrderStatus\",\"symbols\":[\"All\",\"Active\",\"Anulled\",\"NeedsPriceChange\",\"HandOverInProgress\",\"AboutToBeClosed\",\"Closed\",\"InCloseProcess\",\"AnullmentInitialised\",\"ActivationInitialised\"],\"type\":\"enum\"}]"
}

func (_ *UnionNullSalesOrderStatus) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullSalesOrderStatus) SetLong(v int64) {

	r.UnionType = (UnionNullSalesOrderStatusTypeEnum)(v)
}

func (r *UnionNullSalesOrderStatus) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		return &SalesOrderStatusWrapper{Target: (&r.SalesOrderStatus)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullSalesOrderStatus) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullSalesOrderStatus) Finalize()                        {}

func (r *UnionNullSalesOrderStatus) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullSalesOrderStatusTypeEnumSalesOrderStatus:
		return json.Marshal(map[string]interface{}{"Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.SalesOrderStatus": r.SalesOrderStatus})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullSalesOrderStatus")
}

func (r *UnionNullSalesOrderStatus) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Afiextensions.Schemas.InvoiceDebtCollectionEvents.V1.SalesOrderStatus"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.SalesOrderStatus)
	}
	return fmt.Errorf("invalid value for *UnionNullSalesOrderStatus")
}
