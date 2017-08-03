package security

import (
	"fmt"
	"log"
	"net/http"

	"errors"
	meetingStore "github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	userStore "github.com/WeisswurstSystems/WWM-BB/user/store"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/gorilla/mux"
)

//TODO
func MeetingAuthenticationHandler(next http.HandlerFunc) http.HandlerFunc {
	LOG_TAG := "[MeetingAuthenticationHandler]"
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		result, _ := meetingStore.FindOne(vars["meetingId"])

		// if he is the creator of the meeting...
		uname, _, _ := r.BasicAuth()
		if result.Creator == uname {
			log.Printf("%v, User %v is owner of meeting %v.", LOG_TAG, uname, result.ID)
			next(w, r)
			return
		}

		// ... or if he has admin rights
		findByUserMail, _ := userStore.FindByMail(uname)
		if util.Contains(findByUserMail.Roles, "admin") {
			next(w, r)
			return
		}

		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized - Your are not the Owner of this meeting.\n"))
	}
}

//func OrderAuthenticationHandler(next http.HandlerFunc, meetingid, order)
