package movie_booking

type BookingStatus int

const (
	BookingStatusPending BookingStatus = iota
	BookingStatusConfirmed
	BookingStatusCancelled
)

type SeatStatus int

const (
	SeatStatusBooked SeatStatus = iota
	SeatStatusAvailable
)
