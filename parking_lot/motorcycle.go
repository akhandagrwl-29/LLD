package parking_lot

func NewMotorCycle(licenseNumber string) *BaseVehicle {
	return &BaseVehicle{
		LicenseNumber: licenseNumber,
		VehicleType:   MOTORCYCLE,
	}
}
