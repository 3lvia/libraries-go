// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100487_5.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullWorkAssetDto(r []*UnionNullWorkAssetDto, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullWorkAssetDto(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullWorkAssetDtoWrapper struct {
	Target *[]*UnionNullWorkAssetDto
}

func (_ ArrayUnionNullWorkAssetDtoWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionNullWorkAssetDtoWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullWorkAssetDtoWrapper) Finalize()        {}
func (_ ArrayUnionNullWorkAssetDtoWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUnionNullWorkAssetDtoWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullWorkAssetDto, 0, s)
	}
}
func (r ArrayUnionNullWorkAssetDtoWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullWorkAssetDtoWrapper) AppendArray() types.Field {
	var v *UnionNullWorkAssetDto
	v = NewUnionNullWorkAssetDto()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
