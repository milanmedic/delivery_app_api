package main

import (
	deliveryAppDb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

func main() {
	db, err := deliveryAppDb.CreateDeliveryAppDb()
	HandleError(err)

	err = db.CheckConnection()
	HandleError(err)

	err = db.CloseConnection()
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
