package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/security"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	// Opening Database Connection...
	database.Init()
	defer database.DBSession.Close()

	router := mux.NewRouter()
	// unsecured endpoints
	router.HandleFunc("/users", security.DefaultAuthenticationHandler("Please login to see all users.", userHandler.Read)).Methods("GET")
	router.HandleFunc("/users", userHandler.Register).Methods("POST")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Printf("Starting on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
