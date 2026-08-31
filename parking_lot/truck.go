package parking_lot

func NewTruck(licenseNumber string) *BaseVehicle {
	return &BaseVehicle{
		LicenseNumber: licenseNumber,
		VehicleType:   TRUCK,
	}
}
