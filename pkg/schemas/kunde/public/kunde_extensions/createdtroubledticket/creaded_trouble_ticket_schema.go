// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100467_3.avsc
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

type CreadedTroubleTicketSchema struct {
	Customer *UnionNullCustomer `json:"Customer"`

	DateTimeOfReport *UnionNullLong `json:"DateTimeOfReport"`

	Description *UnionNullString `json:"Description"`

	MRID *UnionNullString `json:"MRID"`

	TroubleCode *UnionNullString `json:"TroubleCode"`
}

const CreadedTroubleTicketSchemaAvroCRC64Fingerprint = "Ax\xa1Gs\xf2X?"

func NewCreadedTroubleTicketSchema() CreadedTroubleTicketSchema {
	r := CreadedTroubleTicketSchema{}
	r.Customer = nil
	r.DateTimeOfReport = nil
	r.Description = nil
	r.MRID = nil
	r.TroubleCode = nil
	return r
}

func DeserializeCreadedTroubleTicketSchema(r io.Reader) (CreadedTroubleTicketSchema, error) {
	t := NewCreadedTroubleTicketSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCreadedTroubleTicketSchemaFromSchema(r io.Reader, schema string) (CreadedTroubleTicketSchema, error) {
	t := NewCreadedTroubleTicketSchema()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCreadedTroubleTicketSchema(r CreadedTroubleTicketSchema, w io.Writer) error {
	var err error
	err = writeUnionNullCustomer(r.Customer, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.DateTimeOfReport, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Description, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MRID, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TroubleCode, w)
	if err != nil {
		return err
	}
	return err
}

func (r CreadedTroubleTicketSchema) Serialize(w io.Writer) error {
	return writeCreadedTroubleTicketSchema(r, w)
}

func (r CreadedTroubleTicketSchema) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Customer\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ContactType\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ContactValue\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InstallationOid\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Trigger\",\"type\":[\"null\",{\"name\":\"Trigger\",\"symbols\":[\"EtrChange\",\"InformDispatched\",\"InitialEtr\",\"PowerOut\",\"PowerRestored\"],\"type\":\"enum\"}]}],\"name\":\"Customer\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"DateTimeOfReport\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TroubleCode\",\"type\":[\"null\",\"string\"]}],\"name\":\"kunde_extensions_wop_api.Model.CreadedTroubleTicketSchema\",\"type\":\"record\"}"
}

func (r CreadedTroubleTicketSchema) SchemaName() string {
	return "kunde_extensions_wop_api.Model.CreadedTroubleTicketSchema"
}

func (_ CreadedTroubleTicketSchema) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetString(v string)   { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CreadedTroubleTicketSchema) Get(i int) types.Field {
	switch i {
	case 0:
		r.Customer = NewUnionNullCustomer()

		return r.Customer
	case 1:
		r.DateTimeOfReport = NewUnionNullLong()

		return r.DateTimeOfReport
	case 2:
		r.Description = NewUnionNullString()

		return r.Description
	case 3:
		r.MRID = NewUnionNullString()

		return r.MRID
	case 4:
		r.TroubleCode = NewUnionNullString()

		return r.TroubleCode
	}
	panic("Unknown field index")
}

func (r *CreadedTroubleTicketSchema) SetDefault(i int) {
	switch i {
	case 0:
		r.Customer = nil
		return
	case 1:
		r.DateTimeOfReport = nil
		return
	case 2:
		r.Description = nil
		return
	case 3:
		r.MRID = nil
		return
	case 4:
		r.TroubleCode = nil
		return
	}
	panic("Unknown field index")
}

func (r *CreadedTroubleTicketSchema) NullField(i int) {
	switch i {
	case 0:
		r.Customer = nil
		return
	case 1:
		r.DateTimeOfReport = nil
		return
	case 2:
		r.Description = nil
		return
	case 3:
		r.MRID = nil
		return
	case 4:
		r.TroubleCode = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CreadedTroubleTicketSchema) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) HintSize(int)                     { panic("Unsupported operation") }
func (_ CreadedTroubleTicketSchema) Finalize()                        {}

func (_ CreadedTroubleTicketSchema) AvroCRC64Fingerprint() []byte {
	return []byte(CreadedTroubleTicketSchemaAvroCRC64Fingerprint)
}

func (r CreadedTroubleTicketSchema) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Customer"], err = json.Marshal(r.Customer)
	if err != nil {
		return nil, err
	}
	output["DateTimeOfReport"], err = json.Marshal(r.DateTimeOfReport)
	if err != nil {
		return nil, err
	}
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["MRID"], err = json.Marshal(r.MRID)
	if err != nil {
		return nil, err
	}
	output["TroubleCode"], err = json.Marshal(r.TroubleCode)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CreadedTroubleTicketSchema) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Customer"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Customer); err != nil {
			return err
		}
	} else {
		r.Customer = NewUnionNullCustomer()

		r.Customer = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DateTimeOfReport"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DateTimeOfReport); err != nil {
			return err
		}
	} else {
		r.DateTimeOfReport = NewUnionNullLong()

		r.DateTimeOfReport = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		r.Description = NewUnionNullString()

		r.Description = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["MRID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MRID); err != nil {
			return err
		}
	} else {
		r.MRID = NewUnionNullString()

		r.MRID = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TroubleCode"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TroubleCode); err != nil {
			return err
		}
	} else {
		r.TroubleCode = NewUnionNullString()

		r.TroubleCode = nil
	}
	return nil
}
