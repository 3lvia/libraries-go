// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100147_1.avsc
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

type UnionNullPlantTypeEnum int

const (
	UnionNullPlantTypeEnumPlant UnionNullPlantTypeEnum = 1
)

type UnionNullPlant struct {
	Null      *types.NullVal
	Plant     Plant
	UnionType UnionNullPlantTypeEnum
}

func writeUnionNullPlant(r *UnionNullPlant, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullPlantTypeEnumPlant:
		return writePlant(r.Plant, w)
	}
	return fmt.Errorf("invalid value for *UnionNullPlant")
}

func NewUnionNullPlant() *UnionNullPlant {
	return &UnionNullPlant{}
}

func (r *UnionNullPlant) Serialize(w io.Writer) error {
	return writeUnionNullPlant(r, w)
}

func DeserializeUnionNullPlant(r io.Reader) (*UnionNullPlant, error) {
	t := NewUnionNullPlant()
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

func DeserializeUnionNullPlantFromSchema(r io.Reader, schema string) (*UnionNullPlant, error) {
	t := NewUnionNullPlant()
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

func (r *UnionNullPlant) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"name\":\"FuseSize\",\"type\":\"int\"},{\"name\":\"NominalVoltage\",\"type\":\"int\"},{\"name\":\"Phases\",\"type\":\"int\"},{\"default\":null,\"name\":\"PlantId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Plant\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]"
}

func (_ *UnionNullPlant) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullPlant) SetLong(v int64) {

	r.UnionType = (UnionNullPlantTypeEnum)(v)
}

func (r *UnionNullPlant) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Plant = NewPlant()
		return &types.Record{Target: (&r.Plant)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullPlant) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullPlant) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullPlant) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullPlant) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullPlant) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullPlant) Finalize()                        {}

func (r *UnionNullPlant) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullPlantTypeEnumPlant:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Schemas.Primitives.Plant": r.Plant})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullPlant")
}

func (r *UnionNullPlant) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Schemas.Primitives.Plant"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Plant)
	}
	return fmt.Errorf("invalid value for *UnionNullPlant")
}
