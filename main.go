package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	// core_value_objects "github.com/okaprinarjaya/parking-lot/core/value_objects"
	"github.com/okaprinarjaya/parking-lot/repositories"
)

func createConnection() (*sqlx.DB, error) {
	dsnString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		"doadmin",
		"JNIjPYXDNnc240PQ",
		"db-mysql-sgp1-45875-development-do-user-7392084-0.b.db.ondigitalocean.com:25060",
		"oprex",
	)

	db, errConn := sqlx.Open("mysql", dsnString)
	if errConn != nil {
		return nil, errConn
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errConn
	}

	return db, nil
}

func main() {
	db, err := createConnection()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		parking_lot_repo := repositories.NewParkingLotRepository(db)
		parking_lot, _ := parking_lot_repo.FindOne("L01")

		fmt.Println(parking_lot)

		// vehicle := &core_value_objects.Vehicle{PlateNumber: "B1070ON", Type: "TRUCK", Color: "YELLOW"}
		// _, err := parking_lot.ReserveSlot(vehicle)

		// if err != nil {
		// 	fmt.Println(err.Error())
		// } else {
		// 	parking_lot_repo.Update(&parking_lot)
		// }

		//
		parking_lot.CheckinSlot("B1070ON")

		if err != nil {
			fmt.Println(err.Error())
		} else {
			parking_lot_repo.Update(&parking_lot)
		}

		// parking_lot.CheckoutSlot("B1070HR")
		// parking_lot_repo.Update(&parking_lot)

		fmt.Println(parking_lot)
	}
}
