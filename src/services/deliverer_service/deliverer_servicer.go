package deliverer_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type DelivererServicer interface {
	AddDeliverer(ddto dto.DelivererInputDto) error
	UpdateProperty(property string, value interface{}, id string) error
	ValidateDelivererRegistrationInput(udto dto.DelivererInputDto) error
	GetBy(attr string, value interface{}) (*models.Deliverer, error)
	Exists(email string) (bool, error)
	GetDelivererInfo(id string) (*dto.DelivererOutputDto, error)
}
