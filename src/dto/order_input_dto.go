package dto

type OrderInputDto struct {
	CustomerID string
	Comment    string          `json:"comment"`
	Address    AddressInputDto `json:"address"`
	Basket     BasketInputDto  `json:"basket"`
}
