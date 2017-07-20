package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WeisswurstSystems/WWM-BB/database"
	meetingService "github.com/WeisswurstSystems/WWM-BB/meeting/service"
	"github.com/WeisswurstSystems/WWM-BB/security"
	userService "github.com/WeisswurstSystems/WWM-BB/user/service"
	"github.com/gorilla/mux"
)

func main() {
	// Opening Database Connection...
	database.Init()
	defer database.DBSession.Close()

	router := mux.NewRouter()

	// unsecured endpoints
	router.HandleFunc("/users", userService.Register).Methods("POST")
	router.HandleFunc("/meetings", meetingService.ReadAll).Methods("GET")
	router.HandleFunc("/meetings/{meetingId}", meetingService.ReadSingle).Methods("GET")

	// secured endpoints
	router.HandleFunc("/users", security.DefAuth(userService.Read)).Methods("GET")
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
