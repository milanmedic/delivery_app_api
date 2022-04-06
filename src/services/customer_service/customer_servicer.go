package customer_service

import (
	dto "delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type CustomerServicer interface {
	GetBy(attr string, value interface{}) (*models.Customer, error)
	Exists(email string) (bool, error)
	ValidateCustomerRegistrationInput(udto dto.CustomerInputDto) error
	CreateCustomer(ud dto.CustomerInputDto) error
	UpdateProperty(property string, value interface{}, id string) error
	GetCustomerInfo(id string) (*dto.CustomerOutputDto, error)
}
