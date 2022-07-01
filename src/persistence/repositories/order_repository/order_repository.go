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

func (or *OrderRepository) DeleteOrder(attr string, value interface{}) error {
	return or.db.DeleteOrder(attr, value)
}

func (or *OrderRepository) GetOrderBasketID(id string) (string, error) {
	return or.db.GetOrderBasketID(id)
}

func (or *OrderRepository) GetOrderStatus(id string) (string, error) {
	return or.db.GetOrderStatus(id)
}

func (or *OrderRepository) UpdateProperty(property string, value interface{}, id string) error {
	return or.db.UpdateProperty(property, value, id)
}

func (or *OrderRepository) GetAllOrders(deliveryStatus string, accepted ...string) ([]models.Order, error) {
	return or.db.GetAllOrders(deliveryStatus, accepted...)
}
