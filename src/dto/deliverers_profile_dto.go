package dto

type DeliverersProfileDto struct {
	Id                 string            `json:"id"`
	Username           string            `json:"username"`
	Name               string            `json:"name"`
	Surname            string            `json:"surname"`
	Email              string            `json:"email"`
	DateOfBirth        string            `json:"dateOfBirth"`
	DeliveryInProgess  bool              `json:"deliveryInProgress"`
	VerificationStatus string            `json:"verificationStatus"`
	Address            *AddressOutputDto `json:"address"`
}
