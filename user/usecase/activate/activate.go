package activate

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type ActivateUseCase interface {
	Activate(Request) error
}

type Request struct {
	RegistrationID string `json:"registrationID"`
}
type Interactor struct {
	user.Store
}

var (
	ErrAlreadyActivated = wwm.Error{"User with this ID was already activated", http.StatusBadRequest}
)

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
