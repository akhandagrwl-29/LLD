package movie_booking

type Theatre struct {
	id       int
	name     string
	location string
}

func NewTheatre(id int, name, location string) *Theatre {
	return &Theatre{
		id:       id,
		name:     name,
		location: location,
	}
}
