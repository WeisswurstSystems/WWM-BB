package changePassword

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
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

// ChangePassword changes the password of the User (by login).
func (i Interactor) ChangePassword(request Request) error {
	u, err := i.Authenticate(request.Login)
	if err != nil {
		return err
	}

	u.Password = request.Password
	err = i.Save(u)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("Changed password of user %v", u.Mail)
	return nil
}
