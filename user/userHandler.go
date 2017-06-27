package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/util"

	"gopkg.in/mgo.v2/bson"
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

	if req.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var requestUser RegisterUser

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

	var findByUserMail []User
	err = database.Users.Find(bson.M{"mail": requestUser.Mail}).All(&findByUserMail)
	fmt.Println(findByUserMail)

	if len(findByUserMail) > 0 {
		http.Error(w, "User with Mail "+requestUser.Mail+" already exists.", http.StatusConflict)
		return
	}

	registerUser := User{
		UserID:      util.GetUID(12),
		Mail:        requestUser.Mail,
		Password:    requestUser.Password,
		Roles:       []string{"user"},
		MailEnabled: requestUser.MailEnabled,
	}

	err = database.Users.Insert(&registerUser)
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
