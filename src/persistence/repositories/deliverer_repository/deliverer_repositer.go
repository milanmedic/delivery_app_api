package deliverer_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type DelivererRepositer interface {
	AddDeliverer(d models.Deliverer) error
	UpdateProperty(property string, value interface{}, id string) error
	GetBy(attr string, value interface{}) (*models.Deliverer, error)
	Update(d *models.Deliverer) (bool, error)
	GetAll() ([]dto.DeliverersProfileDto, error)
}
