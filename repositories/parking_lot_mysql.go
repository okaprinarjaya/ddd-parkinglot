package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	core "github.com/okaprinarjaya/parking-lot/core/parking_lot_aggregates"
	core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"
)

type VehicleDTO struct {
	ID         string  `db:"id"`
	VehicleID  *string `db:"vehicle_id"`
	SlotStatus string  `db:"slot_status"`
}

type ParkingLotRepositoryMysql struct {
	*sqlx.DB
}

func NewParkingLotRepository(db *sqlx.DB) *ParkingLotRepositoryMysql {
	return &ParkingLotRepositoryMysql{DB: db}
}

func (plr *ParkingLotRepositoryMysql) FindOne(parking_lot_id string) (core.ParkingLot, error) {
	var vhc_dto_list []VehicleDTO
	qry := "SELECT id, vehicle_id, slot_status FROM slots"
	err := plr.Select(&vhc_dto_list, qry)

	if err != nil {
		fmt.Println(err.Error())
		return core.ParkingLot{}, err
	}

	var slots []core.Slot

	for _, v_dto := range vhc_dto_list {
		if v_dto.VehicleID != nil {
			veh := &core_value_objects.Vehicle{PlateNumber: *v_dto.VehicleID, Type: "TRUCK", Color: "YELLOW"}
			slots = append(slots, core.NewSlot(v_dto.ID, veh, v_dto.SlotStatus, "none"))
		} else {
			slots = append(slots, core.NewSlot(v_dto.ID, nil, v_dto.SlotStatus, "none"))
		}
	}

	return core.NewParkingLot(parking_lot_id, slots), nil
}

func (plr *ParkingLotRepositoryMysql) Update(pl *core.ParkingLot) error {
	for i := 0; i < len(pl.Slots); i++ {
		slot := &pl.Slots[i]
		if pl.Slots[i].PersistenceStatus == "update" {
			if slot.GetVehicle() != nil {
				plr.MustExec("UPDATE slots SET vehicle_id = ?, slot_status = ? WHERE id = ?", slot.GetVehicle().PlateNumber, slot.SlotStatus(), slot.Code())
			} else {
				plr.MustExec("UPDATE slots SET vehicle_id = NULL, slot_status = ? WHERE id = ?", slot.SlotStatus(), slot.Code())
			}

			slot.PersistenceStatus = "none"
		}
	}

	return nil
}
