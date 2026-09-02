package restaurant_management_system

type Payment struct {
	id            string
	amount        float64
	paymentStatus PaymentStatus
	paymentMethod PaymentMethod
}

func NewPayment(id string, amount float64, paymentStatus PaymentStatus, paymentMethod PaymentMethod) *Payment {
	return &Payment{
		id:            id,
		amount:        amount,
		paymentStatus: paymentStatus,
		paymentMethod: paymentMethod,
	}
}
