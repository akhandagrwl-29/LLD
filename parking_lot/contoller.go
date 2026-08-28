package parkinglot

func Run() {
	parkingLot := NewParkingLot()
	parkingLot.ShowAvailability()

	parkingLot.AddLevel(NewLevel(2, 10))
	parkingLot.ShowAvailability()

	car := NewCar("12334")
	motorBike := NewMotorCycle("12345")
	truck := NewTruck("12673q32")

	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(truck)
	parkingLot.ParkVehicle(motorBike)

	parkingLot.ShowAvailability()

}
