package dto

type AddressInputDto struct {
	Id        int    `json:"id"`
	City      string `json:"city"`
	Street    string `json:"street"`
	StreetNum int    `json:"streetNum"`
	Postfix   string `json:"postfix"`
}
