package command

import (
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type Interactor interface {
	register.RegisterUseCase
	activate.ActivateUseCase
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
