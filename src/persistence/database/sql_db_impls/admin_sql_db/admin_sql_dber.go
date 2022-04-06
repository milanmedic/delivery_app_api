package admin_sql_db

import "delivery_app_api.mmedic.com/m/v2/src/models"

type AdminDber interface {
	GetBy(attr string, value interface{}) (*models.Admin, error)
}
