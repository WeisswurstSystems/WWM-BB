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

// Interactor describes the group of methods needed for this command handler.
type Interactor interface {
	register.RegisterUseCase
	activate.ActivateUseCase
	setUpPayPal.SetUpPayPalUseCase
	changePassword.ChangePasswordUseCase
	deleteAccount.DeleteAccountUseCase
}

// Handler holds possible Commands (from CQRS) for users.
type Handler struct {
	Interactor
}

// Register a new user by json
func (ch *Handler) Register(w http.ResponseWriter, req *http.Request) error {
	var e register.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.Register(e)
}

// SetUpPayPal for this user
func (ch *Handler) SetUpPayPal(w http.ResponseWriter, req *http.Request) error {
	var e setUpPayPal.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetUpPayPal(e)
}

// Activate a user
func (ch *Handler) Activate(w http.ResponseWriter, req *http.Request) error {
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

// ChangePassword of the user
func (ch *Handler) ChangePassword(w http.ResponseWriter, req *http.Request) error {
	var e changePassword.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.ChangePassword(e)
}

// DeleteAccount deletes the User object for the user
func (ch *Handler) DeleteAccount(w http.ResponseWriter, req *http.Request) error {
	var e deleteAccount.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.DeleteAccount(e)
}
