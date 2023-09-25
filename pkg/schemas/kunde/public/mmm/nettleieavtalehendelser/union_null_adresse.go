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

type UnionNullAdresseTypeEnum int

const (
	UnionNullAdresseTypeEnumAdresse UnionNullAdresseTypeEnum = 1
)

type UnionNullAdresse struct {
	Null      *types.NullVal
	Adresse   Adresse
	UnionType UnionNullAdresseTypeEnum
}

func writeUnionNullAdresse(r *UnionNullAdresse, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullAdresseTypeEnumAdresse:
		return writeAdresse(r.Adresse, w)
	}
	return fmt.Errorf("invalid value for *UnionNullAdresse")
}

func NewUnionNullAdresse() *UnionNullAdresse {
	return &UnionNullAdresse{}
}

func (r *UnionNullAdresse) Serialize(w io.Writer) error {
	return writeUnionNullAdresse(r, w)
}

func DeserializeUnionNullAdresse(r io.Reader) (*UnionNullAdresse, error) {
	t := NewUnionNullAdresse()
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

func DeserializeUnionNullAdresseFromSchema(r io.Reader, schema string) (*UnionNullAdresse, error) {
	t := NewUnionNullAdresse()
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

func (r *UnionNullAdresse) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Bruksenhetsnummer\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CoAdresse\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Etasje\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Gatenavn\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Husnummer\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Kommunenummer\",\"type\":[\"null\",\"string\"]},{\"name\":\"Landkode\",\"type\":\"string\"},{\"default\":null,\"name\":\"Postboks\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PostboksAnleggsnavn\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Postnummer\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Poststed\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PaaVegneAv\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"VedPerson\",\"type\":[\"null\",\"string\"]}],\"name\":\"Adresse\",\"type\":\"record\"}]"
}

func (_ *UnionNullAdresse) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullAdresse) SetLong(v int64) {

	r.UnionType = (UnionNullAdresseTypeEnum)(v)
}

func (r *UnionNullAdresse) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Adresse = NewAdresse()
		return &types.Record{Target: (&r.Adresse)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullAdresse) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullAdresse) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullAdresse) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullAdresse) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullAdresse) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullAdresse) Finalize()                        {}

func (r *UnionNullAdresse) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullAdresseTypeEnumAdresse:
		return json.Marshal(map[string]interface{}{"Mmm.Schemas.NettleieAvtaleHendelser.V1.Adresse": r.Adresse})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullAdresse")
}

func (r *UnionNullAdresse) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Mmm.Schemas.NettleieAvtaleHendelser.V1.Adresse"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Adresse)
	}
	return fmt.Errorf("invalid value for *UnionNullAdresse")
}
