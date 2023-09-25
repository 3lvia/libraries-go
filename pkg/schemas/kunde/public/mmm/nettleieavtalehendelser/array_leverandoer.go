// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100480_13.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayLeverandoer(r []Leverandoer, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeLeverandoer(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayLeverandoerWrapper struct {
	Target *[]Leverandoer
}

func (_ ArrayLeverandoerWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayLeverandoerWrapper) Finalize()                        {}
func (_ ArrayLeverandoerWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayLeverandoerWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Leverandoer, 0, s)
	}
}
func (r ArrayLeverandoerWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayLeverandoerWrapper) AppendArray() types.Field {
	var v Leverandoer
	v = NewLeverandoer()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
