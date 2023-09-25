// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100044_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayHouse(r []House, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeHouse(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayHouseWrapper struct {
	Target *[]House
}

func (_ ArrayHouseWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayHouseWrapper) Finalize()                        {}
func (_ ArrayHouseWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayHouseWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]House, 0, s)
	}
}
func (r ArrayHouseWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayHouseWrapper) AppendArray() types.Field {
	var v House
	v = NewHouse()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
