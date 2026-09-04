package carrentalsystem

import (
	"fmt"
	models2 "github.com/LLD/carrentalsystem/models"
	"time"
)

func Run() {
	carRentalSystem := NewCarRentalSystem()

	user1 := models2.NewUser("user1", "akhand", "akhand@gmail.com")
	carRentalSystem.AddUser(user1)

	car1 := models2.NewCar("Car1", "Toyata", "SUV", 500)
	carRentalSystem.AddCar(car1)

	car2 := models2.NewCar("Car2", "Maruti", "Mini", 200)
	carRentalSystem.AddCar(car2)

	store1 := models2.NewStore("store1", "Bengaluru")
	store1.AddCar(car1)
	carRentalSystem.AddStore(store1)

	listedCars := carRentalSystem.ListCars(store1.Name)
	for _, car := range listedCars {
		fmt.Println(car.Name)
	}

	bookingID := carRentalSystem.RequestCar(user1, car1.Id, models2.NewSlot(
		"slot1",
		time.Date(2026, time.September, 1, 15, 15, 0, 0, time.UTC),
		time.Date(2026, time.September, 1, 16, 0, 0, 0, time.UTC),
	))

	_ = carRentalSystem.RequestCar(user1, car1.Id, models2.NewSlot(
		"slot1",
		time.Date(2026, time.September, 1, 14, 15, 0, 0, time.UTC),
		time.Date(2026, time.September, 1, 17, 30, 0, 0, time.UTC),
	))

	for _, booking := range carRentalSystem.booking {
		fmt.Printf("Booking ID: %s of car name: %s and booking status: %s\n", booking.Id, booking.Car.Name, booking.Status)
	}

	carRentalSystem.ConfirmBooking(bookingID)

	for _, booking := range carRentalSystem.booking {
		fmt.Printf("Booking ID: %s of car name: %s and booking status: %s\n", booking.Id, booking.Car.Name, booking.Status)
	}

	fmt.Println(bookingID)

}
