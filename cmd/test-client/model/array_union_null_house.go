// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100172_6.avsc
 */
package model

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullHouse(r []*UnionNullHouse, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullHouse(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullHouseWrapper struct {
	Target *[]*UnionNullHouse
}

func (_ ArrayUnionNullHouseWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionNullHouseWrapper) Finalize()                        {}
func (_ ArrayUnionNullHouseWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayUnionNullHouseWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullHouse, 0, s)
	}
}
func (r ArrayUnionNullHouseWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullHouseWrapper) AppendArray() types.Field {
	var v *UnionNullHouse
	v = NewUnionNullHouse()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}