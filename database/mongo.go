package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

const LOG_TAG = "[DATABASE]"

func NewMongoSession() *mgo.Session {
	username := GetEnv("DB_USERNAME", "wwm")
	password := GetEnv("DB_PASSWORD", "wwm")
	url := GetEnv("DB_URL", "ds064649.mlab.com:64649/wwmbb-dev")
	databaseName := GetEnv("DB_NAME", "wwmbb-dev")

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
