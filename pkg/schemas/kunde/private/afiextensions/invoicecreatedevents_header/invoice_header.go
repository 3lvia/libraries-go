// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100397_2.avsc
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

type InvoiceHeader struct {
	ActorId *UnionNullString `json:"ActorId"`

	Amount *UnionNullMoney `json:"Amount"`

	Balance *UnionNullMoney `json:"Balance"`

	Cid string `json:"Cid"`

	CreditationAllowed *UnionNullString `json:"CreditationAllowed"`

	DateCredited *UnionNullString `json:"DateCredited"`

	DueDate string `json:"DueDate"`

	GiroAmount *UnionNullMoney `json:"GiroAmount"`

	GrossPresentation *UnionNullString `json:"GrossPresentation"`

	Id string `json:"Id"`

	InvoiceDate string `json:"InvoiceDate"`

	InvoiceNumber *UnionNullInt `json:"InvoiceNumber"`

	MeterIds *UnionNullArrayString `json:"MeterIds"`

	MeterPointIds *UnionNullArrayString `json:"MeterPointIds"`

	PrintedDate *UnionNullString `json:"PrintedDate"`

	ProductionDate string `json:"ProductionDate"`

	StatusEnum *UnionNullInvoiceStatusEnum `json:"StatusEnum"`

	SumPayableInterestDue *UnionNullLong `json:"SumPayableInterestDue"`

	Term *UnionNullLong `json:"Term"`

	Type *UnionNullString `json:"Type"`

	Year *UnionNullString `json:"Year"`
}

const InvoiceHeaderAvroCRC64Fingerprint = "\x15>2\f9U\xfa\t"

func NewInvoiceHeader() InvoiceHeader {
	r := InvoiceHeader{}
	r.ActorId = nil
	r.Amount = nil
	r.Balance = nil
	r.CreditationAllowed = nil
	r.DateCredited = nil
	r.GiroAmount = nil
	r.GrossPresentation = nil
	r.InvoiceNumber = nil
	r.MeterIds = nil
	r.MeterPointIds = nil
	r.PrintedDate = nil
	r.StatusEnum = nil
	r.SumPayableInterestDue = nil
	r.Term = nil
	r.Type = nil
	r.Year = nil
	return r
}

func DeserializeInvoiceHeader(r io.Reader) (InvoiceHeader, error) {
	t := NewInvoiceHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInvoiceHeaderFromSchema(r io.Reader, schema string) (InvoiceHeader, error) {
	t := NewInvoiceHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInvoiceHeader(r InvoiceHeader, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.ActorId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMoney(r.Amount, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMoney(r.Balance, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CreditationAllowed, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DateCredited, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DueDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMoney(r.GiroAmount, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GrossPresentation, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.InvoiceDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.InvoiceNumber, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.MeterIds, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.MeterPointIds, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PrintedDate, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ProductionDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInvoiceStatusEnum(r.StatusEnum, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.SumPayableInterestDue, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Term, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Type, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Year, w)
	if err != nil {
		return err
	}
	return err
}

func (r InvoiceHeader) Serialize(w io.Writer) error {
	return writeInvoiceHeader(r, w)
}

func (r InvoiceHeader) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"ActorId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Amount\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Amount\",\"type\":{\"logicalType\":\"decimal\",\"precision\":29,\"scale\":14,\"type\":\"bytes\"}},{\"name\":\"CurrencyCode\",\"type\":\"string\"}],\"name\":\"Money\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Balance\",\"type\":[\"null\",\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.Money\"]},{\"name\":\"Cid\",\"type\":\"string\"},{\"default\":null,\"name\":\"CreditationAllowed\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DateCredited\",\"type\":[\"null\",\"string\"]},{\"name\":\"DueDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"GiroAmount\",\"type\":[\"null\",\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.Money\"]},{\"default\":null,\"name\":\"GrossPresentation\",\"type\":[\"null\",\"string\"]},{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"InvoiceDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"InvoiceNumber\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"MeterIds\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"MeterPointIds\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"PrintedDate\",\"type\":[\"null\",\"string\"]},{\"name\":\"ProductionDate\",\"type\":\"string\"},{\"default\":null,\"name\":\"StatusEnum\",\"type\":[\"null\",{\"name\":\"InvoiceStatusEnum\",\"symbols\":[\"Finished\",\"Unfinished\",\"WrittenOff\",\"NotApplicable\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"SumPayableInterestDue\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Term\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Type\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Year\",\"type\":[\"null\",\"string\"]}],\"name\":\"Afiextensions.Schemas.InvoiceCreatedEvents.V1.InvoiceHeader\",\"type\":\"record\"}"
}

func (r InvoiceHeader) SchemaName() string {
	return "Afiextensions.Schemas.InvoiceCreatedEvents.V1.InvoiceHeader"
}

func (_ InvoiceHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InvoiceHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InvoiceHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InvoiceHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InvoiceHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InvoiceHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InvoiceHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ InvoiceHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InvoiceHeader) Get(i int) types.Field {
	switch i {
	case 0:
		r.ActorId = NewUnionNullString()

		return r.ActorId
	case 1:
		r.Amount = NewUnionNullMoney()

		return r.Amount
	case 2:
		r.Balance = NewUnionNullMoney()

		return r.Balance
	case 3:
		w := types.String{Target: &r.Cid}

		return w

	case 4:
		r.CreditationAllowed = NewUnionNullString()

		return r.CreditationAllowed
	case 5:
		r.DateCredited = NewUnionNullString()

		return r.DateCredited
	case 6:
		w := types.String{Target: &r.DueDate}

		return w

	case 7:
		r.GiroAmount = NewUnionNullMoney()

		return r.GiroAmount
	case 8:
		r.GrossPresentation = NewUnionNullString()

		return r.GrossPresentation
	case 9:
		w := types.String{Target: &r.Id}

		return w

	case 10:
		w := types.String{Target: &r.InvoiceDate}

		return w

	case 11:
		r.InvoiceNumber = NewUnionNullInt()

		return r.InvoiceNumber
	case 12:
		r.MeterIds = NewUnionNullArrayString()

		return r.MeterIds
	case 13:
		r.MeterPointIds = NewUnionNullArrayString()

		return r.MeterPointIds
	case 14:
		r.PrintedDate = NewUnionNullString()

		return r.PrintedDate
	case 15:
		w := types.String{Target: &r.ProductionDate}

		return w

	case 16:
		r.StatusEnum = NewUnionNullInvoiceStatusEnum()

		return r.StatusEnum
	case 17:
		r.SumPayableInterestDue = NewUnionNullLong()

		return r.SumPayableInterestDue
	case 18:
		r.Term = NewUnionNullLong()

		return r.Term
	case 19:
		r.Type = NewUnionNullString()

		return r.Type
	case 20:
		r.Year = NewUnionNullString()

		return r.Year
	}
	panic("Unknown field index")
}

func (r *InvoiceHeader) SetDefault(i int) {
	switch i {
	case 0:
		r.ActorId = nil
		return
	case 1:
		r.Amount = nil
		return
	case 2:
		r.Balance = nil
		return
	case 4:
		r.CreditationAllowed = nil
		return
	case 5:
		r.DateCredited = nil
		return
	case 7:
		r.GiroAmount = nil
		return
	case 8:
		r.GrossPresentation = nil
		return
	case 11:
		r.InvoiceNumber = nil
		return
	case 12:
		r.MeterIds = nil
		return
	case 13:
		r.MeterPointIds = nil
		return
	case 14:
		r.PrintedDate = nil
		return
	case 16:
		r.StatusEnum = nil
		return
	case 17:
		r.SumPayableInterestDue = nil
		return
	case 18:
		r.Term = nil
		return
	case 19:
		r.Type = nil
		return
	case 20:
		r.Year = nil
		return
	}
	panic("Unknown field index")
}

func (r *InvoiceHeader) NullField(i int) {
	switch i {
	case 0:
		r.ActorId = nil
		return
	case 1:
		r.Amount = nil
		return
	case 2:
		r.Balance = nil
		return
	case 4:
		r.CreditationAllowed = nil
		return
	case 5:
		r.DateCredited = nil
		return
	case 7:
		r.GiroAmount = nil
		return
	case 8:
		r.GrossPresentation = nil
		return
	case 11:
		r.InvoiceNumber = nil
		return
	case 12:
		r.MeterIds = nil
		return
	case 13:
		r.MeterPointIds = nil
		return
	case 14:
		r.PrintedDate = nil
		return
	case 16:
		r.StatusEnum = nil
		return
	case 17:
		r.SumPayableInterestDue = nil
		return
	case 18:
		r.Term = nil
		return
	case 19:
		r.Type = nil
		return
	case 20:
		r.Year = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ InvoiceHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InvoiceHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InvoiceHeader) HintSize(int)                     { panic("Unsupported operation") }
func (_ InvoiceHeader) Finalize()                        {}

func (_ InvoiceHeader) AvroCRC64Fingerprint() []byte {
	return []byte(InvoiceHeaderAvroCRC64Fingerprint)
}

func (r InvoiceHeader) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ActorId"], err = json.Marshal(r.ActorId)
	if err != nil {
		return nil, err
	}
	output["Amount"], err = json.Marshal(r.Amount)
	if err != nil {
		return nil, err
	}
	output["Balance"], err = json.Marshal(r.Balance)
	if err != nil {
		return nil, err
	}
	output["Cid"], err = json.Marshal(r.Cid)
	if err != nil {
		return nil, err
	}
	output["CreditationAllowed"], err = json.Marshal(r.CreditationAllowed)
	if err != nil {
		return nil, err
	}
	output["DateCredited"], err = json.Marshal(r.DateCredited)
	if err != nil {
		return nil, err
	}
	output["DueDate"], err = json.Marshal(r.DueDate)
	if err != nil {
		return nil, err
	}
	output["GiroAmount"], err = json.Marshal(r.GiroAmount)
	if err != nil {
		return nil, err
	}
	output["GrossPresentation"], err = json.Marshal(r.GrossPresentation)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["InvoiceDate"], err = json.Marshal(r.InvoiceDate)
	if err != nil {
		return nil, err
	}
	output["InvoiceNumber"], err = json.Marshal(r.InvoiceNumber)
	if err != nil {
		return nil, err
	}
	output["MeterIds"], err = json.Marshal(r.MeterIds)
	if err != nil {
		return nil, err
	}
	output["MeterPointIds"], err = json.Marshal(r.MeterPointIds)
	if err != nil {
		return nil, err
	}
	output["PrintedDate"], err = json.Marshal(r.PrintedDate)
	if err != nil {
		return nil, err
	}
	output["ProductionDate"], err = json.Marshal(r.ProductionDate)
	if err != nil {
		return nil, err
	}
	output["StatusEnum"], err = json.Marshal(r.StatusEnum)
	if err != nil {
		return nil, err
	}
	output["SumPayableInterestDue"], err = json.Marshal(r.SumPayableInterestDue)
	if err != nil {
		return nil, err
	}
	output["Term"], err = json.Marshal(r.Term)
	if err != nil {
		return nil, err
	}
	output["Type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["Year"], err = json.Marshal(r.Year)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InvoiceHeader) UnmarshalJSON(data []byte) error {
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
		r.ActorId = NewUnionNullString()

		r.ActorId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Amount"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Amount); err != nil {
			return err
		}
	} else {
		r.Amount = NewUnionNullMoney()

		r.Amount = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Balance"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Balance); err != nil {
			return err
		}
	} else {
		r.Balance = NewUnionNullMoney()

		r.Balance = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CreditationAllowed"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreditationAllowed); err != nil {
			return err
		}
	} else {
		r.CreditationAllowed = NewUnionNullString()

		r.CreditationAllowed = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DateCredited"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DateCredited); err != nil {
			return err
		}
	} else {
		r.DateCredited = NewUnionNullString()

		r.DateCredited = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DueDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DueDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DueDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["GiroAmount"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GiroAmount); err != nil {
			return err
		}
	} else {
		r.GiroAmount = NewUnionNullMoney()

		r.GiroAmount = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["GrossPresentation"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GrossPresentation); err != nil {
			return err
		}
	} else {
		r.GrossPresentation = NewUnionNullString()

		r.GrossPresentation = nil
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
		if v, ok := fields["InvoiceDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for InvoiceDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["InvoiceNumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InvoiceNumber); err != nil {
			return err
		}
	} else {
		r.InvoiceNumber = NewUnionNullInt()

		r.InvoiceNumber = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["MeterIds"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MeterIds); err != nil {
			return err
		}
	} else {
		r.MeterIds = NewUnionNullArrayString()

		r.MeterIds = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["MeterPointIds"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MeterPointIds); err != nil {
			return err
		}
	} else {
		r.MeterPointIds = NewUnionNullArrayString()

		r.MeterPointIds = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PrintedDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PrintedDate); err != nil {
			return err
		}
	} else {
		r.PrintedDate = NewUnionNullString()

		r.PrintedDate = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProductionDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductionDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ProductionDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StatusEnum"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StatusEnum); err != nil {
			return err
		}
	} else {
		r.StatusEnum = NewUnionNullInvoiceStatusEnum()

		r.StatusEnum = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SumPayableInterestDue"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SumPayableInterestDue); err != nil {
			return err
		}
	} else {
		r.SumPayableInterestDue = NewUnionNullLong()

		r.SumPayableInterestDue = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Term"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Term); err != nil {
			return err
		}
	} else {
		r.Term = NewUnionNullLong()

		r.Term = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		r.Type = NewUnionNullString()

		r.Type = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Year"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Year); err != nil {
			return err
		}
	} else {
		r.Year = NewUnionNullString()

		r.Year = nil
	}
	return nil
}
