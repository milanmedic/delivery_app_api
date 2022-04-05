package userrepository

import "delivery_app_api.mmedic.com/m/v2/src/models"

type CustomerRepositer interface {
	CreateCustomer(u models.Customer) error
	GetCustomer(attr string, value interface{}) (*models.Customer, error)
}
