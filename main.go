package main

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm/construction"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	userRouter := router.PathPrefix("/users").Subrouter()
	construction.AddUserRoutes(userRouter)

	meetingRouter := router.PathPrefix("/meetings").Subrouter()
	construction.AddMeetingRoutes(meetingRouter)

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
