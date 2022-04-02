package usersqldb

type UserDber interface {
	GetById(id string) error
	AddOne() error
	Update() error
	Delete() error
}
