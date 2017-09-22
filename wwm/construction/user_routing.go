package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddUserRoutes(r *mux.Router) {
	r.Handle("/", wwm.Handler(UserQuery.FindAll)).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()
	do.Handle("/register", wwm.Handler(UserCommand.Register)).Methods("POST")
	do.Handle("/activate", wwm.Handler(UserCommand.Activate)).Methods("POST")
	do.Handle("/setUpPayPal", wwm.Handler(UserCommand.SetUpPayPal)).Methods("POST")
	do.Handle("/changePassword", wwm.Handler(UserCommand.ChangePassword)).Methods("POST")
	do.Handle("/deleteAccount", wwm.Handler(UserCommand.DeleteAccount)).Methods("POST")

}
