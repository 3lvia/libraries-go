// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100159_3.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullContract(r []*UnionNullContract, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullContract(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullContractWrapper struct {
	Target *[]*UnionNullContract
}

func (_ ArrayUnionNullContractWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionNullContractWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullContractWrapper) Finalize()        {}
func (_ ArrayUnionNullContractWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUnionNullContractWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullContract, 0, s)
	}
}
func (r ArrayUnionNullContractWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullContractWrapper) AppendArray() types.Field {
	var v *UnionNullContract
	v = NewUnionNullContract()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
