package register

import (
	"fmt"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// RegisterUseCase is there to register a new User to the application.
type RegisterUseCase interface {
	Register(Request) error
}

// Interactor can do the register.
type Interactor struct {
	user.Store
	MailService mail.Service
}

// Request for Registering a new user. Mailing can be enabled.
type Request struct {
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	MailEnabled bool   `json:"mailEnabled"`
}

// Register creates the new user, if it does not already exist. It sends a mail with the RegistrationID.
func (i Interactor) Register(req Request) error {
	if err := req.Validate(); err != nil {
		return err
	}

	if err := i.isFree(req.Mail); err != nil {
		return err
	}

	u := buildUser(req)
	err := i.Store.Save(u)
	if err != nil {
		return err
	}

	m := mail.NewRegistrationMail(u.RegistrationID, u.Mail)
	return i.MailService.Send(m)
}

// isFree checks if the mail is not already registered in the application.
func (interactor *Interactor) isFree(mail string) error {
	_, err := interactor.Store.FindByMail(mail)
	if err == user.ErrNotFound {
		return nil
	}
	message := fmt.Sprintf("User with Mail %v already exists.", mail)
	return &wwm.Error{message, http.StatusConflict}
}

// buildUser Builds a User with a random UUID and default roles.
func buildUser(req Request) user.User {
	uid := util.GetUID(60)
	return user.User{
		Login:          user.Login{req.Mail, req.Password},
		RegistrationID: uid,
		Roles:          []string{"user"},
		MailEnabled:    req.MailEnabled,
	}
}

// Validate checks if the Request is valid.
func (e Request) Validate() error {
	if e.Mail == "" {
		return &wwm.Error{"Missing filed: mail", http.StatusBadRequest}
	}
	if e.Password == "" {
		return &wwm.Error{"Missing filed: password", http.StatusBadRequest}
	}
	return nil
}
