package order_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/order_sql_db"
)

type OrderRepository struct {
	db order_sql_db.OrderDber
}

func CreateOrderRepository(db order_sql_db.OrderDber) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrder(odto dto.OrderInputDto) error {
	return or.db.CreateOrder(odto)
}

func (or *OrderRepository) GetOrdersByUserId(id string) ([]models.Order, error) {
	return or.db.GetOrdersByUserId(id)
}
