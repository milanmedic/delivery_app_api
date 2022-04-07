package admin_sql_db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type AdminDb struct {
	dbDriver *sql_driver.DeliveryAppDb
}

func CreateAdminDb(dbDriver *sql_driver.DeliveryAppDb) *AdminDb {
	return &AdminDb{dbDriver: dbDriver}
}

func (ad *AdminDb) GetBy(attr string, value interface{}) (*models.Admin, error) {
	stmt, err := ad.dbDriver.Prepare(fmt.Sprintf(`SELECT adm.id, adm.username, adm.name, adm.surname, adm.email, adm.password, adm.date_of_birth,
	adm.role, a.city, a.street, a.street_num, a.postfix, a.id from administrator adm inner join address a on a.id = adm.address WHERE adm.%s = ?;`, attr))
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

	var admin *models.Admin = models.CreateAdmin()
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

	admin.SetId(id)
	admin.SetName(name)
	admin.SetSurname(surname)
	admin.SetUsername(username)
	admin.SetEmail(email)
	admin.SetPassword(password)
	admin.SetDateOfBirth(dateOfBirth)
	admin.SetRole(role)

	admin.SetAddress(addr)

	return admin, nil
}

func (adb *AdminDb) Update(a *models.Admin) (bool, error) {
	stmt, err := adb.dbDriver.Prepare(`UPDATE administrator SET
	name = ?,
	surname = ?,
	username = ?,
	email = ?,
	password = ?,
	date_of_birth = ?,
	address = ?
	where administrator.id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(a.Name, a.Surname, a.Username, a.Email, a.Password, a.DateOfBirth, a.Address.Id, a.Id)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, err
	}

	return true, nil
}
