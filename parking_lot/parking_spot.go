package parking_lot

type ParkingSpot struct {
	ID            int
	IsOccupied    bool
	VehicleType   VehicleType
	LicenseNumber string
}
