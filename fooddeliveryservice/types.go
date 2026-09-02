package fooddeliveryservice

type OrderStatus int

const (
	OrderPlaced OrderStatus = iota
	OrderConfirmed
	OrderPreparing
	OrderOutForDelivery
	OrderCancelled
	orderDelivered
)

func (s OrderStatus) String() string {
	switch s {
	case OrderPlaced:
		return "OrderPlaced"
	case OrderConfirmed:
		return "OrderConfirmed"
	case OrderPreparing:
		return "OrderPreparing"
	case OrderOutForDelivery:
		return "OrderOutForDelivery"
	case OrderCancelled:
		return "OrderCancelled"
	case orderDelivered:
	default:
		return "Unknown"
	}
	return "Unknown"
}
