package dto

type DelivererOutputDto struct {
	Username          string            `json:"username"`
	Name              string            `json:"name"`
	Surname           string            `json:"surname"`
	Email             string            `json:"email"`
	DateOfBirth       string            `json:"date_of_birth"`
	DeliveryInProgess bool              `json:"delivery_in_progress"`
	Address           *AddressOutputDto `json:"address"`
}
