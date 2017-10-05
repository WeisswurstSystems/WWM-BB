package user

// PublicUser contains all information of a user that are not private.
type PublicUser struct {
	Mail         string   `json:"mail"`
	PayPalMeLink string   `json:"payPalMeLink"`
	Roles        []string `json:"roles"`
}

// PublicUser representation of a user.
func (u User) PublicUser() (p PublicUser) {
	p.Mail = u.Login.Mail
	p.PayPalMeLink = u.PayPal.MeLink
	p.Roles = u.Roles
	return p
}
