package userrepository

import (
	"delivery_app_api.mmedic.com/m/v2/src/models"
	usersqldb "delivery_app_api.mmedic.com/m/v2/src/persistence/database/sql_db_impls/user_sql_db"
)

type UserRepository struct {
	db usersqldb.UserDber
}

func CreateUserRepository(db usersqldb.UserDber) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(u models.User) error {
	return ur.db.AddOne(u)
}

func (ur *UserRepository) GetUser(attr string, value interface{}) (*models.User, error) {
	return ur.db.GetBy(attr, value)
}
