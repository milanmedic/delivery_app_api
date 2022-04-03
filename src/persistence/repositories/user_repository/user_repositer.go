package userrepository

import "delivery_app_api.mmedic.com/m/v2/src/models"

type UserRepositer interface {
	CreateUser(u models.User) error
	GetUser(attr string, value interface{}) (*models.User, error)
}
