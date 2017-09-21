package setUpPayPal

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

// SetUpPayPalUseCase can set up or change the paypal configuration for the user.
type SetUpPayPalUseCase interface {
	SetUpPayPal(Request) error
}

// Interactor can do the logic for setting up the PayPal.
type Interactor struct {
	user.Store
	authenticate.AuthenticateUseCase
}

// Request contains the PayPal data.
type Request struct {
	PayPal user.PayPal `json:"payPal"`
	Login  user.Login  `json:"login"`
}

// SetUpPayPal sets the paypal configuration for a user.
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
