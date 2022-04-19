package dto

type OrderInputDto struct {
	CustomerID string          `json:"customer_id"`
	Comment    string          `json:"comment"`
	Address    AddressInputDto `json:"address"`
	Basket     BasketInputDto  `json:"basket"`
}
