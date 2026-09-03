package vendingmachine

import "fmt"

type VendingMachineState interface {
	SelectProduct(product *Product)
	InsertNote(note Note)
	InsertCoin(coin Coin)
	DispatchProduct()
	ReturnExchange()
}

type IdleState struct {
	vendingMachine *VendingMachine
}

func NewIdleState(vendingMachine *VendingMachine) *IdleState {
	return &IdleState{vendingMachine: vendingMachine}
}

func (s *IdleState) SelectProduct(product *Product) {
	if s.vendingMachine.Inventory.products[product] == 0 {
		fmt.Printf("Product %v has no inventory\n", product)
		return
	}
	s.vendingMachine.SelectedProduct = product
	fmt.Printf("Selected Product %v\n", product)
	s.vendingMachine.CurrentState = s.vendingMachine.ReadyState
	s.vendingMachine.Inventory.products[product] -= 1
}

func (s *IdleState) InsertNote(note Note) {
	fmt.Printf("First select a product")
}

func (s *IdleState) InsertCoin(coin Coin) {
	fmt.Printf("First select a product")
}

func (s *IdleState) ReturnExchange() {
	fmt.Printf("First select a product")
}

func (s *IdleState) DispatchProduct() {
	fmt.Printf("First select a product")
}

type ReadyState struct {
	vendingMachine *VendingMachine
}

func (s *ReadyState) SelectProduct(product *Product) {
	fmt.Printf("Product already selected\n")
}

func (s *ReadyState) InsertNote(note Note) {
	fmt.Printf("Note Inserted with value: %d\n", note)
	s.vendingMachine.TotalMoney += int(note)
	s.DispatchProduct()
	//s.vendingMachine.CurrentState = s.vendingMachine.DispatchState
}

func (s *ReadyState) InsertCoin(coin Coin) {
	fmt.Printf("Coin Inserted with value: %d\n", coin)
	s.vendingMachine.TotalMoney += int(coin)
	s.DispatchProduct()
	//s.vendingMachine.CurrentState = s.vendingMachine.DispatchState
}

func (s *ReadyState) ReturnExchange() {
	fmt.Printf("First insert a note or coin")
}

func (s *ReadyState) DispatchProduct() {
	change := s.vendingMachine.TotalMoney - s.vendingMachine.SelectedProduct.price
	if change >= 0 {
		fmt.Printf("Payment collected, preparing for dispatch\n")
		s.vendingMachine.CurrentState = s.vendingMachine.DispatchState
	}
}

type DispatchState struct {
	vendingMachine *VendingMachine
}

func (s *DispatchState) SelectProduct(product *Product) {
	fmt.Printf("Product already selected\n")
}

func (s *DispatchState) InsertNote(note Note) {
	fmt.Printf("Note already inserted\n")
}

func (s *DispatchState) InsertCoin(coin Coin) {
	fmt.Printf("Coin already inserted\n")
}

func (s *DispatchState) ReturnExchange() {
	fmt.Printf("First dispatch a product\n")
}

func (s *DispatchState) DispatchProduct() {
	fmt.Printf("dispatching a product with name: %s\n", s.vendingMachine.SelectedProduct.name)
	s.vendingMachine.CurrentState = s.vendingMachine.ReturnExchange
}

type ExchangeState struct {
	vendingMachine *VendingMachine
}

func (s *ExchangeState) SelectProduct(product *Product) {
	fmt.Printf("Product already selected\n")
}

func (s *ExchangeState) InsertNote(note Note) {
	fmt.Printf("Note already inserted\n")
}

func (s *ExchangeState) InsertCoin(coin Coin) {
	fmt.Printf("Coin already inserted\n")
}

func (s *ExchangeState) ReturnExchange() {
	fmt.Printf("Please collect the returned Money")
	change := s.vendingMachine.TotalMoney - s.vendingMachine.SelectedProduct.price
	if change > 0 {
		fmt.Printf("Returing money :%d\n", change)
	} else {
		fmt.Printf("No return\n")
	}
	s.vendingMachine.CurrentState = s.vendingMachine.IdleState
	s.vendingMachine.SelectedProduct = nil
	s.vendingMachine.TotalMoney = 0
}

func (s *ExchangeState) DispatchProduct() {
	fmt.Printf("Product already dispatched\n")
}
