package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	DBSession *mgo.Session
	Users     *mgo.Collection
	Meetings  *mgo.Collection
)

func Init() {
	username := getEnv("DB_USERNAME", "wwm")
	password := getEnv("DB_PASSWORD", "wwm")
	url := getEnv("DB_URL", "ds064649.mlab.com:64649/wwmbb-dev")
	databaseName := getEnv("DB_NAME", "wwmbb-dev")

	log.Printf("Establishing connection to database <%v> on %v", databaseName, url)

	var err error
	DBSession, err = mgo.Dial(fmt.Sprintf("mongodb://%v:%v@%v", username, password, url))
	if err != nil {
		panic(err)
	}
	DBSession.SetMode(mgo.Monotonic, true)
	Users = DBSession.DB(databaseName).C("users")
	Meetings = DBSession.DB(databaseName).C("meetings")
	log.Printf("Connections succesfully established!")
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
