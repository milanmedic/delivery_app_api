package article_sql_db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type ArticleDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateArticleDb(dbDriver *dbdrivers.DeliveryAppDb) *ArticleDb {
	return &ArticleDb{dbDriver: dbDriver}
}

func (adb *ArticleDb) GetBy(attr string, value interface{}) (*models.Article, error) {
	stmt, err := adb.dbDriver.Prepare(fmt.Sprintf(` SELECT * from article where %s = ?;`, attr))
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

	var article *models.Article = models.CreateArticle()
	var id int
	var name string
	var description string
	var price int

	err = row.Scan(&id, &name, &description, &price)
	if err != nil {
		return nil, nil
	}

	article.SetId(id)
	article.SetName(name)
	article.SetDescription(description)
	article.SetPrice(price)

	return article, nil
}

func (adb *ArticleDb) AddOne(a models.Article) error {
	tx, err := adb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO article(name, description, price) VALUES(?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Name, a.Description, a.Price)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (adb *ArticleDb) GetAll() ([]*models.Article, error) {
	stmt, err := adb.dbDriver.Prepare(`SELECT * FROM article;`)
	if err != nil {
		return nil, err
	}
	var articles []*models.Article = []*models.Article{}

	rows, err := stmt.Query()
	for rows.Next() {
		var a models.Article
		var id int
		var name string
		var desc string
		var price int
		if err := rows.Scan(&id, &name, &desc, &price); err != nil {
			return nil, err
		}
		a.SetId(id)
		a.SetName(name)
		a.SetDescription(desc)
		a.SetPrice(price)
		articles = append(articles, &a)
	}
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return articles, nil
}
