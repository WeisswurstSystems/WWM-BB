package service

import (
	"encoding/json"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
)

func Read(w http.ResponseWriter, req *http.Request) {
	results, err := store.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var js []byte
	js, err = json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func Register(w http.ResponseWriter, req *http.Request) {

	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var requestUser user.RegisterUser

	err := json.NewDecoder(req.Body).Decode(&requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if requestUser.Mail == "" {
		http.Error(w, "Missing field: mail", http.StatusBadRequest)
		return
	}

	if requestUser.Password == "" {
		http.Error(w, "Missing field: password", http.StatusBadRequest)
		return
	}

	has, _ := store.Has(requestUser.Mail)
	if has {
		http.Error(w, "User with Mail "+requestUser.Mail+" already exists.", http.StatusConflict)
		return
	}

	registerUser := user.User{
		Mail:        requestUser.Mail,
		Password:    requestUser.Password,
		Roles:       []string{"user"},
		MailEnabled: requestUser.MailEnabled,
	}

	err = store.Save(registerUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(registerUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}
