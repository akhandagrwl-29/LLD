package fooddeliveryservice

type DeliveryAgent struct {
	id          string
	name        string
	phone       string
	isAvailable bool
}

func NewDeliveryAgent(
	id string,
	name string,
	phone string,
	isAvailable bool,
) *DeliveryAgent {
	return &DeliveryAgent{
		id:          id,
		name:        name,
		phone:       phone,
		isAvailable: isAvailable,
	}
}
