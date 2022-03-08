package core

import core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"

type Slot struct {
	code              string
	vehicle           *core_value_objects.Vehicle
	slotStatus        string
	PersistenceStatus string
}

func NewSlot(code string, veh *core_value_objects.Vehicle, slotStatus string, ps string) Slot {
	return Slot{code: code, vehicle: veh, slotStatus: slotStatus, PersistenceStatus: ps}
}

func (s *Slot) Code() string {
	return s.code
}

func (s *Slot) SlotStatus() string {
	return s.slotStatus
}

func (s *Slot) Fill(veh *core_value_objects.Vehicle) {
	if veh != nil {
		s.vehicle = veh
		s.slotStatus = "RESERVED"
		s.PersistenceStatus = "update"
	}
}

func (s *Slot) CheckIN() {
	if s.vehicle != nil {
		s.slotStatus = "CHECK-IN"
		s.PersistenceStatus = "update"
	}
}

func (s *Slot) UnFill() {
	s.vehicle = nil
	s.slotStatus = "AVAILABLE"
	s.PersistenceStatus = "update"
}

func (s *Slot) IsFilled() bool {
	return s.vehicle != nil
}

func (s *Slot) GetVehicle() *core_value_objects.Vehicle {
	return s.vehicle
}
