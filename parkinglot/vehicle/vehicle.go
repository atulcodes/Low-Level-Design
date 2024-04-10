package vehicle

import (
	"fmt"

	"github.com/lowleveldesign/parkinglot/common"
)

type IVehicle interface {
	ParkVehicle() bool
	UnparkVehicle() bool
}

type Vehicle struct {
	Id string
	NumberPlate string
	Type common.VehicleType
}

func (v *Vehicle) ParkVehicle() bool {return false}

func (v *Vehicle) UnparkVehicle() bool {return false}

func VehicleFactory(vehicleType common.VehicleType) (IVehicle, error) {
	switch vehicleType {
		case common.TwoWheeler:
			return &TwoWheeler{
				Vehicle: Vehicle{
					Id:          common.GenerateUniqueID(),
					NumberPlate: common.GenerateRandomNumberPlate(),
					Type:        common.TwoWheeler,
				},
			}, nil
		case common.ThreeWheeler:
			return &TwoWheeler{
				Vehicle: Vehicle{
					Id:          common.GenerateUniqueID(),
					NumberPlate: common.GenerateRandomNumberPlate(),
					Type:        common.ThreeWheeler,
				},
			}, nil
		case common.FourWheeler:
			return &TwoWheeler{
				Vehicle: Vehicle{
					Id:          common.GenerateUniqueID(),
					NumberPlate: common.GenerateRandomNumberPlate(),
					Type:        common.FourWheeler,
				},
			}, nil
	}

	return nil, fmt.Errorf("unknown vehicle type %s", vehicleType)
}