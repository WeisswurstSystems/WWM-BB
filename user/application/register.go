package application

import (
	"fmt"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"net/http"
)

func (interactor *Interactor) Register(e Register) error {
	if err := e.Validate(); err != nil {
		return err
	}

	if err := interactor.isFree(e.Mail); err != nil {
		return err
	}

	u := buildUser(e)
	err := interactor.UserStore.Save(u)
	if err != nil {
		return err
	}

	m := mail.NewRegistrationMail(u.RegistrationID, u.Mail)
	interactor.MailService.Send(m)
	return nil
}

func (interactor *Interactor) isFree(mail string) error {
	user, err := interactor.UserStore.HasByMail(mail)
	if err != nil {
		return err
	}
	if user {
		message := fmt.Sprintf("User with Mail %v already exists.", mail)
		return &Error{message, http.StatusConflict}
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
