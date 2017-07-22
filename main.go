package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	meetingService "github.com/WeisswurstSystems/WWM-BB/meeting/service"
	"github.com/WeisswurstSystems/WWM-BB/security"
	userHandler "github.com/WeisswurstSystems/WWM-BB/user/webhandler"
	"github.com/gorilla/mux"
)

func main() {
	mailService := mail.NewSMTPService()
	router := mux.NewRouter()

	// unsecured endpoints
	router.HandleFunc("/users/do/register", userHandler.RegisterHandler(&mailService)).Methods("POST")
	router.HandleFunc("/users/do/activate", userHandler.Activate).Methods("POST")
	router.HandleFunc("/meetings", meetingService.ReadAll).Methods("GET")
	router.HandleFunc("/meetings/{meetingId}", meetingService.ReadSingle).Methods("GET")

	// secured endpoints
	router.HandleFunc("/users", security.DefAuth(userHandler.Read)).Methods("GET")
	router.HandleFunc("/meetings", security.DefAuth(meetingService.Create)).Methods("POST")

	// secured and only meeting owner endpoints
	router.HandleFunc("/meetings/{meetingId}/setPlace", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetPlace))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/setDate", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetDate))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/setBuyer", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetBuyer))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/addProduct", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.AddProduct))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/changeProduct", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.ChangeProduct))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/closeMeeting", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.CloseMeeting))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/openMeeting", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.OpenMeeting))).Methods("POST")

	// Let's go!
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Printf("Starting on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
