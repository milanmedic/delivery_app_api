package basket_sql_db

import (
	"fmt"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type BasketDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateBasketDb(dbDriver *dbdrivers.DeliveryAppDb) *BasketDb {
	return &BasketDb{dbDriver: dbDriver}
}

func (bdb *BasketDb) AddBasket(bdto dto.BasketInputDto) error {
	tx, err := bdb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmts := []string{}
	s1 := formBasketInsertion(bdto)
	s2 := fmt.Sprintf(`INSERT INTO basket(id, price) VALUES('%s', %d);`, bdto.Id, bdto.Price)
	stmts = append(stmts, s2, s1)

	for _, s := range stmts {
		stmt, err := tx.Prepare(s)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func formBasketInsertion(bdto dto.BasketInputDto) string {
	var stmt string = `INSERT INTO article_basket(article, basket, article_quantity) VALUES`
	for i, a := range bdto.Articles {
		if i < len(bdto.Articles)-1 {
			stmt += fmt.Sprintf("(%d, '%s', %d),", a.Id, bdto.Id, a.Quantity)
		} else {
			stmt += fmt.Sprintf("(%d, '%s', %d);", a.Id, bdto.Id, a.Quantity)
		}
	}

	return stmt
}

func (bdb *BasketDb) DeleteBasket(bId string) error {
	tx, err := bdb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmts := []string{}
	s1 := fmt.Sprintf("DELETE FROM article_basket where basket = '%s';", bId)
	s2 := fmt.Sprintf("DELETE FROM basket where id = '%s';", bId)
	stmts = append(stmts, s1, s2)

	for _, s := range stmts {
		stmt, err := tx.Prepare(s)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
