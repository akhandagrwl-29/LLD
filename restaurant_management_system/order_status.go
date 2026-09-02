package restaurant_management_system

type OrderStatus int

const (
	OrderStatusPreparing OrderStatus = iota
	OrderStatusPrepared
	OrderStatusOnTheWay
	OrderStatusDelivered
)
