package fooddeliveryservice

type MenuItem struct {
	id          string
	name        string
	description string
	price       int
	isAvailable bool
}

func NewMenuItem(id string, name string, description string, isAvailable bool, price int) *MenuItem {
	return &MenuItem{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		isAvailable: isAvailable,
	}
}
