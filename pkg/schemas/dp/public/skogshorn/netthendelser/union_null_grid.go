// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100147_24.avsc
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

type UnionNullGridTypeEnum int

const (
	UnionNullGridTypeEnumGrid UnionNullGridTypeEnum = 1
)

type UnionNullGrid struct {
	Null      *types.NullVal
	Grid      Grid
	UnionType UnionNullGridTypeEnum
}

func writeUnionNullGrid(r *UnionNullGrid, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullGridTypeEnumGrid:
		return writeGrid(r.Grid, w)
	}
	return fmt.Errorf("invalid value for *UnionNullGrid")
}

func NewUnionNullGrid() *UnionNullGrid {
	return &UnionNullGrid{}
}

func (r *UnionNullGrid) Serialize(w io.Writer) error {
	return writeUnionNullGrid(r, w)
}

func DeserializeUnionNullGrid(r io.Reader) (*UnionNullGrid, error) {
	t := NewUnionNullGrid()
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

func DeserializeUnionNullGridFromSchema(r io.Reader, schema string) (*UnionNullGrid, error) {
	t := NewUnionNullGrid()
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

func (r *UnionNullGrid) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Grid\",\"namespace\":\"Skogshorn.Nett.Schemas.Primitives\",\"type\":\"record\"}]"
}

func (_ *UnionNullGrid) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullGrid) SetLong(v int64) {

	r.UnionType = (UnionNullGridTypeEnum)(v)
}

func (r *UnionNullGrid) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Grid = NewGrid()
		return &types.Record{Target: (&r.Grid)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullGrid) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullGrid) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullGrid) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullGrid) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullGrid) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullGrid) Finalize()                        {}

func (r *UnionNullGrid) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullGridTypeEnumGrid:
		return json.Marshal(map[string]interface{}{"Skogshorn.Nett.Schemas.Primitives.Grid": r.Grid})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullGrid")
}

func (r *UnionNullGrid) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Skogshorn.Nett.Schemas.Primitives.Grid"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Grid)
	}
	return fmt.Errorf("invalid value for *UnionNullGrid")
}
