package services

import (
	user_repo "delivery_app_api.mmedic.com/m/v2/src/persistence/repositories/user_repository"
)

type UserService struct {
	repository user_repo.UserRepositer
}

func CreateUserService(repo user_repo.UserRepositer) *UserService {
	return &UserService{repository: repo}
}
