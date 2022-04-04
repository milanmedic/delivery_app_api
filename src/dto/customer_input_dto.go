package dto

type CustomerInputDto struct {
	Username    string           `json:"username"`
	Name        string           `json:"name"`
	Surname     string           `json:"surname"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	DateOfBirth string           `json:"dateOfBirth"`
	Address     *AddressInputDto `json:"address"`
}
