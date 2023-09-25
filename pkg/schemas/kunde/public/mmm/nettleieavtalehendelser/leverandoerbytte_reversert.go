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

type LeverandoerbytteReversert struct {
	AvtaleId string `json:"AvtaleId"`

	Hendelsestype string `json:"Hendelsestype"`

	Id string `json:"Id"`

	KundeGuid string `json:"KundeGuid"`

	MaalepunktId string `json:"MaalepunktId"`

	Startdato int64 `json:"Startdato"`
}

const LeverandoerbytteReversertAvroCRC64Fingerprint = "Z\xe4\x02\xfd\x9d\xeap\xf5"

func NewLeverandoerbytteReversert() LeverandoerbytteReversert {
	r := LeverandoerbytteReversert{}
	return r
}

func DeserializeLeverandoerbytteReversert(r io.Reader) (LeverandoerbytteReversert, error) {
	t := NewLeverandoerbytteReversert()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLeverandoerbytteReversertFromSchema(r io.Reader, schema string) (LeverandoerbytteReversert, error) {
	t := NewLeverandoerbytteReversert()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLeverandoerbytteReversert(r LeverandoerbytteReversert, w io.Writer) error {
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
	err = vm.WriteString(r.KundeGuid, w)
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

func (r LeverandoerbytteReversert) Serialize(w io.Writer) error {
	return writeLeverandoerbytteReversert(r, w)
}

func (r LeverandoerbytteReversert) Schema() string {
	return "{\"fields\":[{\"name\":\"AvtaleId\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"Hendelsestype\",\"type\":\"string\"},{\"name\":\"Id\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"KundeGuid\",\"type\":{\"logicalType\":\"uuid\",\"type\":\"string\"}},{\"name\":\"MaalepunktId\",\"type\":\"string\"},{\"name\":\"Startdato\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Mmm.Schemas.NettleieAvtaleHendelser.V1.LeverandoerbytteReversert\",\"type\":\"record\"}"
}

func (r LeverandoerbytteReversert) SchemaName() string {
	return "Mmm.Schemas.NettleieAvtaleHendelser.V1.LeverandoerbytteReversert"
}

func (_ LeverandoerbytteReversert) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetInt(v int32)       { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetLong(v int64)      { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetString(v string)   { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *LeverandoerbytteReversert) Get(i int) types.Field {
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
		w := types.String{Target: &r.KundeGuid}

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

func (r *LeverandoerbytteReversert) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *LeverandoerbytteReversert) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ LeverandoerbytteReversert) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) HintSize(int)                     { panic("Unsupported operation") }
func (_ LeverandoerbytteReversert) Finalize()                        {}

func (_ LeverandoerbytteReversert) AvroCRC64Fingerprint() []byte {
	return []byte(LeverandoerbytteReversertAvroCRC64Fingerprint)
}

func (r LeverandoerbytteReversert) MarshalJSON() ([]byte, error) {
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
	output["KundeGuid"], err = json.Marshal(r.KundeGuid)
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

func (r *LeverandoerbytteReversert) UnmarshalJSON(data []byte) error {
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
