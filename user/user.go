package user

type User struct {
	UserID        string         `json:"userID"`
	Mail          string         `json:"mail"`
	Roles         []string       `json:"roles"`
	DefaultOrders map[string]int `json:"defaultOrders"`
	MailEnabled   bool           `json:"mailEnabled"`
}
