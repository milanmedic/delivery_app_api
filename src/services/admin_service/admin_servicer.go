package admin_service

type AdminServicer interface {
	VerifyCustomer(customerID string) error
	VerifyDeliverer(delivererID string) error
}