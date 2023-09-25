// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100058_21.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PlacementCode int32

const (
	PlacementCodeUnknown                      PlacementCode = 0
	PlacementCodeInsideStation                PlacementCode = 1
	PlacementCodeCabinetOutsideOnStation      PlacementCode = 2
	PlacementCodeCabinetOutsideOnPole         PlacementCode = 3
	PlacementCodeCabinetInside                PlacementCode = 4
	PlacementCodeCabinetInsideStation         PlacementCode = 5
	PlacementCodeCabinetOutsideOnCableCabinet PlacementCode = 6
	PlacementCodeIntegratedInCabinet          PlacementCode = 7
	PlacementCodeStore                        PlacementCode = 8
	PlacementCodeBackstairs                   PlacementCode = 9
	PlacementCodeEntrance                     PlacementCode = 10
	PlacementCodeFacade                       PlacementCode = 11
	PlacementCodeBoilerRoom                   PlacementCode = 12
	PlacementCodeGarage                       PlacementCode = 13
	PlacementCodeElevatorRoom                 PlacementCode = 14
	PlacementCodeCellar                       PlacementCode = 15
	PlacementCodeOffice                       PlacementCode = 16
	PlacementCodeStorageroom                  PlacementCode = 17
	PlacementCodeAttic                        PlacementCode = 18
	PlacementCodeBoardroom                    PlacementCode = 19
	PlacementCodeShelter                      PlacementCode = 20
	PlacementCodeProperty                     PlacementCode = 21
	PlacementCodeWashingroom                  PlacementCode = 22
	PlacementCodeCabinetInEntrance            PlacementCode = 23
	PlacementCodeWorkshop                     PlacementCode = 24
	PlacementCodeCabin                        PlacementCode = 25
	PlacementCodeOuthouse                     PlacementCode = 26
	PlacementCodeKLFacilities                 PlacementCode = 27
	PlacementCodePresumedInside               PlacementCode = 28
	PlacementCodeCompanyCommonArea            PlacementCode = 29
	PlacementCodeCompanyInAppartement         PlacementCode = 30
	PlacementCodeCompanyGround                PlacementCode = 31
	PlacementCodeHouse1stFloor                PlacementCode = 32
	PlacementCodeHouse2FloorOrHigher          PlacementCode = 33
	PlacementCodeHouseBasementUnderGround     PlacementCode = 34
	PlacementCodeEntranceHall                 PlacementCode = 35
	PlacementCodeNotInUseManhole              PlacementCode = 36
	PlacementCodeNotInUseOuthouse             PlacementCode = 37
	PlacementCodeNotInUseOutsideShed          PlacementCode = 38
	PlacementCodeCableCabinet                 PlacementCode = 39
	PlacementCodeBasement                     PlacementCode = 40
	PlacementCodeNS                           PlacementCode = 41
	PlacementCodeTechnicalRoom                PlacementCode = 42
	PlacementCodeStaircase                    PlacementCode = 43
	PlacementCodeOutdoorsFreeAccess           PlacementCode = 44
	PlacementCodeOuthouseLocked               PlacementCode = 45
	PlacementCodeOutside                      PlacementCode = 46
	PlacementCodeOutsideKiosk                 PlacementCode = 47
	PlacementCodeOutsideCabinet               PlacementCode = 48
)

func (e PlacementCode) String() string {
	switch e {
	case PlacementCodeUnknown:
		return "Unknown"
	case PlacementCodeInsideStation:
		return "InsideStation"
	case PlacementCodeCabinetOutsideOnStation:
		return "CabinetOutsideOnStation"
	case PlacementCodeCabinetOutsideOnPole:
		return "CabinetOutsideOnPole"
	case PlacementCodeCabinetInside:
		return "CabinetInside"
	case PlacementCodeCabinetInsideStation:
		return "CabinetInsideStation"
	case PlacementCodeCabinetOutsideOnCableCabinet:
		return "CabinetOutsideOnCableCabinet"
	case PlacementCodeIntegratedInCabinet:
		return "IntegratedInCabinet"
	case PlacementCodeStore:
		return "Store"
	case PlacementCodeBackstairs:
		return "Backstairs"
	case PlacementCodeEntrance:
		return "Entrance"
	case PlacementCodeFacade:
		return "Facade"
	case PlacementCodeBoilerRoom:
		return "BoilerRoom"
	case PlacementCodeGarage:
		return "Garage"
	case PlacementCodeElevatorRoom:
		return "ElevatorRoom"
	case PlacementCodeCellar:
		return "Cellar"
	case PlacementCodeOffice:
		return "Office"
	case PlacementCodeStorageroom:
		return "Storageroom"
	case PlacementCodeAttic:
		return "Attic"
	case PlacementCodeBoardroom:
		return "Boardroom"
	case PlacementCodeShelter:
		return "Shelter"
	case PlacementCodeProperty:
		return "Property"
	case PlacementCodeWashingroom:
		return "Washingroom"
	case PlacementCodeCabinetInEntrance:
		return "CabinetInEntrance"
	case PlacementCodeWorkshop:
		return "Workshop"
	case PlacementCodeCabin:
		return "Cabin"
	case PlacementCodeOuthouse:
		return "Outhouse"
	case PlacementCodeKLFacilities:
		return "KLFacilities"
	case PlacementCodePresumedInside:
		return "PresumedInside"
	case PlacementCodeCompanyCommonArea:
		return "CompanyCommonArea"
	case PlacementCodeCompanyInAppartement:
		return "CompanyInAppartement"
	case PlacementCodeCompanyGround:
		return "CompanyGround"
	case PlacementCodeHouse1stFloor:
		return "House1stFloor"
	case PlacementCodeHouse2FloorOrHigher:
		return "House2FloorOrHigher"
	case PlacementCodeHouseBasementUnderGround:
		return "HouseBasementUnderGround"
	case PlacementCodeEntranceHall:
		return "EntranceHall"
	case PlacementCodeNotInUseManhole:
		return "NotInUseManhole"
	case PlacementCodeNotInUseOuthouse:
		return "NotInUseOuthouse"
	case PlacementCodeNotInUseOutsideShed:
		return "NotInUseOutsideShed"
	case PlacementCodeCableCabinet:
		return "CableCabinet"
	case PlacementCodeBasement:
		return "Basement"
	case PlacementCodeNS:
		return "NS"
	case PlacementCodeTechnicalRoom:
		return "TechnicalRoom"
	case PlacementCodeStaircase:
		return "Staircase"
	case PlacementCodeOutdoorsFreeAccess:
		return "OutdoorsFreeAccess"
	case PlacementCodeOuthouseLocked:
		return "OuthouseLocked"
	case PlacementCodeOutside:
		return "Outside"
	case PlacementCodeOutsideKiosk:
		return "OutsideKiosk"
	case PlacementCodeOutsideCabinet:
		return "OutsideCabinet"
	}
	return "unknown"
}

func writePlacementCode(r PlacementCode, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewPlacementCodeValue(raw string) (r PlacementCode, err error) {
	switch raw {
	case "Unknown":
		return PlacementCodeUnknown, nil
	case "InsideStation":
		return PlacementCodeInsideStation, nil
	case "CabinetOutsideOnStation":
		return PlacementCodeCabinetOutsideOnStation, nil
	case "CabinetOutsideOnPole":
		return PlacementCodeCabinetOutsideOnPole, nil
	case "CabinetInside":
		return PlacementCodeCabinetInside, nil
	case "CabinetInsideStation":
		return PlacementCodeCabinetInsideStation, nil
	case "CabinetOutsideOnCableCabinet":
		return PlacementCodeCabinetOutsideOnCableCabinet, nil
	case "IntegratedInCabinet":
		return PlacementCodeIntegratedInCabinet, nil
	case "Store":
		return PlacementCodeStore, nil
	case "Backstairs":
		return PlacementCodeBackstairs, nil
	case "Entrance":
		return PlacementCodeEntrance, nil
	case "Facade":
		return PlacementCodeFacade, nil
	case "BoilerRoom":
		return PlacementCodeBoilerRoom, nil
	case "Garage":
		return PlacementCodeGarage, nil
	case "ElevatorRoom":
		return PlacementCodeElevatorRoom, nil
	case "Cellar":
		return PlacementCodeCellar, nil
	case "Office":
		return PlacementCodeOffice, nil
	case "Storageroom":
		return PlacementCodeStorageroom, nil
	case "Attic":
		return PlacementCodeAttic, nil
	case "Boardroom":
		return PlacementCodeBoardroom, nil
	case "Shelter":
		return PlacementCodeShelter, nil
	case "Property":
		return PlacementCodeProperty, nil
	case "Washingroom":
		return PlacementCodeWashingroom, nil
	case "CabinetInEntrance":
		return PlacementCodeCabinetInEntrance, nil
	case "Workshop":
		return PlacementCodeWorkshop, nil
	case "Cabin":
		return PlacementCodeCabin, nil
	case "Outhouse":
		return PlacementCodeOuthouse, nil
	case "KLFacilities":
		return PlacementCodeKLFacilities, nil
	case "PresumedInside":
		return PlacementCodePresumedInside, nil
	case "CompanyCommonArea":
		return PlacementCodeCompanyCommonArea, nil
	case "CompanyInAppartement":
		return PlacementCodeCompanyInAppartement, nil
	case "CompanyGround":
		return PlacementCodeCompanyGround, nil
	case "House1stFloor":
		return PlacementCodeHouse1stFloor, nil
	case "House2FloorOrHigher":
		return PlacementCodeHouse2FloorOrHigher, nil
	case "HouseBasementUnderGround":
		return PlacementCodeHouseBasementUnderGround, nil
	case "EntranceHall":
		return PlacementCodeEntranceHall, nil
	case "NotInUseManhole":
		return PlacementCodeNotInUseManhole, nil
	case "NotInUseOuthouse":
		return PlacementCodeNotInUseOuthouse, nil
	case "NotInUseOutsideShed":
		return PlacementCodeNotInUseOutsideShed, nil
	case "CableCabinet":
		return PlacementCodeCableCabinet, nil
	case "Basement":
		return PlacementCodeBasement, nil
	case "NS":
		return PlacementCodeNS, nil
	case "TechnicalRoom":
		return PlacementCodeTechnicalRoom, nil
	case "Staircase":
		return PlacementCodeStaircase, nil
	case "OutdoorsFreeAccess":
		return PlacementCodeOutdoorsFreeAccess, nil
	case "OuthouseLocked":
		return PlacementCodeOuthouseLocked, nil
	case "Outside":
		return PlacementCodeOutside, nil
	case "OutsideKiosk":
		return PlacementCodeOutsideKiosk, nil
	case "OutsideCabinet":
		return PlacementCodeOutsideCabinet, nil
	}

	return -1, fmt.Errorf("invalid value for PlacementCode: '%s'", raw)

}

func (b PlacementCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *PlacementCode) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewPlacementCodeValue(stringVal)
	*b = val
	return err
}

type PlacementCodeWrapper struct {
	Target *PlacementCode
}

func (b PlacementCodeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b PlacementCodeWrapper) SetInt(v int32) {
	*(b.Target) = PlacementCode(v)
}

func (b PlacementCodeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b PlacementCodeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b PlacementCodeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b PlacementCodeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b PlacementCodeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b PlacementCodeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b PlacementCodeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b PlacementCodeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b PlacementCodeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b PlacementCodeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b PlacementCodeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b PlacementCodeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b PlacementCodeWrapper) Finalize() {}
