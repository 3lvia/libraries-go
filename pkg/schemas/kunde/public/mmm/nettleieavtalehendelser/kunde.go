// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100480_13.avsc
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

type Kunde struct {
	Etternavn *UnionNullString `json:"Etternavn"`

	Fornavn *UnionNullString `json:"Fornavn"`

	KundeGuid string `json:"KundeGuid"`

	Virksomhetsnavn *UnionNullString `json:"Virksomhetsnavn"`
}

const KundeAvroCRC64Fingerprint = "OS&\xf5\x8fJ\xb1@"

func NewKunde() Kunde {
	r := Kunde{}
	r.Etternavn = nil
	r.Fornavn = nil
	r.Virksomhetsnavn = nil
	return r
}

func DeserializeKunde(r io.Reader) (Kunde, error) {
	t := NewKunde()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeKundeFromSchema(r io.Reader, schema string) (Kunde, error) {
	t := NewKunde()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeKunde(r Kunde, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Etternavn, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Fornavn, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.KundeGuid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Virksomhetsnavn, w)
	if err != nil {
		return err
	}
	return err
}

func (r Kunde) Serialize(w io.Writer) error {
	return writeKunde(r, w)
}

func (r Kunde) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Etternavn\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Fornavn\",\"type\":[\"null\",\"string\"]},{\"name\":\"KundeGuid\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"default\":null,\"name\":\"Virksomhetsnavn\",\"type\":[\"null\",\"string\"]}],\"name\":\"Mmm.Schemas.NettleieAvtaleHendelser.V1.Kunde\",\"type\":\"record\"}"
}

func (r Kunde) SchemaName() string {
	return "Mmm.Schemas.NettleieAvtaleHendelser.V1.Kunde"
}

func (_ Kunde) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Kunde) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Kunde) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Kunde) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Kunde) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Kunde) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Kunde) SetString(v string)   { panic("Unsupported operation") }
func (_ Kunde) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Kunde) Get(i int) types.Field {
	switch i {
	case 0:
		r.Etternavn = NewUnionNullString()

		return r.Etternavn
	case 1:
		r.Fornavn = NewUnionNullString()

		return r.Fornavn
	case 2:
		w := types.String{Target: &r.KundeGuid}

		return w

	case 3:
		r.Virksomhetsnavn = NewUnionNullString()

		return r.Virksomhetsnavn
	}
	panic("Unknown field index")
}

func (r *Kunde) SetDefault(i int) {
	switch i {
	case 0:
		r.Etternavn = nil
		return
	case 1:
		r.Fornavn = nil
		return
	case 3:
		r.Virksomhetsnavn = nil
		return
	}
	panic("Unknown field index")
}

func (r *Kunde) NullField(i int) {
	switch i {
	case 0:
		r.Etternavn = nil
		return
	case 1:
		r.Fornavn = nil
		return
	case 3:
		r.Virksomhetsnavn = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Kunde) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Kunde) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Kunde) HintSize(int)                     { panic("Unsupported operation") }
func (_ Kunde) Finalize()                        {}

func (_ Kunde) AvroCRC64Fingerprint() []byte {
	return []byte(KundeAvroCRC64Fingerprint)
}

func (r Kunde) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Etternavn"], err = json.Marshal(r.Etternavn)
	if err != nil {
		return nil, err
	}
	output["Fornavn"], err = json.Marshal(r.Fornavn)
	if err != nil {
		return nil, err
	}
	output["KundeGuid"], err = json.Marshal(r.KundeGuid)
	if err != nil {
		return nil, err
	}
	output["Virksomhetsnavn"], err = json.Marshal(r.Virksomhetsnavn)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Kunde) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Etternavn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Etternavn); err != nil {
			return err
		}
	} else {
		r.Etternavn = NewUnionNullString()

		r.Etternavn = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Fornavn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fornavn); err != nil {
			return err
		}
	} else {
		r.Fornavn = NewUnionNullString()

		r.Fornavn = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["KundeGuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.KundeGuid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for KundeGuid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Virksomhetsnavn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Virksomhetsnavn); err != nil {
			return err
		}
	} else {
		r.Virksomhetsnavn = NewUnionNullString()

		r.Virksomhetsnavn = nil
	}
	return nil
}
