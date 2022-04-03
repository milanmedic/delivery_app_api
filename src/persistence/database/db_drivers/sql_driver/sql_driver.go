package sql_driver

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DeliveryAppDb struct {
	*sql.DB
}

func CreateDeliveryAppDb() (*DeliveryAppDb, error) {
	db, err := connectToDb()
	if err != nil {
		return nil, err
	}
	sqlDb := &DeliveryAppDb{db}
	return sqlDb, nil
}

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./delivery_app.db")
	if err != nil {
		return nil, err
	}
	return db, err
}

func (ddb *DeliveryAppDb) CloseConnection() error {
	err := ddb.Close()
	if err != nil {
		return err
	}
	return nil
}

func (ddb *DeliveryAppDb) CheckConnection() error {
	return ddb.Ping()
}
