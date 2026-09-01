package movie_booking

type Seat struct {
	id         string
	seatStatus SeatStatus
	price      int32
	row        int
	col        int
}

func (s *Seat) GetStatus() SeatStatus {
	return s.seatStatus
}

func (s *Seat) SetStatus(status SeatStatus) {
	s.seatStatus = status
}
