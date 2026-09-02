package restaurant_management_system

import "time"

type Order struct {
	id          string
	items       []*MenuItem
	status      OrderStatus
	totalAmount float64
	timestamp   time.Time
}
