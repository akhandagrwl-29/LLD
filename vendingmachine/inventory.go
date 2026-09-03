package vendingmachine

type Inventory struct {
	products map[*Product]int
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[*Product]int),
	}
}
