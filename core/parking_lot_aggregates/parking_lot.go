package core

import (
	"fmt"

	core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"
)

type ParkingLot struct {
	ID    string
	Slots []Slot
}

func NewParkingLot(id string, slots []Slot) ParkingLot {
	return ParkingLot{ID: id, Slots: slots}
}

func (pl *ParkingLot) ReserveSlot(veh *core_value_objects.Vehicle) (*Slot, error) {
	for i := 0; i < len(pl.Slots); i++ {
		if !pl.Slots[i].IsFilled() {
			pl.Slots[i].Fill(veh)
			return &pl.Slots[i], nil
		}
	}

	return nil, fmt.Errorf("no slot available")
}

func (pl *ParkingLot) CheckinSlot(plateNumber string) *Slot {
	for i := 0; i < len(pl.Slots); i++ {
		if pl.Slots[i].GetVehicle() != nil && pl.Slots[i].GetVehicle().PlateNumber == plateNumber {
			pl.Slots[i].CheckIN()
			return &pl.Slots[i]
		}
	}
	return nil
}

func (pl *ParkingLot) CheckoutSlot(plateNumber string) error {
	for i := 0; i < len(pl.Slots); i++ {
		if pl.Slots[i].GetVehicle().PlateNumber == plateNumber {
			pl.Slots[i].UnFill()
		}
	}

	return nil
}
