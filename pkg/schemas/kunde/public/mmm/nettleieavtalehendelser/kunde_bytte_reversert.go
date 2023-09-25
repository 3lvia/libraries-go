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

type KundeBytteReversert struct {
	AvtaleId string `json:"AvtaleId"`

	Hendelsestype string `json:"Hendelsestype"`

	Id string `json:"Id"`

	KundeGuidForKundeSomReverseres string `json:"KundeGuidForKundeSomReverseres"`

	MaalepunktId string `json:"MaalepunktId"`

	Startdato int64 `json:"Startdato"`
}

const KundeBytteReversertAvroCRC64Fingerprint = "\xb6\xf5N\xefcpZ\x0f"

func NewKundeBytteReversert() KundeBytteReversert {
	r := KundeBytteReversert{}
	return r
}

func DeserializeKundeBytteReversert(r io.Reader) (KundeBytteReversert, error) {
	t := NewKundeBytteReversert()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeKundeBytteReversertFromSchema(r io.Reader, schema string) (KundeBytteReversert, error) {
	t := NewKundeBytteReversert()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeKundeBytteReversert(r KundeBytteReversert, w io.Writer) error {
	var err error
	err = vm.WriteString(r.AvtaleId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Hendelsestype, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.KundeGuidForKundeSomReverseres, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MaalepunktId, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Startdato, w)
	if err != nil {
		return err
	}
	return err
}

func (r KundeBytteReversert) Serialize(w io.Writer) error {
	return writeKundeBytteReversert(r, w)
}

func (r KundeBytteReversert) Schema() string {
	return "{\"fields\":[{\"name\":\"AvtaleId\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"Hendelsestype\",\"type\":\"string\"},{\"name\":\"Id\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"KundeGuidForKundeSomReverseres\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"MaalepunktId\",\"type\":\"string\"},{\"name\":\"Startdato\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Mmm.Schemas.NettleieAvtaleHendelser.V1.KundeBytteReversert\",\"type\":\"record\"}"
}

func (r KundeBytteReversert) SchemaName() string {
	return "Mmm.Schemas.NettleieAvtaleHendelser.V1.KundeBytteReversert"
}

func (_ KundeBytteReversert) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetInt(v int32)       { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetLong(v int64)      { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetString(v string)   { panic("Unsupported operation") }
func (_ KundeBytteReversert) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *KundeBytteReversert) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.AvtaleId}

		return w

	case 1:
		w := types.String{Target: &r.Hendelsestype}

		return w

	case 2:
		w := types.String{Target: &r.Id}

		return w

	case 3:
		w := types.String{Target: &r.KundeGuidForKundeSomReverseres}

		return w

	case 4:
		w := types.String{Target: &r.MaalepunktId}

		return w

	case 5:
		w := types.Long{Target: &r.Startdato}

		return w

	}
	panic("Unknown field index")
}

func (r *KundeBytteReversert) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *KundeBytteReversert) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ KundeBytteReversert) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ KundeBytteReversert) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ KundeBytteReversert) HintSize(int)                     { panic("Unsupported operation") }
func (_ KundeBytteReversert) Finalize()                        {}

func (_ KundeBytteReversert) AvroCRC64Fingerprint() []byte {
	return []byte(KundeBytteReversertAvroCRC64Fingerprint)
}

func (r KundeBytteReversert) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["AvtaleId"], err = json.Marshal(r.AvtaleId)
	if err != nil {
		return nil, err
	}
	output["Hendelsestype"], err = json.Marshal(r.Hendelsestype)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["KundeGuidForKundeSomReverseres"], err = json.Marshal(r.KundeGuidForKundeSomReverseres)
	if err != nil {
		return nil, err
	}
	output["MaalepunktId"], err = json.Marshal(r.MaalepunktId)
	if err != nil {
		return nil, err
	}
	output["Startdato"], err = json.Marshal(r.Startdato)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *KundeBytteReversert) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["AvtaleId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AvtaleId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AvtaleId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Hendelsestype"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Hendelsestype); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Hendelsestype")
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
		if v, ok := fields["KundeGuidForKundeSomReverseres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.KundeGuidForKundeSomReverseres); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for KundeGuidForKundeSomReverseres")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MaalepunktId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MaalepunktId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MaalepunktId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Startdato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Startdato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Startdato")
	}
	return nil
}
