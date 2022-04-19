package dto

type BasketInputDto struct {
	Id       string
	Price    int              `json:"price"`
	Articles []basket_article `json:"articles"`
}

type basket_article struct {
	Id       int
	Quantity int
}
