package models

import (
	"fmt"
	"time"
)

type Booking struct {
	Id          string
	User        *User
	Car         *Car
	Slot        *Slot
	TotalAmount int
	Status      BookingStatus
}

func NewBooking(user *User, car *Car, slot *Slot, totalAmount int) *Booking {
	id := NewBookingID()
	return &Booking{
		Id:          id,
		User:        user,
		Car:         car,
		Slot:        slot,
		TotalAmount: totalAmount,
		Status:      BookingStatusRequested,
	}
}

func NewBookingID() string {
	return fmt.Sprintf("%s", time.Now().Format("20060102150405"))
}
