package order_sql_db

import (
	"database/sql"
	"fmt"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	dbdrivers "delivery_app_api.mmedic.com/m/v2/src/persistence/database/db_drivers/sql_driver"
)

type OrderDb struct {
	dbDriver *dbdrivers.DeliveryAppDb
}

func CreateOrderDb(dbDriver *dbdrivers.DeliveryAppDb) *OrderDb {
	return &OrderDb{dbDriver: dbDriver}
}

func (odb *OrderDb) CreateOrder(odto dto.OrderInputDto) error {
	tx, err := odb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO customer_order(comment, address, basket, customer) VALUES(?, ?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(odto.Comment, odto.Address.Id, odto.Basket.Id, odto.CustomerID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (odb *OrderDb) GetOrdersByUserId(id string) ([]models.Order, error) {
	stmt, err := odb.dbDriver.Prepare(`select o.id, o.comment, o.status, o.accepted,
	ifnull(d.name, '') as 'name', ifnull(d.surname, '') as 'surname',
	b.price as "total",
	a.name as "article_name", a.description as "article_description" , a.price as "article_price",
	ab.article_quantity as "article_quantity",
	addr.city, addr.street, addr.street_num, addr.postfix
	from customer_order o
	inner join customer c
	on o.customer = c.id
	left join deliverer d
	on o.deliverer = d.id
	inner join basket b
	on o.basket = b.id
	inner join article_basket ab
	on b.id = ab.basket
	inner join article a
	on ab.article = a.id
	inner join address addr
	on addr.id = o.address
	where c.id=?;`)
	if err != nil {
		return nil, err
	}
	var rows *sql.Rows
	var orders []models.Order = []models.Order{}

	rows, err = stmt.Query(id)

	var previousOrderId string
	var previousArticleName string
	var o *models.Order
	var ba *dto.BasketArticleOutput
	for rows.Next() {
		var orderId string
		var orderComment string
		var orderStatus string
		var accepted bool
		var delivererName string
		var delivererSurname string
		var city string
		var street string
		var street_num int
		var postfix string
		var totalPrice int
		var articleName string
		var articleDescription string
		var articlePrice string
		var articleQuantity int

		if err := rows.Scan(&orderId, &orderComment, &orderStatus, &accepted, &delivererName, &delivererSurname, &totalPrice, &articleName, &articleDescription, &articlePrice, &articleQuantity, &city, &street, &street_num, &postfix); err != nil {
			return nil, err
		}
		if orderId != previousOrderId {
			previousOrderId = orderId
			o = new(models.Order)
			o.Id = orderId
			o.Comment = orderComment
			o.Status = orderStatus
			o.Accepted = accepted
			o.DelivererName = delivererName
			o.DelivererSurname = delivererSurname
			o.Address.City = city
			o.Address.Street = street
			o.Address.StreetNum = street_num
			o.Address.Postfix = postfix
			o.Basket.Price = totalPrice
			previousArticleName = articleName
			ba = new(dto.BasketArticleOutput)
			ba.Name = articleName
			ba.Description = articleDescription
			ba.Price = articlePrice
			ba.Quantity = articleQuantity
			o.Basket.Articles = append(o.Basket.Articles, *ba)
		} else {
			if strings.Compare(articleName, previousArticleName) != 0 {
				previousArticleName = articleName
				ba = new(dto.BasketArticleOutput)
				ba.Name = articleName
				ba.Description = articleDescription
				ba.Price = articlePrice
				ba.Quantity = articleQuantity
				o.Basket.Articles = append(o.Basket.Articles, *ba)
				orders = orders[:len(orders)-1]
			}
		}

		orders = append(orders, *o)
	}
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return orders, nil
}

func (odb *OrderDb) DeleteOrder(attr string, value interface{}) error {
	tx, err := odb.dbDriver.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(fmt.Sprintf(`DELETE FROM customer_order where %s=;`, attr))
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	switch value.(type) {
	case int:
		_, err = stmt.Exec(value.(int))
	case float64:
		_, err = stmt.Exec(value.(float64))
	case bool:
		_, err = stmt.Exec(value.(bool))
	case string:
		_, err = stmt.Exec(value.(string))
	}
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (odb *OrderDb) GetOrderBasketID(id string) (string, error) {
	stmt, err := odb.dbDriver.Prepare(` SELECT basket from customer_order where id = ?;`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var row *sql.Row

	row = stmt.QueryRow(id)

	var basketID string
	err = row.Scan(&basketID)
	if err != nil {
		return "", nil
	}

	return basketID, nil
}

func (odb *OrderDb) GetOrderStatus(id string) (string, error) {
	stmt, err := odb.dbDriver.Prepare(`SELECT status from customer_order where id = ?;`)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var row *sql.Row

	row = stmt.QueryRow(id)

	var basketID string
	err = row.Scan(&basketID)
	if err != nil {
		return "", nil
	}

	return basketID, nil
}

func (odb *OrderDb) UpdateProperty(property string, value interface{}, id string) error {
	stmt, err := odb.dbDriver.Prepare(fmt.Sprintf(`UPDATE customer_order SET %s = ? where customer_order.id = ?;`, property))
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

var matchAllOrders string = `select o.id, o.comment, o.status, o.accepted,
ifnull(d.name, '') as 'name', ifnull(d.surname, '') as 'surname',
b.price as "total",
a.name as "article_name", a.description as "article_description" , a.price as "article_price",
ab.article_quantity as "article_quantity",
addr.city, addr.street, addr.street_num, addr.postfix
from customer_order o
inner join customer c
on o.customer = c.id
left join deliverer d
on o.deliverer = d.id
inner join basket b
on o.basket = b.id
inner join article_basket ab
on b.id = ab.basket
inner join article a
on ab.article = a.id
inner join address addr
on addr.id = o.address;`

var matchOnStatusAndAcceptance string = `select o.id, o.comment, o.status, o.accepted,
ifnull(d.name, '') as 'name', ifnull(d.surname, '') as 'surname',
b.price as "total",
a.name as "article_name", a.description as "article_description" , a.price as "article_price",
ab.article_quantity as "article_quantity",
addr.city, addr.street, addr.street_num, addr.postfix
from customer_order o
inner join customer c
on o.customer = c.id
left join deliverer d
on o.deliverer = d.id
inner join basket b
on o.basket = b.id
inner join article_basket ab
on b.id = ab.basket
inner join article a
on ab.article = a.id
inner join address addr
on addr.id = o.address
where o.status LIKE ? AND o.accepted=?;
`

var matchOnStatus string = `select o.id, o.comment, o.status, o.accepted,
ifnull(d.name, '') as 'name', ifnull(d.surname, '') as 'surname',
b.price as "total",
a.name as "article_name", a.description as "article_description" , a.price as "article_price",
ab.article_quantity as "article_quantity",
addr.city, addr.street, addr.street_num, addr.postfix
from customer_order o
inner join customer c
on o.customer = c.id
left join deliverer d
on o.deliverer = d.id
inner join basket b
on o.basket = b.id
inner join article_basket ab
on b.id = ab.basket
inner join article a
on ab.article = a.id
inner join address addr
on addr.id = o.address
where o.status LIKE ?;
`

func (odb *OrderDb) GetAllOrders(deliveryStatus string, accepted ...string) ([]models.Order, error) {

	var stmt *sql.Stmt
	var err error

	// NO QP
	if strings.Compare(deliveryStatus, "") == 0 && len(accepted) <= 0 {
		stmt, err = odb.dbDriver.Prepare(matchAllOrders)
		// ONLY STATUS
	} else if len(accepted) <= 0 {
		stmt, err = odb.dbDriver.Prepare(matchOnStatus)
		// ALL
	} else {
		stmt, err = odb.dbDriver.Prepare(matchOnStatusAndAcceptance)
	}

	if err != nil {
		return nil, err
	}
	var rows *sql.Rows
	var orders []models.Order = []models.Order{}

	if strings.Compare(deliveryStatus, "") == 0 && len(accepted) <= 0 {
		rows, err = stmt.Query()
	} else if strings.Compare(deliveryStatus, "") != 0 && len(accepted) <= 0 {
		rows, err = stmt.Query(deliveryStatus)
	} else {
		rows, err = stmt.Query(deliveryStatus, accepted[0])
	}

	var previousOrderId string
	var previousArticleName string
	var o *models.Order
	var ba *dto.BasketArticleOutput
	for rows.Next() {
		var orderId string
		var orderComment string
		var orderStatus string
		var accepted bool
		var delivererName string
		var delivererSurname string
		var city string
		var street string
		var street_num int
		var postfix string
		var totalPrice int
		var articleName string
		var articleDescription string
		var articlePrice string
		var articleQuantity int

		if err := rows.Scan(&orderId, &orderComment, &orderStatus, &accepted, &delivererName, &delivererSurname, &totalPrice, &articleName, &articleDescription, &articlePrice, &articleQuantity, &city, &street, &street_num, &postfix); err != nil {
			return nil, err
		}
		if orderId != previousOrderId {
			previousOrderId = orderId
			o = new(models.Order)
			o.Id = orderId
			o.Comment = orderComment
			o.Status = orderStatus
			o.Accepted = accepted
			o.DelivererName = delivererName
			o.DelivererSurname = delivererSurname
			o.Address.City = city
			o.Address.Street = street
			o.Address.StreetNum = street_num
			o.Address.Postfix = postfix
			o.Basket.Price = totalPrice
			previousArticleName = articleName
			ba = new(dto.BasketArticleOutput)
			ba.Name = articleName
			ba.Description = articleDescription
			ba.Price = articlePrice
			ba.Quantity = articleQuantity
			o.Basket.Articles = append(o.Basket.Articles, *ba)
		} else {
			if strings.Compare(articleName, previousArticleName) != 0 {
				previousArticleName = articleName
				ba = new(dto.BasketArticleOutput)
				ba.Name = articleName
				ba.Description = articleDescription
				ba.Price = articlePrice
				ba.Quantity = articleQuantity
				o.Basket.Articles = append(o.Basket.Articles, *ba)
				orders = orders[:len(orders)-1]
			}
		}

		orders = append(orders, *o)
	}
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return orders, nil
}
