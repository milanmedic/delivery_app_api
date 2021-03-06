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
	"delivery_app_api.mmedic.com/m/v2/src/services/mailing_service"
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

	mail := mailing_service.CreateEmail(deliverer.Email, "Account Validation Status", "Your account has been validated.")
	err = mail.SendMail()
	if err != nil {
		c.Error(fmt.Errorf("Error while sending the verification email. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "Verified")
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
	claims.UserId = admin.Id
	claims.Role = admin.Role

	tokenString, err := jwt_utils.CreateToken(claims)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}

func (ac *AdminController) GetAdminInfo(c *gin.Context) {
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

func (ac *AdminController) UpdateAdmin(c *gin.Context) {
	var adminID string = c.Param("id")
	var adminDto dto.AdminInputDto
	err := c.BindJSON(&adminDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = ac.adminService.ValidateAdminDataInput(adminDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while validating input. \nReason: %s", err.Error()))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var addrId int
	addr, err := ac.addrService.GetAddr(*adminDto.Address)
	if err != nil {
		c.Error(fmt.Errorf("Error while searching for address. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		addrId, err = ac.addrService.CreateAddress(*adminDto.Address)
		if err != nil {
			c.Error(fmt.Errorf("Error while creating an address. \nReason: %s", err.Error()))
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		adminDto.Address.Id = addrId
	} else {
		adminDto.Address.Id = addr.Id
	}

	res, err := ac.adminService.UpdateAdmin(adminID, &adminDto)
	if err != nil {
		c.Error(fmt.Errorf("Error while updating the admin. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !res {
		c.Error(fmt.Errorf("Admin has failed to update or doesn't exist"))
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}

func (ac *AdminController) GetAllDeliverers(c *gin.Context) {
	deliverers, err := ac.delivererService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, deliverers)
}
