package user

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
type User struct {
	Login          `json:"login"`
	PayPal         PayPal         `json:"payPal"`
	RegistrationID string         `json:"registrationID"`
	Roles          []string       `json:"roles"`
	DefaultOrders  map[string]int `json:"defaultOrders"`
	MailEnabled    bool           `json:"mailEnabled"`
}

type ReadStore interface {
	FindByMail(string) (User, error)
	FindAll() ([]User, error)
	FindByRegistrationID(registrationID string) (User, error)
}
type WriteStore interface {
	Save(User) error
}

type Store interface {
	ReadStore
	WriteStore
}

var (
	ErrNotFound = wwm.Error{"User does not exist", http.StatusNotFound}
)

func (user User) IsRegistered() bool {
	return user.RegistrationID == ""
}
