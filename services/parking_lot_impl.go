package services

import (
	"github.com/okaprinarjaya/parking-lot/models"
	"github.com/okaprinarjaya/parking-lot/repositories"
)

type ParkingLotService struct {
	repositories.SlotRepository
}

func NewParkingLotService() *ParkingLotService {
	return &ParkingLotService{}
}

func (pl *ParkingLotService) InitializeSlots() {
	slots := []string{
		"B1001",
		"B1002",
		"B1003",
		"B1004",
		"B1005",
		"B1006",
		"B1007",
	}

	for i := 0; i < 7; i++ {
		pl.SlotRepository.Create(slots[i])
	}
}

func (pl *ParkingLotService) AddNewSlot(code string) error {
	pl.SlotRepository.Create(code)
	return nil
}

func (pl *ParkingLotService) GetSlots() []*models.Slot {
	return pl.GetAllSlots()
}

func (pl *ParkingLotService) ReserveSlot() *models.Slot {
	slots := pl.GetSlots()
	for i := 0; i < len(slots); i++ {
		if !slots[i].IsFilled() {
			slots[i].Fill(&models.Vehicle{PlateNumber: "B2872ON", Type: "TRUCK", Color: "YELLOW"})
			return slots[i]
		}
	}
	return nil
}
