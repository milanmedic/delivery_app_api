package usersqldb

import (
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type UserDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateUserDb(dbDriver *dbdrivers.DeliveryAppDb) *UserDb {
	return &UserDb{dbDriver: dbDriver}
}

func (udb *UserDb) GetUser(id string) error {
	return nil
}

func (udb *UserDb) AddOne() error {
	return nil
}

func (udb *UserDb) Update() error {
	return nil
}

func (udb *UserDb) Delete() error {
	return nil
}
