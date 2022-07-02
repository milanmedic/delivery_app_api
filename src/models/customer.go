package models

type Customer struct {
	Id          string
	Username    string
	Name        string
	Surname     string
	Email       string
	Password    string
	DateOfBirth string
	Address     *Address
	Role        string
}

func CreateCustomer() *Customer {
	return &Customer{}
}

func (u *Customer) SetId(id string) {
	u.Id = id
}
func (u *Customer) SetUsername(un string) {
	u.Username = un
}
func (u *Customer) SetName(name string) {
	u.Name = name
}
func (u *Customer) SetSurname(sn string) {
	u.Surname = sn
}
func (u *Customer) SetEmail(email string) {
	u.Email = email
}
func (u *Customer) SetPassword(pass string) {
	u.Password = pass
}
func (u *Customer) SetDateOfBirth(dob string) {
	u.DateOfBirth = dob
}
func (u *Customer) SetAddress(addr *Address) {
	u.Address = addr
}
func (u *Customer) SetRole(role string) {
	u.Role = role
}
