package vendingmachine

type VendingMachine struct {
	IdleState       VendingMachineState
	ReadyState      VendingMachineState
	DispatchState   VendingMachineState
	ReturnExchange  VendingMachineState
	CurrentState    VendingMachineState
	SelectedProduct *Product
	Inventory       *Inventory
	TotalMoney      int
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		Inventory: NewInventory(),
	}
	vm.IdleState = &IdleState{
		vm,
	}
	vm.ReadyState = &ReadyState{vm}
	vm.DispatchState = &DispatchState{vm}
	vm.ReturnExchange = &ExchangeState{vm}
	vm.CurrentState = vm.IdleState
	return vm
}

func (vm *VendingMachine) SelectProduct(product *Product) {
	vm.CurrentState.SelectProduct(product)
}

func (vm *VendingMachine) InsertNote(note Note) {
	vm.CurrentState.InsertNote(note)
}

func (vm *VendingMachine) InsertCoin(coin Coin) {
	vm.CurrentState.InsertCoin(coin)
}

func (vm *VendingMachine) DispatchProduct() {
	vm.CurrentState.DispatchProduct()
}

func (vm *VendingMachine) ReturnChange() {
	vm.CurrentState.ReturnExchange()
}

func (vm *VendingMachine) AddInventory(inventory map[*Product]int) {
	for item, count := range inventory {
		vm.Inventory.products[item] += count
	}
}
