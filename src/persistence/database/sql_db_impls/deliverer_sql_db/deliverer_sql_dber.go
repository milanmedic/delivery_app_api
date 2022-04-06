package deliverer_sql_db

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
)

type DelivererDber interface {
	AddDeliverer(d models.Deliverer) error
	UpdateProperty(property string, value interface{}, id string) error
	GetBy(attr string, value interface{}) (*models.Deliverer, error)
}
