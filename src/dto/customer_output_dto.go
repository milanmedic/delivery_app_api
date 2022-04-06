package dto

type CustomerOutputDto struct {
	Username    string            `json:"username"`
	Name        string            `json:"name"`
	Surname     string            `json:"surname"`
	Email       string            `json:"email"`
	DateOfBirth string            `json:"dateOfBirth"`
	Address     *AddressOutputDto `json:"address"`
}
