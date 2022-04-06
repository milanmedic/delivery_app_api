package models

type Deliverer struct {
	Id                 string
	Username           string
	Name               string
	Surname            string
	Email              string
	Password           string
	DateOfBirth        string
	Address            *Address
	Role               string
	DeliveryInProgress bool
	VerificationStatus string
}

func CreateDeliverer() *Deliverer {
	return &Deliverer{}
}

func (d *Deliverer) SetId(id string) {
	d.Id = id
}
func (d *Deliverer) SetUsername(un string) {
	d.Username = un
}
func (d *Deliverer) SetName(name string) {
	d.Name = name
}
func (d *Deliverer) SetSurname(sn string) {
	d.Surname = sn
}
func (d *Deliverer) SetEmail(email string) {
	d.Email = email
}
func (d *Deliverer) SetPassword(pass string) {
	d.Password = pass
}
func (d *Deliverer) SetDateOfBirth(dob string) {
	d.DateOfBirth = dob
}
func (d *Deliverer) SetAddress(addr *Address) {
	d.Address = addr
}
func (d *Deliverer) SetRole(role string) {
	d.Role = role
}
func (d *Deliverer) SetVerificationStatus(st string) {
	d.VerificationStatus = st
}
func (d *Deliverer) SetDeliveryProgress(progress bool) {
	d.DeliveryInProgress = progress
}
