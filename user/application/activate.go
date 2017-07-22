package application

import (
	"log"
	"net/http"
)

func (i *Interactor) Activate(e Activate) error {
	result, err := i.UserStore.FindByRegistrationID(e.RegistrationID)
	if err != nil {
		return &Error{"User with this ID was already activated", http.StatusBadRequest}
	}
	result.RegistrationID = ""
	err = i.UserStore.Save(result)
	if err != nil {
		return err
	}
	log.Printf("Successfully unlocked user %v", result.Mail)
	return nil
}
