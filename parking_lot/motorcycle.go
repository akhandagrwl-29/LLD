package parkinglot

func NewMotorCycle(licenseNunber string) Vehicle {
	return &BaseVehicle{LicensePlate: licenseNunber, Type: MOTORCYCLE}
}
