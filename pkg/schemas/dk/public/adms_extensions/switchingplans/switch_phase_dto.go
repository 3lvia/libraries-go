// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100486_7.avsc
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

type SwitchPhaseDto struct {
	PhaseSide1 *UnionNullString `json:"PhaseSide1"`

	PhaseSide2 *UnionNullString `json:"PhaseSide2"`
}

const SwitchPhaseDtoAvroCRC64Fingerprint = "!u\xb8G\x97\xb4\x8e\xee"

func NewSwitchPhaseDto() SwitchPhaseDto {
	r := SwitchPhaseDto{}
	r.PhaseSide1 = nil
	r.PhaseSide2 = nil
	return r
}

func DeserializeSwitchPhaseDto(r io.Reader) (SwitchPhaseDto, error) {
	t := NewSwitchPhaseDto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSwitchPhaseDtoFromSchema(r io.Reader, schema string) (SwitchPhaseDto, error) {
	t := NewSwitchPhaseDto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSwitchPhaseDto(r SwitchPhaseDto, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.PhaseSide1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PhaseSide2, w)
	if err != nil {
		return err
	}
	return err
}

func (r SwitchPhaseDto) Serialize(w io.Writer) error {
	return writeSwitchPhaseDto(r, w)
}

func (r SwitchPhaseDto) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"PhaseSide1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PhaseSide2\",\"type\":[\"null\",\"string\"]}],\"name\":\"SesamResponseServices.SesamDomainObjects.SwitchPhaseDto\",\"type\":\"record\"}"
}

func (r SwitchPhaseDto) SchemaName() string {
	return "SesamResponseServices.SesamDomainObjects.SwitchPhaseDto"
}

func (_ SwitchPhaseDto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetString(v string)   { panic("Unsupported operation") }
func (_ SwitchPhaseDto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SwitchPhaseDto) Get(i int) types.Field {
	switch i {
	case 0:
		r.PhaseSide1 = NewUnionNullString()

		return r.PhaseSide1
	case 1:
		r.PhaseSide2 = NewUnionNullString()

		return r.PhaseSide2
	}
	panic("Unknown field index")
}

func (r *SwitchPhaseDto) SetDefault(i int) {
	switch i {
	case 0:
		r.PhaseSide1 = nil
		return
	case 1:
		r.PhaseSide2 = nil
		return
	}
	panic("Unknown field index")
}

func (r *SwitchPhaseDto) NullField(i int) {
	switch i {
	case 0:
		r.PhaseSide1 = nil
		return
	case 1:
		r.PhaseSide2 = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ SwitchPhaseDto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SwitchPhaseDto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SwitchPhaseDto) HintSize(int)                     { panic("Unsupported operation") }
func (_ SwitchPhaseDto) Finalize()                        {}

func (_ SwitchPhaseDto) AvroCRC64Fingerprint() []byte {
	return []byte(SwitchPhaseDtoAvroCRC64Fingerprint)
}

func (r SwitchPhaseDto) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["PhaseSide1"], err = json.Marshal(r.PhaseSide1)
	if err != nil {
		return nil, err
	}
	output["PhaseSide2"], err = json.Marshal(r.PhaseSide2)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SwitchPhaseDto) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["PhaseSide1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PhaseSide1); err != nil {
			return err
		}
	} else {
		r.PhaseSide1 = NewUnionNullString()

		r.PhaseSide1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PhaseSide2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PhaseSide2); err != nil {
			return err
		}
	} else {
		r.PhaseSide2 = NewUnionNullString()

		r.PhaseSide2 = nil
	}
	return nil
}
