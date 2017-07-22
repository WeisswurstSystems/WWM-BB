package middleware

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
	ErrNotValid                   = errors.New("Username or password is wrong")
)

func (ch *MiddlewareHandler) Identity(r *http.Request) (user.User, error) {
	mail, password, ok := r.BasicAuth()
	if !ok {
		return nil, ErrNoBasicAuthenticationGiven
	}

	u, err := ch.UserStore.FindByMail(mail)
	if err != nil {
		return nil, err
	}
	if !u.IsRegistered() {
		return nil, ErrNotFullyRegistered
	}
	if !u.Authenticate(password) {
		return nil, ErrNotValid
	}
	return u, nil
}

func (ch *MiddlewareHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	LOG_TAG := "[Authenticated]"
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := ch.Identity(r)

		switch err {
		case ErrNotFullyRegistered:
			log.Printf("%v User didn't finish his registration process.", LOG_TAG)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Please finish your registration before using this app."))
			w.WriteHeader(401)
			w.Write([]byte("401 Unregistered\n"))
			return
		case err != nil:
			log.Printf("%v User: wrong mail/password or not registered.", LOG_TAG)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized - Please login to access this resource.\n"))
		}

		log.Printf("%v User %v authorized for request to %v.", LOG_TAG, u.Mail, r.URL)
		next(w, r)
		return
	}
}
