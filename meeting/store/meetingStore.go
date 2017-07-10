package store

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"gopkg.in/mgo.v2/bson"
)

type Store interface {
	Count() (int, error)
	Has(id string) (bool, error)
	FindAll() ([]meeting.Meeting, error)
	FindAllReduced() ([]meeting.ReducedMeeting, error)
	FindOne(id string) (meeting.Meeting, error)
	Create(meeting meeting.Meeting) (meeting.Meeting, error)
	Save(meeting meeting. Meeting) error
}

func Count() (int, error) {
	return database.Meetings.Find(nil).Count()
}

func Has(id string) (bool, error) {
	count, err := database.Meetings.Find(bson.M{"id": id}).Count()
	return count != 0, err
}

func FindAll() ([]meeting.Meeting, error) {
	var results []meeting.Meeting
	err := database.Meetings.Find(nil).All(&results)
	return results, err
}

func FindAllReduced() ([]meeting.ReducedMeeting, error) {
	var results []meeting.ReducedMeeting
	err := database.Meetings.Find(nil).All(&results)
	return results, err
}

func FindOne(id string) (meeting.Meeting, error) {
	var result meeting.Meeting
	err := database.Meetings.Find(bson.M{"id": id}).One(&result)
	return result, err
}


func Create(meeting meeting.Meeting) (meeting.Meeting, error) {
	err := database.Meetings.Insert(&meeting)
	return meeting, err
}


func Save(meeting meeting. Meeting) error {
	_, err := database.Meetings.UpsertId(meeting.ID, bson.M{"$set": meeting})
	return err
}