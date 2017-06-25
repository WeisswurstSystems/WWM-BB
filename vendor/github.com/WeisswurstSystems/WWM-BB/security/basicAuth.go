package security

import (
	"fmt"
	"net/http"
)

func DefaultAuthenticationHandler(realm string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkBasicAuth(r) {
			next(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	}
}

//func MeetingAuthenticationHandler(next http.HandlerFunc), meetingid)

//func OrderAuthenticationHandler(next http.HandlerFunc, meetingid, order)

func checkBasicAuth(r *http.Request) bool {
	u, p, ok := r.BasicAuth()
	if !ok {
		return false
	}
	//TODO Datenbankabfrage
	return u == "user" && p == "user"
}
