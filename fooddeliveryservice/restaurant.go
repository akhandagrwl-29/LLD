package fooddeliveryservice

type Restaurant struct {
	id       string
	name     string
	location string
	items    []*MenuItem
}

func NewRestaurant(id string, name string, location string) *Restaurant {
	return &Restaurant{
		id:       id,
		name:     name,
		location: location,
	}
}

func (r *Restaurant) AddItem(menuItem *MenuItem) {
	r.items = append(r.items, menuItem)
}

func (r *Restaurant) GetItems() []*MenuItem {
	return r.items
}
