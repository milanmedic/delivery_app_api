package addr_repository

import (
	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	addrsqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/addr_sql_db"
)

type AddrRepository struct {
	db addrsqldb.AddrDber
}

func CreateAddrRepository(db addrsqldb.AddrDber) *AddrRepository {
	return &AddrRepository{db: db}
}

func (ar *AddrRepository) CreateAddr(a dto.AddressInputDto) (int, error) {
	return ar.db.AddOne(a)
}

func (ar *AddrRepository) GetAddr(a dto.AddressInputDto) (*models.Address, error) {
	return ar.db.GetAddr(a)
}

func (ar *AddrRepository) GetById(id string) (*models.Address, error) {
	return ar.db.GetByID(id)
}

func (ar *AddrRepository) GetUserAddress(userId string) (*dto.AddressOutputDto, error) {
	return ar.db.GetUserAddress(userId)
}
