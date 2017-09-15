package deleteAccount

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

type DeleteAccountUseCase interface {
	DeleteAccount(Request) error
}

type Request struct {
	Login user.Login `json:"login"`
}

type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

func (i *Interactor) DeleteAccount(request Request) error {
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
