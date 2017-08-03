package main

import (
	meeting "github.com/WeisswurstSystems/WWM-BB/meeting/construction"
	user "github.com/WeisswurstSystems/WWM-BB/user/construction"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	userRouter := router.PathPrefix("/users").Subrouter()
	user.AddUserRoutes(userRouter)

	meetingRouter := router.PathPrefix("/meetings").Subrouter()
	meeting.AddMeetingRoutes(meetingRouter)

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
