package webhandler

import (
	"fmt"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/go-errors/errors"
	"log"
	"net/http"
)

var (
	ErrNotFullyRegistered         = errors.New("User didn't finish his registration process.")
	ErrNoBasicAuthenticationGiven = errors.New("No Basic Authentication was set.")
)

func (ch *MiddleWareHandler) Identity(r *http.Request) (user.User, error) {
	mail, password, ok := r.BasicAuth()
	if !ok {
		return nil, ErrNoBasicAuthenticationGiven
	}

	u, err := ch.UserStore.FindByMail(mail)
	if err != nil {
		return nil, err
	}

	ok = u.Authenticate(password)
	if !ok {
		return nil, errors.New("User could not be authenticated")
	}
	return u, nil
}

func (ch *MiddleWareHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	LOG_TAG := "[Authenticated]"
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := ch.Identity(r)
		if u.RegistrationID != nil {
			log.Printf("%v User %v didn't finish his registration process.", LOG_TAG, uname)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Please finish your registration before using this app."))
			w.WriteHeader(401)
			w.Write([]byte("401 Unregistered\n"))
			return
		}

		if u != nil {
			log.Printf("%v User %v authorized for request to %v.", LOG_TAG, uname, r.URL)
			next(w, r)
			return
		}

		log.Printf("%v User %v: wrong password or not registered.", LOG_TAG, uname)
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized - Please login to access this resource.\n"))
	}
}
