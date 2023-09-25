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

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ExternalOrigin struct {
	Reference *UnionNullString `json:"Reference"`

	System *UnionNullExternalSystem `json:"System"`

	UpdatedAt *UnionNullString `json:"UpdatedAt"`

	UpdatedBy *UnionNullUser `json:"UpdatedBy"`
}

const ExternalOriginAvroCRC64Fingerprint = "E!}e[R3t"

func NewExternalOrigin() ExternalOrigin {
	r := ExternalOrigin{}
	r.Reference = nil
	r.System = nil
	r.UpdatedAt = nil
	r.UpdatedBy = nil
	return r
}

func DeserializeExternalOrigin(r io.Reader) (ExternalOrigin, error) {
	t := NewExternalOrigin()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeExternalOriginFromSchema(r io.Reader, schema string) (ExternalOrigin, error) {
	t := NewExternalOrigin()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeExternalOrigin(r ExternalOrigin, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Reference, w)
	if err != nil {
		return err
	}
	err = writeUnionNullExternalSystem(r.System, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UpdatedAt, w)
	if err != nil {
		return err
	}
	err = writeUnionNullUser(r.UpdatedBy, w)
	if err != nil {
		return err
	}
	return err
}

func (r ExternalOrigin) Serialize(w io.Writer) error {
	return writeExternalOrigin(r, w)
}

func (r ExternalOrigin) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Reference\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"System\",\"type\":[\"null\",{\"name\":\"ExternalSystem\",\"symbols\":[\"Unknown\",\"Ifs\",\"Elwin\",\"Geonis\",\"Quant\",\"Cab\",\"Mmm\",\"Mdmx\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"UpdatedAt\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UpdatedBy\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]}],\"name\":\"User\",\"namespace\":\"Msim.Domain.Model.Primitives\",\"type\":\"record\"}]}],\"name\":\"Msim.Domain.Model.Events.ExternalOrigin\",\"type\":\"record\"}"
}

func (r ExternalOrigin) SchemaName() string {
	return "Msim.Domain.Model.Events.ExternalOrigin"
}

func (_ ExternalOrigin) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ExternalOrigin) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ExternalOrigin) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ExternalOrigin) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ExternalOrigin) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ExternalOrigin) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ExternalOrigin) SetString(v string)   { panic("Unsupported operation") }
func (_ ExternalOrigin) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ExternalOrigin) Get(i int) types.Field {
	switch i {
	case 0:
		r.Reference = NewUnionNullString()

		return r.Reference
	case 1:
		r.System = NewUnionNullExternalSystem()

		return r.System
	case 2:
		r.UpdatedAt = NewUnionNullString()

		return r.UpdatedAt
	case 3:
		r.UpdatedBy = NewUnionNullUser()

		return r.UpdatedBy
	}
	panic("Unknown field index")
}

func (r *ExternalOrigin) SetDefault(i int) {
	switch i {
	case 0:
		r.Reference = nil
		return
	case 1:
		r.System = nil
		return
	case 2:
		r.UpdatedAt = nil
		return
	case 3:
		r.UpdatedBy = nil
		return
	}
	panic("Unknown field index")
}

func (r *ExternalOrigin) NullField(i int) {
	switch i {
	case 0:
		r.Reference = nil
		return
	case 1:
		r.System = nil
		return
	case 2:
		r.UpdatedAt = nil
		return
	case 3:
		r.UpdatedBy = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ExternalOrigin) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ExternalOrigin) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ExternalOrigin) HintSize(int)                     { panic("Unsupported operation") }
func (_ ExternalOrigin) Finalize()                        {}

func (_ ExternalOrigin) AvroCRC64Fingerprint() []byte {
	return []byte(ExternalOriginAvroCRC64Fingerprint)
}

func (r ExternalOrigin) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Reference"], err = json.Marshal(r.Reference)
	if err != nil {
		return nil, err
	}
	output["System"], err = json.Marshal(r.System)
	if err != nil {
		return nil, err
	}
	output["UpdatedAt"], err = json.Marshal(r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	output["UpdatedBy"], err = json.Marshal(r.UpdatedBy)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ExternalOrigin) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Reference"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Reference); err != nil {
			return err
		}
	} else {
		r.Reference = NewUnionNullString()

		r.Reference = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["System"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.System); err != nil {
			return err
		}
	} else {
		r.System = NewUnionNullExternalSystem()

		r.System = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["UpdatedAt"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdatedAt); err != nil {
			return err
		}
	} else {
		r.UpdatedAt = NewUnionNullString()

		r.UpdatedAt = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["UpdatedBy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdatedBy); err != nil {
			return err
		}
	} else {
		r.UpdatedBy = NewUnionNullUser()

		r.UpdatedBy = nil
	}
	return nil
}
