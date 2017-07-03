package security

import (
	"fmt"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/user/store"
)

func DefAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkBasicAuth(r) {
			next(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	}
}

//func MeetingAuthenticationHandler(next http.HandlerFunc), meetingid)

//func OrderAuthenticationHandler(next http.HandlerFunc, meetingid, order)

func checkBasicAuth(r *http.Request) bool {
	usermail, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	findByUserMail, err := store.FindByMail(usermail)

	if err != nil {
		return false
	}

	return password == findByUserMail.Password
}
