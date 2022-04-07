package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
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

	if strings.Compare(deliverer.VerificationStatus, "VERIFIED") != 0 {
		c.String(http.StatusUnauthorized, "Account not verified.")
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

func (dc *DelivererController) UpdateDeliverer(c *gin.Context) {
	var delivererID string = c.Param("id")
	var delivererDto dto.DelivererInputDto
	err := c.BindJSON(&delivererDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dc.delivererService.ValidateDelivererDataInput(delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var addrId int
	addr, err := dc.addrService.GetAddr(*delivererDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = dc.addrService.CreateAddress(*delivererDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		delivererDto.Address.Id = addrId
	} else {
		delivererDto.Address.Id = addr.Id
	}

	res, err := dc.delivererService.UpdateDeliverer(delivererID, &delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while updating the deliverer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !res {
		c.Error(fmt.Errorf("Deliverer has failed to update or doesn't exist"))
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}
