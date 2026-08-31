package parking_lot

import "fmt"

var instance *ParkingLot

type ParkingLot struct {
	ID     int
	levels []*Level
}

func NewParkingLot(id int) *ParkingLot {
	if instance != nil {
		return instance
	}
	return &ParkingLot{
		ID:     id,
		levels: make([]*Level, 0),
	}
}

func (parkingLot *ParkingLot) AddLevel(level *Level) {
	parkingLot.levels = append(parkingLot.levels, level)
}

func (parkingLot *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range parkingLot.levels {
		if level.ParkVehicle(vehicle.GetLicenseNumber(), vehicle.getVehicleType()) {
			return true
		}
	}
	return false
}

func (parkingLot *ParkingLot) UnParkVehicle(vehicle Vehicle) bool {
	for _, level := range parkingLot.levels {
		if level.UnParkVehicle(vehicle.GetLicenseNumber(), vehicle.getVehicleType()) {
			return true
		}
	}
	return false
}

func (parkingLot *ParkingLot) ShowAvailability() {
	for _, level := range parkingLot.levels {
		fmt.Printf("Parking Lot ID: %d and level : %d\n", parkingLot.ID, level.floor)
		for _, spot := range level.parkingSpots {
			status := "empty"
			if spot.IsOccupied {
				status = "occupied"
			}
			fmt.Printf("Spot: %d, status: %s, licenseNumber: %s\n", spot.ID, status, string(spot.LicenseNumber))
		}
	}
}
