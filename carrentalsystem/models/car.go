package models

type Car struct {
	Id          string
	Name        string
	VehicleType string
	Slots       []*Slot
	Price       int
}

func NewCar(id, name string, vehicleType string, price int) *Car {
	return &Car{
		Id:          id,
		Name:        name,
		VehicleType: vehicleType,
		Price:       price,
	}
}

func (c *Car) BookCar(reqSlot *Slot) bool {
	for _, slot := range c.Slots {
		if canBook(slot, reqSlot) {
			reqSlot.IsOccupied = true
			c.Slots = append(c.Slots, reqSlot)
			return true
		}
	}
	return false
}

func canBook(slot *Slot, reqSlot *Slot) bool {
	// logic to be checked
	//req : [10, 12]
	// [8 11]
	if !slot.IsOccupied {
		return true
	}
	if reqSlot.StartTime.Before(slot.EndTime) && reqSlot.StartTime.After(slot.StartTime) {
		return false
	}
	if reqSlot.EndTime.Before(slot.EndTime) && reqSlot.EndTime.After(slot.StartTime) {
		return false
	}
	return true
}
