package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"github.com/gorilla/mux"
)

func AddMeetingRoutes(r *mux.Router) {
	r.Handle("/", wwm.Handler(MeetingQuery.FindAll)).Methods("GET")
	r.Handle("/{meetingId}", wwm.Handler(MeetingQuery.FindByID)).Methods("GET")

	// Nicht meeting-spezifische Aktionen
	doNoID := r.PathPrefix("/do").Subrouter()
	doNoID.Handle("/createMeeting", wwm.Handler(MeetingCommand.CreateMeeting)).Methods("POST")

	// Meeting-spezifische Aktionen
	do := r.PathPrefix("/{meetingId}/do").Subrouter()
	do.Handle("/invite", wwm.Handler(MeetingCommand.Invite)).Methods("POST")
	do.Handle("/notify", wwm.Handler(MeetingCommand.Notify)).Methods("POST")
	do.Handle("/toggleOrderPayed", wwm.Handler(MeetingCommand.ToggleOrderPayed)).Methods("POST")
	do.Handle("/closeMeeting", wwm.Handler(MeetingCommand.CloseMeeting)).Methods("POST")
	do.Handle("/putProduct", wwm.Handler(MeetingCommand.PutProduct)).Methods("POST")
	do.Handle("/removeProduct", wwm.Handler(MeetingCommand.RemoveProduct)).Methods("POST")
	do.Handle("/setBuyer", wwm.Handler(MeetingCommand.SetBuyer)).Methods("POST")
	do.Handle("/setPlace", wwm.Handler(MeetingCommand.SetPlace)).Methods("POST")
}
