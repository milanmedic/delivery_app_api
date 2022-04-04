package usersqldb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type CustomerDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateCustomerDb(dbDriver *dbdrivers.DeliveryAppDb) *CustomerDb {
	return &CustomerDb{dbDriver: dbDriver}
}

func getUnderlyingAsValue(data interface{}) reflect.Value {
	return reflect.ValueOf(data)
}

func (cdb *CustomerDb) GetBy(attr string, value interface{}) (*models.Customer, error) {
	stmt, err := cdb.dbDriver.Prepare(fmt.Sprintf(` SELECT c.id, c.username, c.name, c.surname, c.email, c.password, c.date_of_birth,
	c.role, c.verification_status, a.city, a.street, a.street_num, a.postfix, a.id from customer c inner join address a on a.id = c.address WHERE %s = ?;`, attr))
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

	var customer *models.Customer = models.CreateCustomer()
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
	var city string
	var street string
	var streetNum int
	var postfix string
	err = row.Scan(&id, &username, &name, &surname, &email, &password, &dateOfBirth, &role, &status, &city, &street, &streetNum, &postfix, &addrId)
	if err != nil {
		return nil, nil
	}

	var addr *models.Address = models.CreateAddress(addrId, streetNum, city, street, postfix)

	customer.SetId(id)
	customer.SetName(name)
	customer.SetSurname(surname)
	customer.SetUsername(username)
	customer.SetEmail(email)
	customer.SetPassword(password)
	customer.SetDateOfBirth(dateOfBirth)
	customer.SetRole(role)
	customer.SetVerificationStatus(status)

	customer.SetAddress(addr)

	return customer, nil
}

func (cdb *CustomerDb) AddOne(u models.Customer) error {
	tx, err := cdb.dbDriver.Begin()
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

func (cdb *CustomerDb) Update() error {
	return nil
}

func (cdb *CustomerDb) Delete() error {
	return nil
}
