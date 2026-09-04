package models

type Store struct {
	Name     string
	Location string
	Cars     []*Car
}

func NewStore(name string, location string) *Store {
	return &Store{
		Name:     name,
		Location: location,
	}
}

func (s *Store) AddCar(car *Car) {
	s.Cars = append(s.Cars, car)
}
