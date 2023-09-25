// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100450_1.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullEnergizedUsagePoint(r []*UnionNullEnergizedUsagePoint, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullEnergizedUsagePoint(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullEnergizedUsagePointWrapper struct {
	Target *[]*UnionNullEnergizedUsagePoint
}

func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetBoolean(v bool)  { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetInt(v int32)     { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetLong(v int64)    { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetBytes(v []byte)  { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetUnionElem(v int64) {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullEnergizedUsagePointWrapper) Get(i int) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullEnergizedUsagePointWrapper) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ArrayUnionNullEnergizedUsagePointWrapper) Finalize()        {}
func (_ ArrayUnionNullEnergizedUsagePointWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r ArrayUnionNullEnergizedUsagePointWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullEnergizedUsagePoint, 0, s)
	}
}
func (r ArrayUnionNullEnergizedUsagePointWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullEnergizedUsagePointWrapper) AppendArray() types.Field {
	var v *UnionNullEnergizedUsagePoint
	v = NewUnionNullEnergizedUsagePoint()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
