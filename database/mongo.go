package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

const LOG_TAG = "[DATABASE]"

func NewMongoSession() *mgo.Session {
	username := GetEnv("db.username", "wwm")
	password := GetEnv("db.password", "wwm")
	url := GetEnv("db.url", "ds064649.mlab.com:64649/wwmbb-dev")
	databaseName := GetEnv("db.name", "wwmbb-dev")

	log.Printf("%v Establishing connection to database <%v> on %v", LOG_TAG, databaseName, url)

	var err error
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%v:%v@%v", username, password, url))
	if err != nil {
		panic(err)
	}
	return session
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
