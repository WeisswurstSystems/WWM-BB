package database

import (
	"gopkg.in/mgo.v2"
)

var (
	DBSession         *mgo.Session
	UserCollection    *mgo.Collection
	MeetingCollection *mgo.Collection
	PeopleCollection  *mgo.Collection
)
