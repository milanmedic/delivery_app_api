package admin_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type AdminServicer interface {
	GetBy(attr string, value interface{}) (*models.Admin, error)
	GetAdminInfo(id string) (*dto.AdminOutputDto, error)
	UpdateAdmin(id string, adto *dto.AdminInputDto) (bool, error)
	ValidateAdminDataInput(adto dto.AdminInputDto) error
}
