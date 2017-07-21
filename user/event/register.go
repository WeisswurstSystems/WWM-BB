package event

import (
	"fmt"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/go-errors/errors"
	"net/http"
)

type Register struct {
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	MailEnabled bool   `json:"mailEnabled"`
}

func (e Register) Execute() error {
	if err := e.validate(); err != nil {
		return err
	}

	if err := isFree(e.Mail); err != nil {
		return err
	}

	u := buildUser(e)
	u, err := store.Save(u)
	if err != nil {
		return errors.New("Failed to save the User.")
	}
	mail.SendRegistrationMail(u.RegistrationID, u.Mail)
	return nil
}

func isFree(mail string) error {
	has, err := store.Has(mail)
	if err != nil {
		return err
	}
	if has {
		message := fmt.Sprintf("User with Mail %v already exists.", mail)
		return &EventError{message, http.StatusConflict}
	}
	return nil
}

func (e Register) validate() error {
	if e.Mail == "" {
		return &EventError{"Missing filed: mail", http.StatusBadRequest}
	}
	if e.Password == "" {
		return &EventError{"Missing filed: password", http.StatusBadRequest}
	}
	return nil
}

func buildUser(e Register) user.User {
	uid := util.GetUID(60)
	return user.User{
		Mail:           e.Mail,
		Password:       e.Password,
		RegistrationID: uid,
		Roles:          []string{"user"},
		MailEnabled:    e.MailEnabled,
	}
}
