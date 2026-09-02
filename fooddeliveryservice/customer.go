package fooddeliveryservice

type Customer struct {
	id    string
	name  string
	email string
}

func NewCustomer(id string, name string, email string) *Customer {
	return &Customer{
		id:    id,
		name:  name,
		email: email,
	}
}
