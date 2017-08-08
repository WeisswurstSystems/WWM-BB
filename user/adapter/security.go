package adapter

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
		return user.User{}, ErrNoBasicAuthenticationGiven
	}

	u, err := ch.UserStore.FindByMail(mail)
	if err != nil {
		return user.User{}, err
	}
	if !u.IsRegistered() {
		return user.User{}, ErrNotFullyRegistered
	}
	if !u.Authenticate(password) {
		return user.User{}, ErrNotValid
	}
	return u, nil
}

func (ch *MiddlewareHandler) WithRoles(roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := ch.Identity(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		has := make(map[string]bool)
		for _, role := range u.Roles {
			has[role] = true
		}
		for _, role := range roles {
			_, ok := has[role]
			if !ok {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (ch *MiddlewareHandler) Authenticated(next http.Handler) http.Handler {
	LOG_TAG := "[Authenticated]"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := ch.Identity(r)

		switch err {
		case nil:
		case ErrNotFullyRegistered:
			log.Printf("%v User didn't finish his registration process.", LOG_TAG)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Please finish your registration before using this app."))
			w.WriteHeader(401)
			w.Write([]byte("401 Unregistered\n"))
			return
		default:
			log.Printf("%v User: wrong mail/password or not registered.", LOG_TAG)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "Your are not authenticated. Please sign in!"))
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized - Please login to access this resource.\n"))
		}

		log.Printf("%v User %v authorized for request to %v.", LOG_TAG, u.Mail, r.URL)
		next.ServeHTTP(w, r)
		return
	})
}
