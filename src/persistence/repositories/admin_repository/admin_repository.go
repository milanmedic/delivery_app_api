package admin_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/admin_sql_db"
)

type AdminRepository struct {
	db admin_sql_db.AdminDber
}

func CreateAdminRepository(db admin_sql_db.AdminDber) *AdminRepository {
	return &AdminRepository{db: db}
}

func (ar *AdminRepository) GetBy(property string, value interface{}) (*models.Admin, error) {
	return ar.db.GetBy(property, value)
}

func (ar *AdminRepository) Update(a *models.Admin) (bool, error) {
	return ar.db.Update(a)
}
