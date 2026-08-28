package parkinglot

func NewCar(licenseNunber string) Vehicle {
	return &BaseVehicle{LicensePlate: licenseNunber, Type: CAR}
}
