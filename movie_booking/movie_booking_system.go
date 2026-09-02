package movie_booking

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var instance *MovieBookingSystem

type MovieBookingSystem struct {
	theatres      []*Theatre
	movies        []*Movie
	shows         map[int64]*Show
	bookings      map[string]*Booking
	bookingsCount int64
	mu            sync.RWMutex
}

func NewMovieBookingSystem() *MovieBookingSystem {
	if instance != nil {
		return instance
	}
	return &MovieBookingSystem{
		theatres:      make([]*Theatre, 0),
		movies:        make([]*Movie, 0),
		shows:         make(map[int64]*Show),
		bookings:      make(map[string]*Booking),
		bookingsCount: 0,
	}
}

func (sys *MovieBookingSystem) AddTheatre(theatre *Theatre) {
	sys.theatres = append(sys.theatres, theatre)
}

func (sys *MovieBookingSystem) AddMovie(movie *Movie) {
	sys.movies = append(sys.movies, movie)
}

func (sys *MovieBookingSystem) AddShow(show *Show) {
	sys.shows[show.ShowID()] = show
}

func (sys *MovieBookingSystem) BookTicket(user *User, show *Show, seats []*Seat) (*Booking, error) {
	for _, seat := range seats {
		showSeats, exists := show.seatMap[seat.id]
		if !exists || showSeats.GetStatus() == SeatStatusBooked {
			fmt.Printf("seat is not avaialble for booking\n")
			return nil, errors.New("seat not available")
		}
	}

	for _, seat := range seats {
		show.seatMap[seat.id].seatStatus = SeatStatusBooked
	}
	var totalprice int32

	for _, seat := range seats {
		totalprice += seat.price
	}

	bookingID := sys.generateBookingID()
	booking := NewBooking(bookingID, user, show, totalprice, seats, BookingStatusPending)
	sys.bookings[bookingID] = booking
	return booking, nil
}

func (sys *MovieBookingSystem) generateBookingID() string {
	id := atomic.AddInt64(&sys.bookingsCount, 1)
	return fmt.Sprintf("booking:%d", id)
}

func (sys *MovieBookingSystem) ConfirmBooking(id string) error {
	sys.mu.Lock()
	defer sys.mu.Unlock()
	for _, booking := range sys.bookings {
		if booking.getBookingID() == id {
			booking.bookingStatus = BookingStatusConfirmed
			return nil
		}
	}
	return errors.New("booking not found")
}

func (sys *MovieBookingSystem) CancelBooking(id string) error {
	for _, booking := range sys.bookings {
		if booking.getBookingID() == id && booking.bookingStatus != BookingStatusCancelled {
			booking.bookingStatus = BookingStatusCancelled
			for _, seat := range booking.seats {
				seat.SetStatus(SeatStatusAvailable)
			}
			return nil
		}
	}
	return errors.New("booking not found")
}

func GenerateSeats(row, col int) map[string]*Seat {
	seatMap := make(map[string]*Seat)
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			var price int32
			if i <= 2 || j <= 2 {
				price = 200
			} else {
				price = 100
			}
			seatMap[fmt.Sprintf("%d:%d", i, j)] = &Seat{
				id:         fmt.Sprintf("%d:%d", i, j),
				seatStatus: SeatStatusAvailable,
				price:      price,
				row:        i,
				col:        j,
			}

		}
	}
	return seatMap
}
