package restaurant_management_system

type Staff struct {
	id      string
	name    string
	contact string
	role    string
}

func NewStaff(id string, name string, contact string, role string) *Staff {
	return &Staff{
		id:      id,
		name:    name,
		contact: contact,
		role:    role,
	}
}
