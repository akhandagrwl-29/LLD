package parking_lot

type VehicleType string

var CAR VehicleType
var TRUCK VehicleType
var MOTORCYCLE VehicleType

type BaseVehicle struct {
	LicenseNumber string
	VehicleType   VehicleType
}
type Vehicle interface {
	GetLicenseNumber() string
	getVehicleType() VehicleType
}

func (b *BaseVehicle) GetLicenseNumber() string {
	return b.LicenseNumber
}

func (b *BaseVehicle) getVehicleType() VehicleType {
	return b.VehicleType
}
