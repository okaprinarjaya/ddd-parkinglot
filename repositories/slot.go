package repositories

import "github.com/okaprinarjaya/parking-lot/models"

type ISlotRepository interface {
	Create() *models.Slot
	GetAllSlots() []*models.Slot
}
