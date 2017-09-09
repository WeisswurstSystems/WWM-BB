package authenticate

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type AuthenticateUseCase interface {
	Authenticate(user.Login) (user.User, error)
}
type Interactor struct {
	user.ReadStore
}

var (
	ErrNotAuthenticated   = wwm.Error{"Invalid Mail or Password", http.StatusUnauthorized}
	ErrNotFullyRegistered = wwm.Error{"User didn't finish his registration process.", http.StatusUnauthorized}
)

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

func authenticated(user user.User, l user.Login) bool {
	return user.RegistrationID == "" && l == user.Login
}
