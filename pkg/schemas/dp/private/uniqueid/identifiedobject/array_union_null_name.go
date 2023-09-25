// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100254_5.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullName(r []*UnionNullName, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullName(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullNameWrapper struct {
	Target *[]*UnionNullName
}

func (_ ArrayUnionNullNameWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionNullNameWrapper) Finalize()                        {}
func (_ ArrayUnionNullNameWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayUnionNullNameWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullName, 0, s)
	}
}
func (r ArrayUnionNullNameWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullNameWrapper) AppendArray() types.Field {
	var v *UnionNullName
	v = NewUnionNullName()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
