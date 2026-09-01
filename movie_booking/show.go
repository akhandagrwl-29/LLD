package movie_booking

import "time"

type Show struct {
	movie         *Movie
	theatre       *Theatre
	showId        int64
	showStartTime time.Time
	showEndTime   time.Time
	seatMap       map[string]*Seat
}

func NewShow(movie *Movie, startTime, endTime time.Time, seatMap map[string]*Seat, theatre *Theatre) *Show {
	return &Show{
		movie:         movie,
		theatre:       theatre,
		showId:        time.Now().Unix(),
		showStartTime: startTime,
		showEndTime:   endTime,
		seatMap:       seatMap,
	}
}

func (show *Show) ShowID() int64 {
	return show.showId
}
