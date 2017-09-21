package authenticate

import (
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// AuthenticateUseCase can authenticate a user by the combination of mail and password.
type AuthenticateUseCase interface {
	Authenticate(user.Login) (user.User, error)
}

// Interactor can authenticate a user.
type Interactor struct {
	user.ReadStore
}

var (
	// ErrNotAuthenticated if the user has an invalid mail or password.
	ErrNotAuthenticated = wwm.Error{"Invalid Mail or Password", http.StatusUnauthorized}
	// ErrNotFullyRegistered if the user has not yet confirmed his mail.
	ErrNotFullyRegistered = wwm.Error{"User didn't finish his registration process.", http.StatusUnauthorized}
)

// Authenticate takes the Login information and checks if the user is authenticated.
func (i Interactor) Authenticate(l user.Login) (user.User, error) {
	u, err := i.ReadStore.FindByMail(l.Mail)
	if err != nil {
		return user.User{}, err
	}
	if !authenticated(u, l) {
		return user.User{}, ErrNotAuthenticated
	}

	if !u.IsRegistered() {
		return user.User{}, ErrNotFullyRegistered
	}
	return u, nil
}

// authenticated checks the user against the login.
func authenticated(user user.User, l user.Login) bool {
	return user.RegistrationID == "" && l == user.Login
}
