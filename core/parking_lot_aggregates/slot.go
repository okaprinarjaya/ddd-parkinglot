package core

import core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"

type Slot struct {
	code              string
	vehicle           core_value_objects.Vehicle
	PersistenceStatus string
}

func NewSlot(code string, veh core_value_objects.Vehicle) Slot {
	return Slot{code: code, vehicle: veh, PersistenceStatus: "none"}
}

func (s *Slot) Code() string {
	return s.code
}

func (s *Slot) Fill(v core_value_objects.Vehicle) {
	if v.PlateNumber != "" {
		s.vehicle = v
		s.PersistenceStatus = "update"
	}
}

func (s *Slot) UnFill() {
	s.vehicle.PlateNumber = ""
	s.vehicle.Type = ""
	s.vehicle.Color = ""
	s.PersistenceStatus = "update"
}

func (s *Slot) IsFilled() bool {
	return s.vehicle.PlateNumber != ""
}

func (s *Slot) GetVehicle() core_value_objects.Vehicle {
	return s.vehicle
}
