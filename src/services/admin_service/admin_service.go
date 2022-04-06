package admin_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
)

type AdminService struct {
	customerService customer_service.CustomerServicer
}

func CreateAdminService(cs customer_service.CustomerServicer) *AdminService {
	return &AdminService{customerService: cs}
}
