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

var _ = fmt.Printf

type Grid struct {
	Created string `json:"Created"`

	FunctionalObjectId *UnionNullString `json:"FunctionalObjectId"`

	GridArea *UnionNullGridArea `json:"GridArea"`

	GridId *UnionNullString `json:"GridId"`

	GridOwner *UnionNullString `json:"GridOwner"`

	Updated *UnionNullString `json:"Updated"`
}

const GridAvroCRC64Fingerprint = "\xfc\x93\x8e2\x93v\b\x02"

func NewGrid() Grid {
	r := Grid{}
	r.FunctionalObjectId = nil
	r.GridArea = nil
	r.GridId = nil
	r.GridOwner = nil
	r.Updated = nil
	return r
}

func DeserializeGrid(r io.Reader) (Grid, error) {
	t := NewGrid()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeGridFromSchema(r io.Reader, schema string) (Grid, error) {
	t := NewGrid()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeGrid(r Grid, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Created, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FunctionalObjectId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullGridArea(r.GridArea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GridId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GridOwner, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Updated, w)
	if err != nil {
		return err
	}
	return err
}

func (r Grid) Serialize(w io.Writer) error {
	return writeGrid(r, w)
}

func (r Grid) Schema() string {
	return "{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FunctionalObjectId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridArea\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Created\",\"type\":\"string\"},{\"default\":null,\"name\":\"FlowDirection\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridAreaId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"GridArea\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GridId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"GridOwner\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Updated\",\"type\":[\"null\",\"string\"]}],\"name\":\"Skogshorn.Nett.Schemas.Primitives.Grid\",\"type\":\"record\"}"
}

func (r Grid) SchemaName() string {
	return "Skogshorn.Nett.Schemas.Primitives.Grid"
}

func (_ Grid) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Grid) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Grid) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Grid) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Grid) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Grid) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Grid) SetString(v string)   { panic("Unsupported operation") }
func (_ Grid) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Grid) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Created}

		return w

	case 1:
		r.FunctionalObjectId = NewUnionNullString()

		return r.FunctionalObjectId
	case 2:
		r.GridArea = NewUnionNullGridArea()

		return r.GridArea
	case 3:
		r.GridId = NewUnionNullString()

		return r.GridId
	case 4:
		r.GridOwner = NewUnionNullString()

		return r.GridOwner
	case 5:
		r.Updated = NewUnionNullString()

		return r.Updated
	}
	panic("Unknown field index")
}

func (r *Grid) SetDefault(i int) {
	switch i {
	case 1:
		r.FunctionalObjectId = nil
		return
	case 2:
		r.GridArea = nil
		return
	case 3:
		r.GridId = nil
		return
	case 4:
		r.GridOwner = nil
		return
	case 5:
		r.Updated = nil
		return
	}
	panic("Unknown field index")
}

func (r *Grid) NullField(i int) {
	switch i {
	case 1:
		r.FunctionalObjectId = nil
		return
	case 2:
		r.GridArea = nil
		return
	case 3:
		r.GridId = nil
		return
	case 4:
		r.GridOwner = nil
		return
	case 5:
		r.Updated = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Grid) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Grid) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Grid) HintSize(int)                     { panic("Unsupported operation") }
func (_ Grid) Finalize()                        {}

func (_ Grid) AvroCRC64Fingerprint() []byte {
	return []byte(GridAvroCRC64Fingerprint)
}

func (r Grid) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Created"], err = json.Marshal(r.Created)
	if err != nil {
		return nil, err
	}
	output["FunctionalObjectId"], err = json.Marshal(r.FunctionalObjectId)
	if err != nil {
		return nil, err
	}
	output["GridArea"], err = json.Marshal(r.GridArea)
	if err != nil {
		return nil, err
	}
	output["GridId"], err = json.Marshal(r.GridId)
	if err != nil {
		return nil, err
	}
	output["GridOwner"], err = json.Marshal(r.GridOwner)
	if err != nil {
		return nil, err
	}
	output["Updated"], err = json.Marshal(r.Updated)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Grid) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Created"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Created); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Created")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FunctionalObjectId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FunctionalObjectId); err != nil {
			return err
		}
	} else {
		r.FunctionalObjectId = NewUnionNullString()

		r.FunctionalObjectId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["GridArea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GridArea); err != nil {
			return err
		}
	} else {
		r.GridArea = NewUnionNullGridArea()

		r.GridArea = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["GridId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GridId); err != nil {
			return err
		}
	} else {
		r.GridId = NewUnionNullString()

		r.GridId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["GridOwner"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GridOwner); err != nil {
			return err
		}
	} else {
		r.GridOwner = NewUnionNullString()

		r.GridOwner = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Updated"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Updated); err != nil {
			return err
		}
	} else {
		r.Updated = NewUnionNullString()

		r.Updated = nil
	}
	return nil
}
