package parkinglot

type VehicleType int

const (
	CAR        VehicleType = iota
	TRUCK                  = 1
	MOTORCYCLE             = 2
)

type Vehicle interface {
	GetLicensePlate() string
	GetVehicleType() VehicleType
}

type BaseVehicle struct {
	LicensePlate string
	Type         VehicleType
}

func NewVehicle(licensePlate string, vehicleType VehicleType) *BaseVehicle {
	return &BaseVehicle{
		LicensePlate: licensePlate,
		Type:         vehicleType,
	}
}

func (b *BaseVehicle) GetLicensePlate() string {
	return b.LicensePlate
}

func (b *BaseVehicle) GetVehicleType() VehicleType {
	return b.Type
}
