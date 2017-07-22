package store

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoStore struct {
	session  *mgo.Session
	meetings *mgo.Collection
}

func NewMongoStore() meeting.Store {
	databaseName := database.GetEnv("DB_NAME", "wwmbb-dev")
	var store mongoStore

	store.session = database.NewMongoSession()
	store.session.SetMode(mgo.Monotonic, true)
	store.meetings = store.session.DB(databaseName).C("meetings")
	return &store
}

func (store *mongoStore) Count() (int, error) {
	return store.meetings.Find(nil).Count()
}

func (store *mongoStore) Has(id string) (bool, error) {
	count, err := store.meetings.Find(bson.M{"id": id}).Count()
	return count != 0, err
}

func (store *mongoStore) FindAll() ([]meeting.Meeting, error) {
	var results []meeting.Meeting
	err := store.meetings.Find(nil).All(&results)
	return results, err
}

func (store *mongoStore) FindAllReduced() ([]meeting.ReducedMeeting, error) {
	var results []meeting.ReducedMeeting
	err := store.meetings.Find(nil).All(&results)
	return results, err
}

func (store *mongoStore) FindOne(id string) (meeting.Meeting, error) {
	var result meeting.Meeting
	err := store.meetings.Find(bson.M{"id": id}).One(&result)
	return result, err
}

func (store *mongoStore) Save(meeting meeting.Meeting) error {
	err := store.meetings.Insert(&meeting)
	return err
}

func (store *mongoStore) Update(meeting meeting.Meeting) error {
	return store.meetings.Update(bson.M{"id": meeting.ID}, meeting)
}
