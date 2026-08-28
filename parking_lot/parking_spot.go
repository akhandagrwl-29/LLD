package parkinglot

type ParkingSpot struct {
	id                   int
	isOccupied           bool
	vehicleLicenseNumber string
	vehileType           VehicleType
}

func (p *ParkingSpot) IsOccupied() bool {
	return p.isOccupied
}

func (p *ParkingSpot) GetParkingSpot() *ParkingSpot {
	return p
}
