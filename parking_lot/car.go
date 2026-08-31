package parking_lot

func NewCar(licenseNumber string) *BaseVehicle {
	return &BaseVehicle{
		LicenseNumber: licenseNumber,
		VehicleType:   CAR,
	}
}
