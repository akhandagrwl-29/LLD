package models

type BookingStatus int

const (
	BookingStatusRequested BookingStatus = iota
	BookingStatusConfirmed
	BookingStatusCancelled
)
