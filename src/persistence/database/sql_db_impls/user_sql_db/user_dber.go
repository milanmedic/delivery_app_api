package usersqldb

import "delivery_app_api.mmedic.com/m/v2/src/models"

type UserDber interface {
	GetBy(attr string, value interface{}) error
	AddOne(u models.User) error
	Update() error
	Delete() error
}
