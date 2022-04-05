package services

import (
	dto "delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type CustomerServicer interface {
	GetCustomer(attr string, value interface{}) (*models.Customer, error)
	Exists(email string) (bool, error)
	ValidateCustomerRegistrationInput(udto dto.CustomerInputDto) error
	CreateCustomer(ud dto.CustomerInputDto) error
}
