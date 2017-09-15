package command

import (
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/changePassword"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/deleteAccount"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/setUpPayPal"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

type Interactor interface {
	register.RegisterUseCase
	activate.ActivateUseCase
	setUpPayPal.SetUpPayPalUseCase
	changePassword.ChangePasswordUseCase
	deleteAccount.DeleteAccountUseCase
}

type CommandHandler struct {
	Interactor
}

func (ch *CommandHandler) Register(w http.ResponseWriter, req *http.Request) error {
	var e register.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.Register(e)
}

func (ch *CommandHandler) SetUpPayPal(w http.ResponseWriter, req *http.Request) error {
	var e setUpPayPal.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetUpPayPal(e)
}

func (ch *CommandHandler) Activate(w http.ResponseWriter, req *http.Request) error {
	var e activate.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	err = ch.Interactor.Activate(e)
	if err != nil {
		return err
	}
	http.Redirect(w, req, "http://www.google.com", 301)
	return nil
}

func (ch *CommandHandler) ChangePassword(w http.ResponseWriter, req *http.Request) error {
	var e changePassword.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.ChangePassword(e)
}

func (ch *CommandHandler) DeleteAccount(w http.ResponseWriter, req *http.Request) error {
	var e deleteAccount.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.DeleteAccount(e)
}
