package vendingmachine

func Run() {
	vendingMachine := NewVendingMachine()
	product1 := NewProduct("1", "Coke", 20)
	product2 := NewProduct("2", "Pepsi", 30)
	product3 := NewProduct("2", "Chips", 10)

	vendingMachine.AddInventory(map[*Product]int{
		product1: 10,
		product2: 10,
		product3: 10,
	})

	vendingMachine.SelectProduct(product1)
	vendingMachine.InsertNote(TenRupeeNote)
	vendingMachine.InsertCoin(TenRupeeCoin)
	vendingMachine.DispatchProduct()
	vendingMachine.ReturnChange()

	vendingMachine.SelectProduct(product2)
	vendingMachine.InsertCoin(TenRupeeCoin)
	vendingMachine.InsertCoin(TwoRupeeCoin)
	vendingMachine.DispatchProduct()
	vendingMachine.ReturnChange()
}
