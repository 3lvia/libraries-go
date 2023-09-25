// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100242_2.avsc
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

type UnionNullCompanyTypeEnum int

const (
	UnionNullCompanyTypeEnumCompany UnionNullCompanyTypeEnum = 1
)

type UnionNullCompany struct {
	Null      *types.NullVal
	Company   Company
	UnionType UnionNullCompanyTypeEnum
}

func writeUnionNullCompany(r *UnionNullCompany, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullCompanyTypeEnumCompany:
		return writeCompany(r.Company, w)
	}
	return fmt.Errorf("invalid value for *UnionNullCompany")
}

func NewUnionNullCompany() *UnionNullCompany {
	return &UnionNullCompany{}
}

func (r *UnionNullCompany) Serialize(w io.Writer) error {
	return writeUnionNullCompany(r, w)
}

func DeserializeUnionNullCompany(r io.Reader) (*UnionNullCompany, error) {
	t := NewUnionNullCompany()
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

func DeserializeUnionNullCompanyFromSchema(r io.Reader, schema string) (*UnionNullCompany, error) {
	t := NewUnionNullCompany()
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

func (r *UnionNullCompany) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"CompanyName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CompanyOrgNo\",\"type\":[\"null\",\"string\"]}],\"name\":\"Company\",\"type\":\"record\"}]"
}

func (_ *UnionNullCompany) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullCompany) SetLong(v int64) {

	r.UnionType = (UnionNullCompanyTypeEnum)(v)
}

func (r *UnionNullCompany) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Company = NewCompany()
		return &types.Record{Target: (&r.Company)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullCompany) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullCompany) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullCompany) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullCompany) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullCompany) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullCompany) Finalize()                        {}

func (r *UnionNullCompany) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullCompanyTypeEnumCompany:
		return json.Marshal(map[string]interface{}{"GridTariff.Events.V1.Company": r.Company})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullCompany")
}

func (r *UnionNullCompany) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["GridTariff.Events.V1.Company"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Company)
	}
	return fmt.Errorf("invalid value for *UnionNullCompany")
}
