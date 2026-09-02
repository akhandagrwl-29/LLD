package restaurant_management_system

type MenuItem struct {
	id          string
	name        string
	description string
	isAvailable bool
}

func NewMenuItem(id string, name string, description string, isAvailable bool) *MenuItem {
	return &MenuItem{
		id:          id,
		name:        name,
		description: description,
		isAvailable: isAvailable,
	}
}
