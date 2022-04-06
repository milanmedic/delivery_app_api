package models

type Admin struct {
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

func CreateAdmin() *Admin {
	return &Admin{}
}

func (a *Admin) SetId(id string) {
	a.Id = id
}
func (a *Admin) SetUsername(un string) {
	a.Username = un
}
func (a *Admin) SetName(name string) {
	a.Name = name
}
func (a *Admin) SetSurname(sn string) {
	a.Surname = sn
}
func (a *Admin) SetEmail(email string) {
	a.Email = email
}
func (a *Admin) SetPassword(pass string) {
	a.Password = pass
}
func (a *Admin) SetDateOfBirth(dob string) {
	a.DateOfBirth = dob
}
func (a *Admin) SetAddress(addr *Address) {
	a.Address = addr
}
func (a *Admin) SetRole(role string) {
	a.Role = role
}
