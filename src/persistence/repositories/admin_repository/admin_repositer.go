package admin_repository

import "delivery_app_api.mmedic.com/m/v2/src/models"

type AdminRepositer interface {
	GetBy(attr string, value interface{}) (*models.Admin, error)
	Update(c *models.Admin) (bool, error)
}
