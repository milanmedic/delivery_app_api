package models

type Article struct {
	Id          int
	Name        string
	Description string
	Price       int
}

func CreateArticle() *Article {
	return &Article{}
}

func (a *Article) SetId(id int) {
	a.Id = id
}

func (a *Article) SetName(n string) {
	a.Name = n
}

func (a *Article) SetDescription(d string) {
	a.Description = d
}

func (a *Article) SetPrice(p int) {
	a.Price = p
}
