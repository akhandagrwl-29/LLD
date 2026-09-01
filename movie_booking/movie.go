package movie_booking

type Movie struct {
	name        string
	description string
	duration    int
}

func NewMovie(name string, description string, duration int) *Movie {
	return &Movie{
		name:        name,
		description: description,
		duration:    duration,
	}
}
