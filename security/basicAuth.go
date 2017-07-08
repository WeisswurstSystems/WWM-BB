package security

import (
	"fmt"
	"net/http"
	"log"

	userStore "github.com/WeisswurstSystems/WWM-BB/user/store"
	meetingStore "github.com/WeisswurstSystems/WWM-BB/meeting/store"
	"github.com/gorilla/mux"
)

func DefAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uname, _, _ := r.BasicAuth()
		if checkBasicAuth(r) {
			log.Printf("User %v authorized.", uname);
			next(w, r)
			return
		}

		log.Printf("User %v: wrong password or not registered.", uname)
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
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
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		result, _ := meetingStore.FindOne(vars["meetingId"])

		if uname,_,_ :=  r.BasicAuth(); result.Creator == uname {
			next(w, r)
			return
		}

		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	}
}

//func OrderAuthenticationHandler(next http.HandlerFunc, meetingid, order)

func checkBasicAuth(r *http.Request) bool {
	usermail, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	findByUserMail, err := userStore.FindByMail(usermail)

	if err != nil {
		return false
	}

	return password == findByUserMail.Password
}
