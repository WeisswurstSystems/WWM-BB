package activate

import (
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// ActivateUseCase can confirm the mail adress of a new registration.
type ActivateUseCase interface {
	Activate(Request) error
}

// Request contains the required RegistrationID as proof that the email was received.
type Request struct {
	RegistrationID string `json:"registrationID"`
}

// Interactor can do the Activation logic.
type Interactor struct {
	user.Store
}

var (
	// ErrAlreadyActivated if the user was already confirmed.
	ErrAlreadyActivated = wwm.Error{"User with this ID was already activated", http.StatusBadRequest}
)

// Activate checks if the RegistrationID matches and enables the account.
func (i Interactor) Activate(req Request) error {
	result, err := i.Store.FindByRegistrationID(req.RegistrationID)
	if err != nil {
		return ErrAlreadyActivated
	}
	result.RegistrationID = ""
	err = i.Store.Save(result)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("Successfully unlocked user %v", result.Mail)
	return nil
}
