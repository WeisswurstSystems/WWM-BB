package handler

import (
	"encoding/json"
	"github.com/WeisswurstSystems/WWM-BB/user/event"
	"net/http"
)

func Register(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var register event.Register
	err := json.NewDecoder(req.Body).Decode(&register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = register.Execute(); err != nil {
		if ee, ok := err.(*event.EventError); ok {
			http.Error(w, ee.Message, ee.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Activate(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	var activate event.Activate
	err := json.NewDecoder(req.Body).Decode(&activate)

	if err = activate.Execute(); err != nil {
		if ee, ok := err.(*event.EventError); ok {
			http.Error(w, ee.Message, ee.Code)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, req, "http://www.google.com", 301)
}
