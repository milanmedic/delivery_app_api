package deliverer_sql_db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type DelivererDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateDelivererDb(dbDriver *dbdrivers.DeliveryAppDb) *DelivererDb {
	return &DelivererDb{dbDriver: dbDriver}
}

func (dDb *DelivererDb) AddDeliverer(d models.Deliverer) error {
	tx, err := dDb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO deliverer(id, username, name, surname, email, password, date_of_birth, address, role, delivery_in_progress, verification_status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Id, d.Username, d.Name, d.Surname, d.Email, d.Password, d.DateOfBirth, d.Address.Id, d.Role, d.DeliveryInProgress, d.VerificationStatus)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (dDb *DelivererDb) UpdateProperty(property string, value interface{}, id string) error {
	stmt, err := dDb.dbDriver.Prepare(fmt.Sprintf(`UPDATE deliverer SET %s = ? where deliverer.id = ?;`, property))
	if err != nil {
		return err
	}
	defer stmt.Close()

	val := reflect.ValueOf(value)
	ptr := val

	switch ptr.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		concreteValue := strconv.FormatInt(ptr.Int(), 10)
		_, err = stmt.Exec(concreteValue, id)
	case reflect.String:
		concreteValue := ptr.String()
		_, err = stmt.Exec(concreteValue, id)
	case reflect.Float32, reflect.Float64:
		concreteValue := strconv.FormatFloat(ptr.Float(), 'f', 2, 32)
		_, err = stmt.Exec(concreteValue, id)
	case reflect.Bool:
		concreteValue := strconv.FormatBool(ptr.Bool())
		_, err = stmt.Exec(concreteValue, id)
	}

	if err != nil {
		return err
	}

	return nil
}

func (dDb *DelivererDb) GetBy(attr string, value interface{}) (*models.Deliverer, error) {
	stmt, err := dDb.dbDriver.Prepare(fmt.Sprintf(` SELECT d.id, d.username, d.name, d.surname, d.email, d.password, d.date_of_birth,
	d.role, d.verification_status, d.delivery_in_progress, a.city, a.street, a.street_num, a.postfix, a.id from deliverer d inner join address a on a.id = d.address WHERE d.%s = ?;`, attr))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var row *sql.Row

	val := reflect.ValueOf(value)
	ptr := val

	switch ptr.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		concreteValue := strconv.FormatInt(ptr.Int(), 10)
		row = stmt.QueryRow(concreteValue)
	case reflect.String:
		concreteValue := ptr.String()
		row = stmt.QueryRow(concreteValue)
	case reflect.Float32, reflect.Float64:
		concreteValue := strconv.FormatFloat(ptr.Float(), 'f', 2, 32)
		row = stmt.QueryRow(concreteValue)
	case reflect.Bool:
		concreteValue := strconv.FormatBool(ptr.Bool())
		row = stmt.QueryRow(concreteValue)
	}

	var deliverer *models.Deliverer = models.CreateDeliverer()
	var id string
	var name string
	var username string
	var surname string
	var email string
	var password string
	var dateOfBirth string
	var addrId int
	var role string
	var status string
	var deliveryInProgress bool
	var city string
	var street string
	var streetNum int
	var postfix string
	err = row.Scan(&id, &username, &name, &surname, &email, &password, &dateOfBirth, &role, &status, &deliveryInProgress, &city, &street, &streetNum, &postfix, &addrId)
	if err != nil {
		return nil, nil
	}

	var addr *models.Address = models.CreateAddress(addrId, streetNum, city, street, postfix)

	deliverer.SetId(id)
	deliverer.SetName(name)
	deliverer.SetSurname(surname)
	deliverer.SetUsername(username)
	deliverer.SetEmail(email)
	deliverer.SetPassword(password)
	deliverer.SetDateOfBirth(dateOfBirth)
	deliverer.SetRole(role)
	deliverer.SetVerificationStatus(status)
	deliverer.SetDeliveryProgress(deliveryInProgress)
	deliverer.SetAddress(addr)

	return deliverer, nil
}
