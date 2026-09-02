package restaurant_management_system

import "time"

type Reservation struct {
	id            string
	customerName  string
	startTime     time.Time
	endTime       time.Time
	customerEmail string
	partySize     int
}

func NewReservation(
	id string,
	customerName string,
	startTime time.Time,
	endTime time.Time,
	customerEmail string,
	partySize int) *Reservation {
	return &Reservation{
		id:            id,
		customerName:  customerName,
		startTime:     startTime,
		endTime:       endTime,
		customerEmail: customerEmail,
		partySize:     partySize,
	}
}
