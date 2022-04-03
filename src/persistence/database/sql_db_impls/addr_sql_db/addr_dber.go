package addrsqldb

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type AddrDber interface {
	GetByID(id string) (*models.Address, error)
	AddOne(u dto.AddressInputDto) (int, error)
}
