package services

import "github.com/okaprinarjaya/parking-lot/models"

type IParkingLotService interface {
	InitializeSlots()
	AddNewSlot(code string) error
	GetSlots() []*models.Slot
	ReserveSlot() *models.Slot
}
