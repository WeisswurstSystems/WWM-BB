package security

import (
	"fmt"
	"net/http"
	"log"

	userStore "github.com/WeisswurstSystems/WWM-BB/user/store"
	meetingStore "github.com/WeisswurstSystems/WWM-BB/meeting/store"
	"github.com/gorilla/mux"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"errors"
)

func DefAuth(next http.HandlerFunc) http.HandlerFunc {
	LOG_TAG := "[DefAuth]"
	return func(w http.ResponseWriter, r *http.Request) {
		uname, _, _ := r.BasicAuth()

		authenticated, err := checkBasicAuth(r)

		if(err != nil) {
			log.Printf("%v User %v didn't finish his registration process.", LOG_TAG, uname);
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Please finish your registration before using this app."))
			w.WriteHeader(401)
			w.Write([]byte("401 Unregistered\n"))
			return
		}

		if authenticated {
			log.Printf("%v User %v authorized for request to %v.", LOG_TAG, uname, r.URL);
			next(w, r)
			return
		}

		log.Printf("%v User %v: wrong password or not registered.", LOG_TAG, uname)
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized - Please login to access this resource.\n"))
	}
}

func GetCurrentUser(r *http.Request) string {
	usermail, _, ok := r.BasicAuth();
	if ok {
		return usermail
	} else {
		return ""
	}
}

func MeetingAuthenticationHandler(next http.HandlerFunc) http.HandlerFunc{
	LOG_TAG := "[MeetingAuthenticationHandler]"
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		result, _ := meetingStore.FindOne(vars["meetingId"])

		// if he is the creator of the meeting...
		uname,_,_ :=  r.BasicAuth()
		if result.Creator == uname {
			log.Printf("%v, User %v is owner of meeting %v.", LOG_TAG, uname, result.ID);
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

func checkBasicAuth(r *http.Request) (bool, error) {
	usermail, password, ok := r.BasicAuth()
	if !ok {
		return false, nil
	}

	findByUserMail, err := userStore.FindByMail(usermail)

	if err != nil {
		return false, nil
	}

	if findByUserMail.RegistrationID == "" {
		return password == findByUserMail.Password, nil
	} else {
		return false, errors.New("Please finish your registration!")
	}
}
