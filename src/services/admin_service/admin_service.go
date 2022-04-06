package admin_service

import (
	"delivery_app_api.mmedic.com/m/v2/src/services/customer_service"
	"delivery_app_api.mmedic.com/m/v2/src/services/deliverer_service"
)

type AdminService struct {
	customerService  customer_service.CustomerServicer
	delivererService deliverer_service.DelivererServicer
}

func CreateAdminService(cs customer_service.CustomerServicer, ds deliverer_service.DelivererServicer) *AdminService {
	return &AdminService{customerService: cs, delivererService: ds}
}

func (as *AdminService) VerifyCustomer(customerID string) error {
	return as.customerService.UpdateProperty("verification_status", "VERIFIED", customerID)
}

func (as *AdminService) VerifyDeliverer(delivererID string) error {
	return as.delivererService.UpdateProperty("verification_status", "VERIFIED", delivererID)
}
