// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100486_7.avsc
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

type UnionNullArrayUnionNullSwitchingPlanDtoTypeEnum int

const (
	UnionNullArrayUnionNullSwitchingPlanDtoTypeEnumArrayUnionNullSwitchingPlanDto UnionNullArrayUnionNullSwitchingPlanDtoTypeEnum = 1
)

type UnionNullArrayUnionNullSwitchingPlanDto struct {
	Null                           *types.NullVal
	ArrayUnionNullSwitchingPlanDto []*UnionNullSwitchingPlanDto
	UnionType                      UnionNullArrayUnionNullSwitchingPlanDtoTypeEnum
}

func writeUnionNullArrayUnionNullSwitchingPlanDto(r *UnionNullArrayUnionNullSwitchingPlanDto, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayUnionNullSwitchingPlanDtoTypeEnumArrayUnionNullSwitchingPlanDto:
		return writeArrayUnionNullSwitchingPlanDto(r.ArrayUnionNullSwitchingPlanDto, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullSwitchingPlanDto")
}

func NewUnionNullArrayUnionNullSwitchingPlanDto() *UnionNullArrayUnionNullSwitchingPlanDto {
	return &UnionNullArrayUnionNullSwitchingPlanDto{}
}

func (r *UnionNullArrayUnionNullSwitchingPlanDto) Serialize(w io.Writer) error {
	return writeUnionNullArrayUnionNullSwitchingPlanDto(r, w)
}

func DeserializeUnionNullArrayUnionNullSwitchingPlanDto(r io.Reader) (*UnionNullArrayUnionNullSwitchingPlanDto, error) {
	t := NewUnionNullArrayUnionNullSwitchingPlanDto()
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

func DeserializeUnionNullArrayUnionNullSwitchingPlanDtoFromSchema(r io.Reader, schema string) (*UnionNullArrayUnionNullSwitchingPlanDto, error) {
	t := NewUnionNullArrayUnionNullSwitchingPlanDto()
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

func (r *UnionNullArrayUnionNullSwitchingPlanDto) Schema() string {
	return "[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"AuthorName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CreatedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastModifiedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastModifiedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Name\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Outage\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]}],\"name\":\"OutageDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Purpose\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Rank\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SwitchActions\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"CrewMember\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Person\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]}],\"name\":\"PersonDto\",\"type\":\"record\"}]}],\"name\":\"CrewMemberSDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExecutedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ExecutedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"IsFreeSequence\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IssuedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IssuedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Kind\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastModifiedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastModifiedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"OperatedSwitch\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Assets\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]}],\"name\":\"AssetSDto\",\"type\":\"record\"}],\"type\":\"array\"}]},{\"default\":null,\"name\":\"MRID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SwitchPhase\",\"type\":[\"null\",{\"items\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"PhaseSide1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PhaseSide2\",\"type\":[\"null\",\"string\"]}],\"name\":\"SwitchPhaseDto\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"SwitchDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Operator\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Person\",\"type\":[\"null\",\"SesamResponseServices.SesamDomainObjects.PersonDto\"]}],\"name\":\"OperatorDto\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"PlannedDateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PlannedDateTimeSpecified\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"SequenceNumber\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Status\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"DateTime\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"StatusSDto\",\"type\":\"record\"}]}],\"name\":\"SwitchActionDto\",\"type\":\"record\"}],\"type\":\"array\"}]}],\"name\":\"SwitchingPlanDto\",\"type\":\"record\"}],\"type\":\"array\"}]"
}

func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetBoolean(v bool)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetInt(v int32)     { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetBytes(v []byte)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullArrayUnionNullSwitchingPlanDto) SetLong(v int64) {

	r.UnionType = (UnionNullArrayUnionNullSwitchingPlanDtoTypeEnum)(v)
}

func (r *UnionNullArrayUnionNullSwitchingPlanDto) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayUnionNullSwitchingPlanDto = make([]*UnionNullSwitchingPlanDto, 0)
		return &ArrayUnionNullSwitchingPlanDtoWrapper{Target: (&r.ArrayUnionNullSwitchingPlanDto)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullArrayUnionNullSwitchingPlanDto) Finalize() {}

func (r *UnionNullArrayUnionNullSwitchingPlanDto) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayUnionNullSwitchingPlanDtoTypeEnumArrayUnionNullSwitchingPlanDto:
		return json.Marshal(map[string]interface{}{"array": r.ArrayUnionNullSwitchingPlanDto})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayUnionNullSwitchingPlanDto")
}

func (r *UnionNullArrayUnionNullSwitchingPlanDto) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayUnionNullSwitchingPlanDto)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayUnionNullSwitchingPlanDto")
}
