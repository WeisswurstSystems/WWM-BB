package webhandler

import (
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/user/event"
	"net/http"
)

func (ch *CommandHandler) Register(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var e event.Register
	err := json.NewDecoder(req.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = ch.UserInteractor.Register(e); err != nil {
		if ee, ok := err.(*event.Error); ok {
			http.Error(w, ee.Message, ee.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ch *CommandHandler) Activate(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	var e event.Activate
	err := json.NewDecoder(req.Body).Decode(&e)

	if err = ch.UserInteractor.Activate(e); err != nil {
		if ee, ok := err.(*event.Error); ok {
			http.Error(w, ee.Message, ee.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, req, "http://www.google.com", 301)
}
