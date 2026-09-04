package carrentalsystem

import (
	"fmt"
	models2 "github.com/LLD/carrentalsystem/models"
)

type CarRentalSystem struct {
	users         []*models2.User
	stores        map[string]*models2.Store
	cars          map[string]*models2.Car
	booking       map[string]*models2.Booking
	totalBookings int
}

var carRentalSystem *CarRentalSystem

func NewCarRentalSystem() *CarRentalSystem {
	if carRentalSystem != nil {
		return carRentalSystem
	}
	return &CarRentalSystem{
		booking: make(map[string]*models2.Booking),
		users:   make([]*models2.User, 0),
		stores:  make(map[string]*models2.Store),
		cars:    make(map[string]*models2.Car),
	}
}

func (carRentalSystem *CarRentalSystem) AddUser(user *models2.User) {
	carRentalSystem.users = append(carRentalSystem.users, user)
}

func (carRentalSystem *CarRentalSystem) AddStore(store *models2.Store) {
	carRentalSystem.stores[store.Name] = store
}

func (carRentalSystem *CarRentalSystem) AddCar(car *models2.Car) {
	carRentalSystem.cars[car.Id] = car
}

func (carRentalSystem *CarRentalSystem) ListCars(storeId string) []*models2.Car {
	store, exists := carRentalSystem.stores[storeId]
	if !exists {
		fmt.Printf("carRentalSystem.ListCars store %s does not exist\n", storeId)
	}
	return store.Cars
}

func (carRentalSystem *CarRentalSystem) RequestCar(user *models2.User, carID string, slot *models2.Slot) string {
	car, exists := carRentalSystem.cars[carID]
	if !exists {
		fmt.Printf("carRentalSystem.RequestCar car %s does not exist\n", carID)
	}
	if !car.BookCar(slot) {
		fmt.Printf("carRentalSystem.RequestCar car %s is not available\n", carID)
	}

	booking := models2.NewBooking(user, car, slot, car.Price)
	carRentalSystem.booking[booking.Id] = booking
	return booking.Id
}

func (carRentalSystem *CarRentalSystem) ConfirmBooking(bookingID string) {
	booking, exists := carRentalSystem.booking[bookingID]
	if !exists {
		fmt.Printf("carRentalSystem.ConfirmBooking booking %s does not exist\n", bookingID)
	}
	booking.Status = models2.BookingStatusConfirmed // after payment is completed
}

func (carRentalSystem *CarRentalSystem) CancelBooking(bookingID string) {
	booking, exists := carRentalSystem.booking[bookingID]
	if !exists {
		fmt.Printf("carRentalSystem.CancelBooking booking %s does not exist\n", bookingID)
	}

	if booking.Status == models2.BookingStatusCancelled {
		fmt.Printf("carRentalSystem.CancelBooking booking %s is already cancelled\n", bookingID)
	}

	booking.Status = models2.BookingStatusCancelled
	car, exists := carRentalSystem.cars[booking.Car.Id]
	if !exists {
		fmt.Printf("carRentalSystem.CancelBooking car %s does not exist\n", booking.Car.Id)
	}

	for _, slot := range car.Slots {
		if slot.Id == booking.Slot.Id {
			slot.IsOccupied = false
		}
	}
}
