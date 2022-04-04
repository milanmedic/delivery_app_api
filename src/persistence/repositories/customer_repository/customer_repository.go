package userrepository

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	customersqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/customer_sql_db"
)

type CustomerRepository struct {
	db customersqldb.CustomerDber
}

func CreateCustomerRepository(db customersqldb.CustomerDber) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (ur *CustomerRepository) CreateCustomer(u models.Customer) error {
	return ur.db.AddOne(u)
}

func (ur *CustomerRepository) GetCustomer(attr string, value interface{}) (*models.Customer, error) {
	return ur.db.GetBy(attr, value)
}
