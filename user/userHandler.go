package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/security"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func Read(w http.ResponseWriter, req *http.Request) {
	var results []User
	err := database.Users.Find(nil).All(&results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func Register(w http.ResponseWriter, req *http.Request) {

	var requestUser User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var findByUserMail []User
	err = database.User.Find(bson.M{"mail": requestUser.Mail}).All(&findByUserMail)
	fmt.Println(findByUserMail)

	if len(findByUserMail) > 0 {
		http.Error(w, "User with Mail "+requestUser.Mail+" already exists.", http.StatusConflict)
		return
	}

	registerUser := User{
		UserID:      util.GetUID,
		Mail:        requestUser.Mail,
		Password:    requestUser.Password,
		Roles:       []string{"user"},
		MailEnabled: requestUser.MailEnabled,
	}

	err := database.Users.Insert(&registerUser)
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
