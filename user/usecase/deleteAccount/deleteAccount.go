package deleteAccount

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

// DeleteAccountUseCase can remove an account from the application.
type DeleteAccountUseCase interface {
	DeleteAccount(Request) error
}

// Request for deleting an account. The account is given by the Login Mail.
type Request struct {
	Login user.Login `json:"login"`
}

// Interactor can do the actual delete account logic.
type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

// DeleteAccount deletes the account of a user.
func (i Interactor) DeleteAccount(request Request) error {
	u, err := i.Authenticate(request.Login)
	if err != nil {
		return err
	}

	err = i.RemoveByMail(u.Mail)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("Deleted account of user %v", u.Mail)
	return nil
}
