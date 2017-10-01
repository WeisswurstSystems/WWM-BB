package user

import (
	"fmt"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// Login describes the information necessary for a user.
type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

// ToString creates a console representation, that does hide the password.
func (l Login) ToString() string {
	return fmt.Sprintf("Login[%v, %v]", l.Mail, "...")
}

// User describes the aggregate of a user.
type User struct {
	Login          `json:"login"`
	PayPal         PayPal         `json:"payPal"`
	RegistrationID string         `json:"registrationID"`
	Roles          []string       `json:"roles"`
	DefaultOrders  map[string]int `json:"defaultOrders"`
	MailEnabled    bool           `json:"mailEnabled"`
}

// ReadStore can be used to do query operations for a User. The Mail of Login is used as id.
type ReadStore interface {
	FindByMail(string) (User, error)
	FindAll() ([]User, error)
	FindByRegistrationID(registrationID string) (User, error)
}

// WriteStore saves or removes a user.
type WriteStore interface {
	Save(User) error
	RemoveByMail(string) error
}

// Store can read and write for a user.
type Store interface {
	ReadStore
	WriteStore
}

var (
	// ErrNotFound if the user does not exist in a store.
	ErrNotFound = wwm.Error{"User does not exist", http.StatusNotFound}
)

// IsRegistered returns if the user is fully registered.
func (user User) IsRegistered() bool {
	return user.RegistrationID == ""
}

var (
	// Hans is a pseudo user with admin role
	Hans = User{
		Login: Login{
			Mail:     "hans@gmail.com",
			Password: "hansistsogeil",
		},
		Roles: []string{"admin"},
		PayPal: PayPal{
			MeLink: "http://www.paypal.me/hans",
		},
	}
)
