package services

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type AddrServicer interface {
	GetById(id string) (*models.Address, error)
	CreateAddress(a dto.AddressInputDto) (int, error)
}
