package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeSlots(t *testing.T) {
	var parkingLotService IParkingLotService = NewParkingLotService()
	parkingLotService.InitializeSlots()

	assert.Equal(t, 7, len(parkingLotService.GetSlots()))
}

func TestReserveSlot(t *testing.T) {
	var parkingLotRepo IParkingLotService = NewParkingLotService()
	parkingLotRepo.InitializeSlots()

	slot := parkingLotRepo.ReserveSlot()

	slots := parkingLotRepo.GetSlots()
	slotOrigin := slots[0]

	assert.Equal(t, slot.IsFilled(), slotOrigin.IsFilled())
	assert.Equal(t, "B2872ON", slot.GetVehicle().PlateNumber)
}
