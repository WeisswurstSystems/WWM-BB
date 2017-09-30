package query

import (
	"encoding/json"
	"net/http"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

type QueryHandler struct {
	user.Store
	authenticate.AuthenticateUseCase
}

func (ch *QueryHandler) FindAll(w http.ResponseWriter, req *http.Request) error {
	var e user.Login
	var ok bool
	e.Mail, e.Password, ok = req.BasicAuth()

	if(!ok) {
		return wwm.Error{"Unauthenticated", http.StatusUnauthorized}
	}

	_, err := ch.AuthenticateUseCase.Authenticate(e)
	if err != nil {
		return err
	}

	var results []user.User
	results, err = ch.Store.FindAll()
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(results)
}

func (ch *QueryHandler) Identity(w http.ResponseWriter, req *http.Request) error {
	var e user.Login
	e.Mail, e.Password, _ = req.BasicAuth()
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}

	results, err := ch.AuthenticateUseCase.Authenticate(e)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(results)
}
