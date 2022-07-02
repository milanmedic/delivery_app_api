package deliverer_sql_db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
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
	//TODO: Remove reflection
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

func (dDb *DelivererDb) Update(d *models.Deliverer) (bool, error) {
	stmt, err := dDb.dbDriver.Prepare(`UPDATE deliverer SET
	name = ?,
	surname = ?,
	username = ?,
	email = ?,
	password = ?,
	date_of_birth = ?,
	address = ?
	where deliverer.id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(d.Name, d.Surname, d.Username, d.Email, d.Password, d.DateOfBirth, d.Address.Id, d.Id)
	if err != nil {
		return false, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dDb *DelivererDb) GetAll() ([]dto.DeliverersProfileDto, error) {
	stmt, err := dDb.dbDriver.Prepare(`select d.id, d.name, d.surname, d.username,
	d.email, d.date_of_birth, d.delivery_in_progress, d.verification_status,
	a.city, a.street, a.street_num, a.postfix from deliverer d
	inner join address a
	on d.address = a.id;`)

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	var deliverers []dto.DeliverersProfileDto = []dto.DeliverersProfileDto{}

	rows, err = stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var deliverer *dto.DeliverersProfileDto = new(dto.DeliverersProfileDto)
		var id string
		var name string
		var username string
		var surname string
		var email string
		var dateOfBirth string
		var deliveryInProgress bool
		var city string
		var street string
		var streetNum int
		var postfix string
		var verificationStatus string

		if err := rows.Scan(&id, &name, &surname, &username, &email, &dateOfBirth, &deliveryInProgress, &verificationStatus, &city, &street, &streetNum, &postfix); err != nil {
			return nil, err
		}

		deliverer.Id = id
		deliverer.Name = name
		deliverer.Surname = surname
		deliverer.Username = username
		deliverer.Email = email
		deliverer.DateOfBirth = dateOfBirth
		deliverer.DeliveryInProgess = deliveryInProgress
		deliverer.VerificationStatus = verificationStatus
		deliverer.Address = new(dto.AddressOutputDto)
		deliverer.Address.City = city
		deliverer.Address.Street = street
		deliverer.Address.StreetNum = streetNum
		deliverer.Address.Postfix = postfix

		deliverers = append(deliverers, *deliverer)
	}

	return deliverers, nil
}
