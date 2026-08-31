package parking_lot

type Level struct {
	parkingSpots []*ParkingSpot
	floor        int
}

func NewLevel(floor int, numSpots int) *Level {
	level := &Level{
		floor: floor,
	}
	carSpots := int(0.3 * float64(numSpots))
	motorcycleSpots := int(0.5 * float64(numSpots))
	for i := 0; i < carSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, &ParkingSpot{
			VehicleType: CAR,
			ID:          i + 1,
		})
	}

	for i := carSpots; i < carSpots+motorcycleSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, &ParkingSpot{
			VehicleType: MOTORCYCLE,
			ID:          i + 1,
		})
	}

	for i := carSpots + motorcycleSpots; i < numSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, &ParkingSpot{
			VehicleType: TRUCK,
			ID:          i + 1,
		})
	}
	return level
}

func (level *Level) ParkVehicle(licenseNumber string, vehicleType VehicleType) bool {
	for _, spot := range level.parkingSpots {
		if spot.IsOccupied {
			continue
		}
		if spot.VehicleType == vehicleType {
			spot.IsOccupied = true
			spot.VehicleType = vehicleType
			spot.LicenseNumber = licenseNumber
			return true
		}
	}
	return false
}

func (level *Level) UnParkVehicle(licenseNumber string, vehicleType VehicleType) bool {
	for _, spot := range level.parkingSpots {
		if !spot.IsOccupied {
			continue
		}
		if spot.LicenseNumber == licenseNumber {
			spot.IsOccupied = false
			spot.LicenseNumber = ""
			return true
		}
	}
	return false
}
