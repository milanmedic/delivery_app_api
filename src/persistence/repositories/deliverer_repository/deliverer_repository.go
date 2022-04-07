package deliverer_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/deliverer_sql_db"
)

type DelivererRepository struct {
	db deliverer_sql_db.DelivererDber
}

func CreateDelivererRepository(db deliverer_sql_db.DelivererDber) *DelivererRepository {
	return &DelivererRepository{db: db}
}

func (dr *DelivererRepository) AddDeliverer(d models.Deliverer) error {
	return dr.db.AddDeliverer(d)
}

func (dr *DelivererRepository) GetBy(attr string, value interface{}) (*models.Deliverer, error) {
	return dr.db.GetBy(attr, value)
}

func (dr *DelivererRepository) UpdateProperty(property string, value interface{}, id string) error {
	return dr.db.UpdateProperty(property, value, id)
}

func (dr *DelivererRepository) Update(d *models.Deliverer) (bool, error) {
	return dr.db.Update(d)
}
