package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/weisswurstSystems/WWM-BB/security"
)

func GetHelloWorldEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

func main() {
	router := mux.NewRouter()
	// unsecured endpoints
	router.HandleFunc("/", GetHelloWorldEndpoint).Methods("GET")

	// secured endpoints
	router.HandleFunc("/secured", security.DefaultAuthenticationHandler("Please login to see the details of a person", GetHelloWorldEndpoint)).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Printf("Starting on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
