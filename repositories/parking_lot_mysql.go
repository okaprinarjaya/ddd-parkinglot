package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	core "github.com/okaprinarjaya/parking-lot/core/parking_lot_aggregates"
	core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"
)

type VehicleDTO struct {
	ID        string  `db:"id"`
	VehicleID *string `db:"vehicle_id"`
}

type ParkingLotRepositoryMysql struct {
	*sqlx.DB
}

func NewParkingLotRepository(db *sqlx.DB) *ParkingLotRepositoryMysql {
	return &ParkingLotRepositoryMysql{DB: db}
}

func (plr *ParkingLotRepositoryMysql) FindOne(parking_lot_id string) (core.ParkingLot, error) {
	var vhc_dto_list []VehicleDTO
	qry := "SELECT id, vehicle_id FROM slots"
	err := plr.Select(&vhc_dto_list, qry)

	if err != nil {
		fmt.Println(err.Error())
		return core.ParkingLot{}, err
	}

	var slots []core.Slot

	for _, v := range vhc_dto_list {
		if v.VehicleID != nil {
			slots = append(slots, core.NewSlot(v.ID, core_value_objects.Vehicle{PlateNumber: *v.VehicleID, Type: "TRUCK", Color: "YELLOW"}))
		} else {
			slots = append(slots, core.NewSlot(v.ID, core_value_objects.Vehicle{}))
		}
	}

	return core.NewParkingLot(parking_lot_id, slots), nil
}

func (plr *ParkingLotRepositoryMysql) Update(pl *core.ParkingLot) error {
	for i := 0; i < len(pl.Slots); i++ {
		slot := &pl.Slots[i]
		if pl.Slots[i].PersistenceStatus == "update" {
			if slot.GetVehicle().PlateNumber != "" {
				plr.MustExec("UPDATE slots SET vehicle_id = ? WHERE id = ?", slot.GetVehicle().PlateNumber, slot.Code())
			} else {
				plr.MustExec("UPDATE slots SET vehicle_id = NULL WHERE id = ?", slot.Code())
			}

			slot.PersistenceStatus = "none"
		}
	}

	return nil
}
