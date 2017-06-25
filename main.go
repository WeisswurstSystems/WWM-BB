package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/security"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

//<----- Testcode for Database connection //
type Person struct {
	Name  string
	Phone string
}

func GetWriteEndpoint(w http.ResponseWriter, req *http.Request) {
	person := Person{"Ale", "+55 53 8116"}

	err := database.PeopleCollection.Insert(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - " + err.Error()))
		return
	}

	js, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func GetReadEndpoint(w http.ResponseWriter, req *http.Request) {
	var results []Person
	err := database.PeopleCollection.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// Testcode for Database connection ----->//

func main() {
	// Opening Database Connection...
	initDatabase()

	defer database.DBSession.Close()

	router := mux.NewRouter()
	// unsecured endpoints
	router.HandleFunc("/create", GetWriteEndpoint).Methods("GET")
	router.HandleFunc("/read", GetReadEndpoint).Methods("GET")

	// secured endpoints
	router.HandleFunc("/secured", security.DefaultAuthenticationHandler("Please login to see all persons", GetReadEndpoint)).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Printf("Starting on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}

func initDatabase() {
	username := getEnv("DB_USERNAME", "wwm")
	password := getEnv("DB_PASSWORD", "wwm")
	url := getEnv("DB_URL", "ds064649.mlab.com:64649/wwmbb-dev")
	databaseName := getEnv("DB_NAME", "wwmbb-dev")

	log.Printf("Establishing connection to database <%v> on %v", databaseName, url)

	var err error
	database.DBSession, err = mgo.Dial(fmt.Sprintf("mongodb://%v:%v@%v", username, password, url))
	if err != nil {
		panic(err)
	}
	database.DBSession.SetMode(mgo.Monotonic, true)
	database.PeopleCollection = database.DBSession.DB(databaseName).C("people")
	log.Printf("Connection succesfully established!")
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
