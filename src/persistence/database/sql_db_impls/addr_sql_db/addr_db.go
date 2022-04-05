package addrsqldb

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type AddrDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateAddrDb(dbDrivers *dbdrivers.DeliveryAppDb) *AddrDb {
	return &AddrDb{dbDriver: dbDrivers}
}

func (ad *AddrDb) GetAddr(addr dto.AddressInputDto) (*models.Address, error) {
	stmt, err := ad.dbDriver.Prepare(`SELECT * from address WHERE city = ? and street = ? and street_num = ? and postfix = ?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(addr.City, addr.Street, addr.StreetNum, addr.Postfix)

	var addrID int
	var street string
	var city string
	var streetNum int
	var postfix string

	err = row.Scan(&addrID, &city, &street, &streetNum, &postfix)
	if err != nil {
		return nil, nil
	}

	foundAddr := new(models.Address)
	foundAddr.Id = addrID
	foundAddr.City = city
	foundAddr.Street = street
	foundAddr.StreetNum = streetNum
	foundAddr.Postfix = postfix

	return foundAddr, nil
}

func (ad *AddrDb) GetByID(id string) (*models.Address, error) {
	stmt, err := ad.dbDriver.Prepare(`SELECT * from address WHERE id = ?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var addrID int
	var street string
	var city string
	var streetNum int
	var postfix string
	err = row.Scan(&addrID, &city, &street, &streetNum, &postfix)
	if err != nil {
		return nil, nil
	}

	addr := new(models.Address)
	addr.Id = addrID
	addr.City = city
	addr.Street = street
	addr.StreetNum = streetNum
	addr.Postfix = postfix

	return addr, nil
}

func (ad *AddrDb) AddOne(addr dto.AddressInputDto) (int, error) {
	tx, err := ad.dbDriver.Begin()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(`INSERT INTO address(city, street, street_num, postfix) VALUES(?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(addr.City, addr.Street, addr.StreetNum, addr.Postfix)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
