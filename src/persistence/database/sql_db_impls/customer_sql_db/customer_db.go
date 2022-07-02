package customer_sql_db

import (
	"database/sql"
	"fmt"

	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type CustomerDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateCustomerDb(dbDriver *dbdrivers.DeliveryAppDb) *CustomerDb {
	return &CustomerDb{dbDriver: dbDriver}
}

func (cdb *CustomerDb) GetBy(attr string, value interface{}) (*models.Customer, error) {
	stmt, err := cdb.dbDriver.Prepare(fmt.Sprintf(` SELECT c.id, c.username, c.name, c.surname, c.email, c.password, c.date_of_birth,
	c.role, a.city, a.street, a.street_num, a.postfix, a.id from customer c inner join address a on a.id = c.address WHERE c.%s = ?;`, attr))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var row *sql.Row

	switch value.(type) {
	case int:
		row = stmt.QueryRow(value.(int))
	case float64:
		row = stmt.QueryRow(value.(float64))
	case bool:
		row = stmt.QueryRow(value.(bool))
	case string:
		row = stmt.QueryRow(value.(string))
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
	var city string
	var street string
	var streetNum int
	var postfix string
	err = row.Scan(&id, &username, &name, &surname, &email, &password, &dateOfBirth, &role, &city, &street, &streetNum, &postfix, &addrId)
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

	customer.SetAddress(addr)

	return customer, nil
}

func (cdb *CustomerDb) AddOne(u models.Customer) error {
	tx, err := cdb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO customer(id, username, name, surname, email, password, date_of_birth, address, role) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id, u.Username, u.Name, u.Surname, u.Email, u.Password, u.DateOfBirth, u.Address.Id, u.Role)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (cdb *CustomerDb) UpdateProperty(property string, value interface{}, id string) error {
	stmt, err := cdb.dbDriver.Prepare(fmt.Sprintf(`UPDATE customer SET %s = ? where customer.id = ?;`, property))
	if err != nil {
		return err
	}
	defer stmt.Close()

	switch value.(type) {
	case int:
		_, err = stmt.Exec(value.(int), id)
	case float64:
		_, err = stmt.Exec(value.(float64), id)
	case bool:
		_, err = stmt.Exec(value.(bool), id)
	case string:
		_, err = stmt.Exec(value.(string), id)
	}

	if err != nil {
		return err
	}

	return nil
}

func (cdb *CustomerDb) Update(c *models.Customer) (bool, error) {
	stmt, err := cdb.dbDriver.Prepare(`UPDATE customer SET
	name = ?,
	surname = ?,
	username = ?,
	email = ?,
	password = ?,
	date_of_birth = ?,
	address = ?
	where customer.id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.Name, c.Surname, c.Username, c.Email, c.Password, c.DateOfBirth, c.Address.Id, c.Id)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (cdb *CustomerDb) Delete() error {
	return nil
}
