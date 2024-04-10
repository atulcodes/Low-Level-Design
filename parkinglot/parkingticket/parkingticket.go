package parkingticket

import "time"

type ParkingTicket struct {
	id string
	vehicleId string
	inTime time.Time
	outTime time.Time
	parkingCost float32
}

