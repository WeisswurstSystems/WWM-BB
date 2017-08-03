package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddMeetingRoutes(r *mux.Router) {
	r.Handle("/", wwm.Handler(Query.FindAll)).Methods("GET")
	r.Handle("/{meetingId}", wwm.Handler(Query.FindByID)).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()
	do.Handle("/closeMeeting", wwm.Handler(Command.CloseMeeting)).Methods("POST")
	do.Handle("/createMeeting", wwm.Handler(Command.CreateMeeting)).Methods("POST")
	do.Handle("/putProduct", wwm.Handler(Command.PutProduct)).Methods("POST")
	do.Handle("/removeProduct", wwm.Handler(Command.RemoveProduct)).Methods("POST")
	do.Handle("/setBuyer", wwm.Handler(Command.SetBuyer)).Methods("POST")
	do.Handle("/setPlace", wwm.Handler(Command.SetPlace)).Methods("POST")
}
