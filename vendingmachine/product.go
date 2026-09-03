package vendingmachine

type Product struct {
	id    string
	name  string
	price int
}

func NewProduct(id string, name string, price int) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}
}
