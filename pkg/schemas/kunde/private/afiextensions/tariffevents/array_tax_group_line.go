// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100444_4.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayTaxGroupLine(r []TaxGroupLine, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeTaxGroupLine(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayTaxGroupLineWrapper struct {
	Target *[]TaxGroupLine
}

func (_ ArrayTaxGroupLineWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayTaxGroupLineWrapper) Finalize()                        {}
func (_ ArrayTaxGroupLineWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayTaxGroupLineWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]TaxGroupLine, 0, s)
	}
}
func (r ArrayTaxGroupLineWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayTaxGroupLineWrapper) AppendArray() types.Field {
	var v TaxGroupLine
	v = NewTaxGroupLine()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
