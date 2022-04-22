package dto

type BasketOutputDto struct {
	Price    int                   `json:"price"`
	Articles []BasketArticleOutput `json:"articles"`
}

type BasketArticleOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    int    `json:"quantity"`
}
