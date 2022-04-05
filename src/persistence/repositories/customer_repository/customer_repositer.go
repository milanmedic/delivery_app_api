package customer_repository

import "delivery_app_api.mmedic.com/m/v2/src/models"

type CustomerRepositer interface {
	CreateCustomer(u models.Customer) error
	GetCustomer(attr string, value interface{}) (*models.Customer, error)
	UpdateProperty(property string, value interface{}, id string) error
}
