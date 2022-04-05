package customer_repository

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

func (cr *CustomerRepository) CreateCustomer(u models.Customer) error {
	return cr.db.AddOne(u)
}

func (cr *CustomerRepository) GetCustomer(attr string, value interface{}) (*models.Customer, error) {
	return cr.db.GetBy(attr, value)
}

func (cr *CustomerRepository) UpdateProperty(property string, value interface{}, id string) error {
	return cr.db.UpdateProperty(property, value, id)
}
