package event

import (
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"log"
	"net/http"
)

type Activate struct {
	RegistrationID string `json:"registrationID"`
}

func (e Activate) Execute() error {
	result, err := store.FindByRegID(e.RegistrationID)
	if err != nil {
		return &EventError{"User with this ID was already activated", http.StatusBadRequest}
	}
	result.RegistrationID = ""
	err = store.Update(result)
	if err != nil {
		return err
	}
	log.Printf("Successfully unlocked user %v", result.Mail)
	return nil
}
