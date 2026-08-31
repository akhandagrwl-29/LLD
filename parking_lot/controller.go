package parking_lot

func Run() {
	parkingLot := NewParkingLot(1)
	parkingLot.AddLevel(NewLevel(1, 10))

	parkingLot.ShowAvailability()

	car := NewCar("car_123")
	truck := NewTruck("truck_123")
	bike := NewMotorCycle("bike_123")
	parkingLot.ParkVehicle(car)
	parkingLot.ParkVehicle(truck)
	parkingLot.ParkVehicle(bike)

	parkingLot.ShowAvailability()

}
