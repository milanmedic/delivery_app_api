package usersqldb

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type UserDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateUserDb(dbDriver *dbdrivers.DeliveryAppDb) *UserDb {
	return &UserDb{dbDriver: dbDriver}
}

func (udb *UserDb) GetBy(attr string, value interface{}) error {
	return nil
}

func (udb *UserDb) AddOne(u models.User) error {
	tx, err := udb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO customer(id, username, name, surname, email, password, date_of_birth, address, role, verification_status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id, u.Username, u.Name, u.Surname, u.Email, u.Password, u.DateOfBirth, u.Address.Id, u.Role, u.VerificationStatus)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (udb *UserDb) Update() error {
	return nil
}

func (udb *UserDb) Delete() error {
	return nil
}
