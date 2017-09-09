package register

import (
	"fmt"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type RegisterUseCase interface {
	Register(Request) error
}
type Interactor struct {
	user.Store
	MailService mail.Service
}

type Request struct {
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	PaypalLink	string `json:"paypalLink"`
	MailEnabled bool   `json:"mailEnabled"`
}

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

func (interactor *Interactor) isFree(mail string) error {
	_, err := interactor.Store.FindByMail(mail)
	if err == user.ErrNotFound {
		return nil
	}
	message := fmt.Sprintf("User with Mail %v already exists.", mail)
	return &wwm.Error{message, http.StatusConflict}
}

func buildUser(req Request) user.User {
	uid := util.GetUID(60)
	return user.User{
		Login:          user.Login{req.Mail, req.Password},
		PaypalLink:		req.PaypalLink,
		RegistrationID: uid,
		Roles:          []string{"user"},
		MailEnabled:    req.MailEnabled,
	}
}

func (e Request) Validate() error {
	if e.Mail == "" {
		return &wwm.Error{"Missing filed: mail", http.StatusBadRequest}
	}
	if e.Password == "" {
		return &wwm.Error{"Missing filed: password", http.StatusBadRequest}
	}
	return nil
}
