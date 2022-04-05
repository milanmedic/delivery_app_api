package customer_sql_db

import "delivery_app_api.mmedic.com/m/v2/src/models"

type CustomerDber interface {
	GetBy(attr string, value interface{}) (*models.Customer, error)
	AddOne(c models.Customer) error
	UpdateProperty(property string, value interface{}, id string) error
	Delete() error
}
