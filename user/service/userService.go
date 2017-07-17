package service

import (
	"encoding/json"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/gorilla/mux"
	"log"
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

	uid := util.GetUID(60)

	registerUser := user.User{
		Mail:        requestUser.Mail,
		Password:    requestUser.Password,
		RegistrationID: uid,
		Roles:       []string{"user"},
		MailEnabled: requestUser.MailEnabled,
	}

	mail.SendRegistrationMail(uid, requestUser.Mail)

	registerUser, err = store.Save(registerUser)

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

func Activate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindByRegID(vars["registrationID"])
	if err != nil {
		http.Error(w, "User with this ID was already activated", http.StatusBadRequest)
		return
	}
	result.RegistrationID = ""
	err = store.Update(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Successfully unlocked user %v", result.Mail)
	http.Redirect(w, req, "http://www.google.com", 301)
}