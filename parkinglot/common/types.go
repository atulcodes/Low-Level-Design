package common

type VehicleType string

// Types of vehicles
const (
	TwoWheeler VehicleType = "twowheeler"
	ThreeWheeler VehicleType = "threewheeler"
	FourWheeler VehicleType = "fourwheeler"
)

type ParkingFee float32

// Parking fee per hour
const (
	TwoWheelerParkingFee ParkingFee = 15.0
	ThreeWheelerParkingFee ParkingFee = 30.0
	FourWheelerParkingFee ParkingFee = 50.0
)