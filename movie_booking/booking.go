package movie_booking

type Booking struct {
	bookingID     string
	user          *User
	Show          *Show
	totalPrice    int32
	seats         []*Seat
	bookingStatus BookingStatus
}

func (b *Booking) getBookingID() string {
	return b.bookingID
}

func NewBooking(id string, user *User, show *Show, totalPrice int32, seats []*Seat, bookingStatus BookingStatus) *Booking {
	return &Booking{
		bookingID:     id,
		user:          user,
		Show:          show,
		totalPrice:    totalPrice,
		seats:         seats,
		bookingStatus: bookingStatus,
	}
}
