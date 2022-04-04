package usersqldb

import "delivery_app_api.mmedic.com/m/v2/src/models"

type CustomerDber interface {
	GetBy(attr string, value interface{}) (*models.Customer, error)
	AddOne(c models.Customer) error
	Update() error
	Delete() error
}