package order_sql_db

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type OrderDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateOrderDb(dbDriver *dbdrivers.DeliveryAppDb) *OrderDb {
	return &OrderDb{dbDriver: dbDriver}
}

func (odb *OrderDb) CreateOrder(odto dto.OrderInputDto) error {
	tx, err := odb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO customer_order(comment, address, basket, customer) VALUES(?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(odto.Comment, odto.Address.Id, odto.Basket.Id, odto.CustomerID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
