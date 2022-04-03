package dto

type AddressInputDto struct {
	Id        int    `json:"id"`
	City      string `json:"city"`
	Street    string `json:"street"`
	StreetNum int    `json:"street_num"`
	Postfix   string `json:"postfix"`
}
