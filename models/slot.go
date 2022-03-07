package models

// Entities encapsulate Enterprise wide business rules.
// An entity can be an object with methods, or it can be a set of data structures and functions.
// It doesn’t matter so long as the entities could be used by many different applications in the enterprise.

type Slot struct {
	code    string
	vehicle *Vehicle
}

func NewSlot(code string) *Slot {
	return &Slot{code: code}
}

func (s *Slot) Fill(v *Vehicle) {
	if v != nil {
		s.vehicle = v
	}
}

func (s *Slot) UnFill() {
	s.vehicle = nil
}

func (s *Slot) IsFilled() bool {
	return s.vehicle != nil
}

func (s *Slot) GetVehicle() *Vehicle {
	return s.vehicle
}
