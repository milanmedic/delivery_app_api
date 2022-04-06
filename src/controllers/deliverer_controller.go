package controllers

import (
	"fmt"
	"net/http"

	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
)

type DelivererController struct {
	delivererService deliverer_service.DelivererServicer
	addrService      addr_service.AddrServicer
}

func CreateDelivererController(ds deliverer_service.DelivererServicer, as addr_service.AddrServicer) *DelivererController {
	return &DelivererController{delivererService: ds, addrService: as}
}

func (dc *DelivererController) DelivererLogin(c *gin.Context) {
	var credentials jwt_utils.Credentials
	err := c.BindJSON(&credentials)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = validations.ValidateEmail(credentials.Email)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	deliverer, err := dc.delivererService.GetBy("email", credentials.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the deliverer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if deliverer == nil {
		c.String(http.StatusNotFound, "Deliverer with the provided email was not found.")
		return
	}

	if !security.CheckPasswordHash(credentials.Password, deliverer.Password) {
		c.String(http.StatusUnauthorized, "Wrong password.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = credentials.Email
	claims.UserId = deliverer.Id
	claims.Role = deliverer.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func (dc *DelivererController) GetDelivererInfo(c *gin.Context) {
	var id string = c.Param("id")

	delivererOut, err := dc.delivererService.GetDelivererInfo(id)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the deliverer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if delivererOut == nil {
		c.String(http.StatusNotFound, "Deliverer doesn't exist.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": delivererOut})
	return
}
