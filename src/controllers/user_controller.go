package controllers

import (
	user_services "delivery_app_api.mmedic.com/m/v2/src/services/user_service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service user_services.UserServicer
}

func CreateUserController(service user_services.UserServicer) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Register(c *gin.Context) {
	c.String(200, "Register")
}

func (uc *UserController) Login(c *gin.Context) {
	c.String(200, "Login")
}
