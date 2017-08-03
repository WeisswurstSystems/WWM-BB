package driver

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddUserRoutes(r *mux.Router) {
	r.Handle("/", Middleware.Authenticated(wwm.Handler(Query.FindAll))).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()
	do.Handle("/register", wwm.Handler(Command.Register)).Methods("POST")
	do.Handle("/activate", wwm.Handler(Command.Activate)).Methods("POST")
}
