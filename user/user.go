package user

type User struct {
	Mail           string         `json:"mail"`
	Password       string         `json:"-"`
	RegistrationID string         `json:"-"`
	Roles          []string       `json:"roles"`
	DefaultOrders  map[string]int `json:"defaultOrders"`
	MailEnabled    bool           `json:"mailEnabled"`
}

type Store interface {
	HasByMail(mail string) (bool, error)
	Save(User) error
	FindByMail(string) (User, error)
	FindAll() ([]User, error)
	FindByRegistrationID(registrationID string) (User, error)
}
