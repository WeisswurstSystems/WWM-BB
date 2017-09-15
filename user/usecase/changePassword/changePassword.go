package changePassword

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

type ChangePasswordUseCase interface {
	Activate(Request) error
}

type Request struct {
	Login    user.Login `json:"login"`
	Password string     `json:"password"`
}

type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

func (i *Interactor) Activate(request Request) error {
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
