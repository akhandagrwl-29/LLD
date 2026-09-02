package restaurant_management_system

type Restaurant struct {
	name         string
	menuItems    []*MenuItem
	orders       map[string]*Order
	reservations map[int]*Reservation
	payment      map[int]*Payment
	staff        []*Staff
}

var restaurantSystem *Restaurant

func NewRestaurantSystem() *Restaurant {
	if restaurantSystem == nil {
		restaurantSystem = &Restaurant{
			name:         "Restaurant",
			menuItems:    []*MenuItem{},
			orders:       map[string]*Order{},
			reservations: map[int]*Reservation{},
			payment:      map[int]*Payment{},
			staff:        []*Staff{},
		}
	}
	return restaurantSystem
}

func (sys *Restaurant) AddItems(menuItem *MenuItem) {
	sys.menuItems = append(sys.menuItems, menuItem)
}

func (sys *Restaurant) AddOrder(order *Order) {
	sys.orders[order.id] = order
}
