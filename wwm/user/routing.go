package user

import "github.com/gorilla/mux"

func AddUserRoutes(r *mux.Router) {
	r.HandleFunc("/", Middleware.Authenticated(Query.Read)).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()
	do.HandleFunc("/register", Command.Register).Methods("POST")
	do.HandleFunc("/activate", Command.Activate).Methods("POST")
}
