package user

type User struct {
	UserID        string         `json:"userID"`
	Mail          string         `json:"mail"`
	Password      string         `json:"-"`
	Roles         []string       `json:"roles"`
	DefaultOrders map[string]int `json:"defaultOrders"`
	MailEnabled   bool           `json:"mailEnabled"`
}

type RegisterUser struct {
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	MailEnabled bool   `json:"mailEnabled"`
}
