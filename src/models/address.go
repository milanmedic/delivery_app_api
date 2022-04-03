package models

type Address struct {
	Id        int
	City      string
	Street    string
	StreetNum int
	Postfix   string
}

func CreateAddress(id, num int, city, street, pf string) *Address {
	return &Address{
		Id:        id,
		StreetNum: num,
		City:      city,
		Street:    street,
		Postfix:   pf,
	}
}

func (a *Address) SetID(id int) {
	a.Id = id
}
func (a *Address) SetCity(city string) {
	a.City = city
}
func (a *Address) SetStreet(street string) {
	a.Street = street
}
func (a *Address) SetStreetNum(num int) {
	a.StreetNum = num
}
func (a *Address) SetStreetPostfix(pf string) {
	a.Postfix = pf
}
