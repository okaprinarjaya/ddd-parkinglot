package repositories

import "github.com/okaprinarjaya/parking-lot/models"

type SlotRepository struct {
	slots []*models.Slot
}

func (sp *SlotRepository) Create(code string) *models.Slot {
	slot := models.NewSlot(code)
	sp.slots = append(sp.slots, slot)
	return slot
}

func (sp *SlotRepository) GetAllSlots() []*models.Slot {
	return sp.slots
}
