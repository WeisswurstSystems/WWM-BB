package setUpPayPal

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type SetUpPayPalUseCase interface {
	SetUpPayPal(Request) error
}
type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	PayPal user.PayPal `json:"payPal"`
	Login  user.Login  `json:"login"`
}

func (i Interactor) SetUpPayPal(req Request) error {
	u, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	u.PayPal = req.PayPal

	err = i.Save(u)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("did %v", req)
	return nil
}
