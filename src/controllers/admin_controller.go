package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	addr_service "delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	services "delivery_app_api.mmedic.com/m/v2/src/services/admin_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/jwt_utils"
	"delivery_app_api.mmedic.com/m/v2/src/utils/security"
	"delivery_app_api.mmedic.com/m/v2/src/utils/validations"
	"github.com/gin-gonic/gin"
)

//TODO: Refactor
//TODO: Change all errors to custom error
type AdminController struct {
	adminService     services.AdminServicer
	customerService  customer_service.CustomerServicer
	addrService      addr_service.AddrServicer
	delivererService deliverer_service.DelivererServicer
}

func CreateAdminController(as services.AdminServicer, cs customer_service.CustomerServicer, ads addr_service.AddrServicer, ds deliverer_service.DelivererServicer) *AdminController {
	return &AdminController{adminService: as, customerService: cs, addrService: ads, delivererService: ds}
}

func (ac *AdminController) VerifyCustomer(c *gin.Context) {

	customerID := c.Query("customerID")

	customer, err := ac.customerService.GetCustomer("id", customerID)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the customer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if customer == nil {
		c.Status(http.StatusNotFound)
		return
	}

	if strings.Compare(customer.VerificationStatus, "VERIFIED") == 0 {
		c.String(http.StatusBadRequest, "Customer already verified")
		return
	}

	ac.customerService.UpdateProperty("verification_status", "VERIFIED", customerID)
	if err != nil {
		c.Error(fmt.Errorf("Error while verifying customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}

func (ac *AdminController) AddDeliverer(c *gin.Context) {
	var delivererDto dto.DelivererInputDto
	err := c.BindJSON(&delivererDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = ac.delivererService.ValidateDelivererRegistrationInput(delivererDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	var addrId int
	addr, err := ac.addrService.GetAddr(*delivererDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = ac.addrService.CreateAddress(*delivererDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		delivererDto.Address.Id = addrId
	} else {
		delivererDto.Address.Id = addr.Id
	}

	exists, err := ac.delivererService.Exists(delivererDto.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while checking for a customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		c.String(http.StatusBadRequest, "User with the provided email already exists")
		return
	}

	if !exists {
		err = ac.delivererService.AddDeliverer(delivererDto)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating a customer. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Status(204)
	return
}

func (ac *AdminController) VerifyDeliverer(c *gin.Context) {
	delivererID := c.Query("delivererID")

	deliverer, err := ac.delivererService.GetBy("id", delivererID)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the deliverer info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if deliverer == nil {
		c.Status(http.StatusNotFound)
		return
	}

	if strings.Compare(deliverer.VerificationStatus, "VERIFIED") == 0 {
		c.String(http.StatusBadRequest, "Deliverer already verified")
		return
	}

	ac.delivererService.UpdateProperty("verification_status", "VERIFIED", delivererID)
	if err != nil {
		c.Error(fmt.Errorf("Error while verifying customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}

func (ac *AdminController) AdminLogin(c *gin.Context) {
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

	admin, err := ac.adminService.GetBy("email", credentials.Email)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the admin info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if admin == nil {
		c.String(http.StatusNotFound, "Admin with the provided email was not found.")
		return
	}

	if !security.CheckPasswordHash(credentials.Password, admin.Password) {
		c.String(http.StatusUnauthorized, "Wrong password.")
		return
	}

	claims := jwt_utils.CreateClaims()
	claims.Email = credentials.Email
	claims.Role = admin.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func (ac *AdminController) GetInfo(c *gin.Context) {
	var id string = c.Param("id")

	adminOut, err := ac.adminService.GetAdminInfo(id)
	if err != nil {
		c.Error(fmt.Errorf("Error while retrieving the admin info. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if adminOut == nil {
		c.String(http.StatusNotFound, "Admin doesn't exist.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": adminOut})
	return
}
