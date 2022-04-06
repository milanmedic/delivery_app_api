package dto

type AddressOutputDto struct {
	City      string `json:"city"`
	Street    string `json:"street"`
	StreetNum int    `json:"street_num"`
	Postfix   string `json:"postfix"`
}
