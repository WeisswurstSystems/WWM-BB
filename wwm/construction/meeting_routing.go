package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddMeetingRoutes(r *mux.Router) {
	r.Handle("/", wwm.Handler(MeetingQuery.FindAll)).Methods("GET")
	r.Handle("/{meetingId}", wwm.Handler(MeetingQuery.FindByID)).Methods("GET")

	do := r.PathPrefix("/do").Subrouter()

	do.Handle("/closeMeeting",
		UserMiddleware.WithRoles([]string{"ADMIN"},
			wwm.Handler(MeetingCommand.CloseMeeting),
		),
	).Methods("POST")

	do.Handle("/createMeeting", UserMiddleware.WithRoles([]string{"ADMIN"}, wwm.Handler(MeetingCommand.CreateMeeting))).Methods("POST")
	do.Handle("/putProduct", wwm.Handler(MeetingCommand.PutProduct)).Methods("POST")
	do.Handle("/removeProduct", wwm.Handler(MeetingCommand.RemoveProduct)).Methods("POST")
	do.Handle("/setBuyer", wwm.Handler(MeetingCommand.SetBuyer)).Methods("POST")
	do.Handle("/setPlace", wwm.Handler(MeetingCommand.SetPlace)).Methods("POST")
}
