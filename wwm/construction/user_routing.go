package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddUserRoutes(r *mux.Router) {
	r.Handle("/", UserMiddleware.Authenticated(wwm.Handler(UserQuery.FindAll))).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()
	do.Handle("/register", wwm.Handler(UserCommand.Register)).Methods("POST")
	do.Handle("/activate", wwm.Handler(UserCommand.Activate)).Methods("POST")
}
