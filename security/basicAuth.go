package security

import (
	"fmt"
	"net/http"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/user"

	"gopkg.in/mgo.v2/bson"
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

	var findByUserMail user.User
	err := database.Users.Find(bson.M{"mail": u}).One(&findByUserMail)

	if err != nil {

		return false
	}

	return p == findByUserMail.Password
}
