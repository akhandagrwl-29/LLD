package movie_booking

import (
	"fmt"
	"time"
)

func Run() {
	bookingSystem := NewMovieBookingSystem()

	movie1 := NewMovie("movie1", "bollywood classics", 120)
	movie2 := NewMovie("movie2", "bollywood", 120)
	bookingSystem.AddMovie(movie1)
	bookingSystem.AddMovie(movie2)

	theatre1 := NewTheatre(1, "T1", "Marathalli")
	theatre2 := NewTheatre(2, "T2", "Kormangla")
	bookingSystem.AddTheatre(theatre1)
	bookingSystem.AddTheatre(theatre2)

	show1 := NewShow(movie1, time.Now(), time.Now().Add(time.Duration(movie1.duration)*time.Minute), GenerateSeats(10, 10), theatre1)
	show2 := NewShow(movie2, time.Now(), time.Now().Add(time.Duration(movie2.duration)*time.Minute), GenerateSeats(10, 10), theatre2)
	bookingSystem.AddShow(show1)
	bookingSystem.AddShow(show2)

	user := NewUser(1, "akhand", "akhand@gmail.com")
	booking, err := bookingSystem.BookTicket(user, show1, []*Seat{show1.seatMap["1:5"], show2.seatMap["1:6"]})
	if err != nil {
		fmt.Printf("Booking Ticket Error: %v\n", err)
	}
	fmt.Printf("Ticket booked: %s\n", booking.getBookingID())

	err = bookingSystem.ConfirmBooking(booking.getBookingID())
	if err != nil {
		fmt.Printf("Confirm Ticket Error: %v\n", err)
	}
	fmt.Printf("Ticket Confirmed: %s\n", booking.getBookingID())

	err = bookingSystem.CancelBooking(booking.getBookingID())
	if err != nil {
		fmt.Printf("Cancel Ticket Error: %v\n", err)
	}
	fmt.Printf("Ticket Canceled: %s\n", booking.getBookingID())
}
