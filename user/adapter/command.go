package adapter

import (
	"github.com/WeisswurstSystems/WWM-BB/user/application"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

func (ch *CommandHandler) Register(w http.ResponseWriter, req *http.Request) error {
	var e application.Register
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.UserInteractor.Register(e)
}

func (ch *CommandHandler) Activate(w http.ResponseWriter, req *http.Request) error {
	var e application.Register
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	err = ch.UserInteractor.Register(e)
	if err != nil {
		return err
	}
	http.Redirect(w, req, "http://www.google.com", 301)
	return nil
}
