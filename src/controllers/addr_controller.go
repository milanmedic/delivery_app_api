package controllers

import (
	"fmt"
	"net/http"

	"delivery_app_api.mmedic.com/m/v2/src/services/addr_service"
	"github.com/gin-gonic/gin"
)

type AddrController struct {
	addrService addr_service.AddrServicer
}

func CreateAddrController(as addr_service.AddrServicer) *AddrController {
	return &AddrController{addrService: as}
}

func (ac *AddrController) GetCustomerAddr(c *gin.Context) {
	id, ok := c.Get("user_id")

	if !ok {
		c.Error(fmt.Errorf("user not provided in token"))
		c.Status(http.StatusInternalServerError)
		return
	}

	addr, err := ac.addrService.GetUserAddress(id.(string))
	if err != nil {
		c.Error(fmt.Errorf("error while searching for the article. \nReason: %s", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if addr == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, addr)
}
