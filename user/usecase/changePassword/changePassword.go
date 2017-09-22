package changePassword

import (
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// ChangePasswordUseCase changes the password of the user.
//
// The password must be repeated to ensure no type error.
type ChangePasswordUseCase interface {
	ChangePassword(Request) error
}

// Request for changing a password.
type Request struct {
	Login            user.Login `json:"login"`
	Password         string     `json:"password"`
	PasswordRepeated string     `json:"passwordRepeated"`
}

// Interactor for doing the actual usecase logic.
type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

var (
	// ErrPasswordsNotEqual if the passwords are not equal
	ErrPasswordsNotEqual = wwm.Error{
		Message: "Passwords are not the same",
		Code:    http.StatusBadRequest,
	}
)

// ChangePassword changes the password of the User (by login).
func (i Interactor) ChangePassword(request Request) error {
	u, err := i.Authenticate(request.Login)
	if err != nil {
		return err
	}

	if request.Password != request.PasswordRepeated {
		return ErrPasswordsNotEqual
	}

	u.Password = request.Password
	err = i.Save(u)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("Changed password of user %v", u.Mail)
	return nil
}
