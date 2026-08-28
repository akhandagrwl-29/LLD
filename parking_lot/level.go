package parkinglot

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor int, numSpots int) *Level {
	l := &Level{}
	l.floor = floor
	carSpots := int(float32(numSpots) * 0.5)
	motorcycleSpots := int(float32(numSpots) * 0.3)

	for i := 0; i < carSpots; i++ {
		l.parkingSpots = append(l.parkingSpots, &ParkingSpot{id: i + 1, vehileType: CAR})
	}
	for i := carSpots; i < carSpots+motorcycleSpots; i++ {
		l.parkingSpots = append(l.parkingSpots, &ParkingSpot{id: i + 1, vehileType: MOTORCYCLE})
	}
	for i := carSpots + motorcycleSpots; i < numSpots; i++ {
		l.parkingSpots = append(l.parkingSpots, &ParkingSpot{id: i + 1, vehileType: TRUCK})
	}
	return l
}

func (l *Level) ParkVehicle(veh Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.IsOccupied() {
			continue
		}
		if spot.vehileType == veh.GetVehicleType() {
			spot.isOccupied = true
			spot.vehicleLicenseNumber = veh.GetLicensePlate()
			return true
		}
	}
	return false
}

func (l *Level) UnParkVehicle(veh Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if !spot.IsOccupied() {
			continue
		}
		if spot.vehicleLicenseNumber == veh.GetLicensePlate() {
			spot.isOccupied = false
			spot.vehicleLicenseNumber = ""
			return true
		}
	}
	return false
}
