package controllers

import (
	"fmt"
	"net/http"
	"strings"

	services "delivery_app_api.mmedic.com/m/v2/src/services/admin_service"
	customer_service "delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService    services.AdminServicer
	customerService customer_service.CustomerServicer
}

func CreateAdminController(adminService services.AdminServicer, customerService customer_service.CustomerServicer) *AdminController {
	return &AdminController{adminService: adminService, customerService: customerService}
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

	// MOVE THIS TO ADMIN_SERVICE
	err = ac.adminService.VerifyCustomer(customerID)
	if err != nil {
		c.Error(fmt.Errorf("Error while verifying customer. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
	return
}
