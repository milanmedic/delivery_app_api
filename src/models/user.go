package models

type User struct {
	Id                 string
	Username           string
	Name               string
	Surname            string
	Email              string
	Password           string
	DateOfBirth        string
	Address            *Address
	Role               string
	VerificationStatus string
}

func CreateUser() *User {
	return &User{}
}

func (u *User) SetId(id string) {
	u.Id = id
}
func (u *User) SetUsername(un string) {
	u.Username = un
}
func (u *User) SetName(name string) {
	u.Name = name
}
func (u *User) SetSurname(sn string) {
	u.Surname = sn
}
func (u *User) SetEmail(email string) {
	u.Email = email
}
func (u *User) SetPassword(pass string) {
	u.Password = pass
}
func (u *User) SetDateOfBirth(dob string) {
	u.DateOfBirth = dob
}
func (u *User) SetAddress(addr *Address) {
	u.Address = addr
}
func (u *User) SetRole(role string) {
	u.Role = role
}
func (u *User) SetVerificationStatus(st string) {
	u.VerificationStatus = st
}
