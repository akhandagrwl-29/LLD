package parkinglot

func NewTruck(licenseNunber string) Vehicle {
	return &BaseVehicle{LicensePlate: licenseNunber, Type: TRUCK}
}
