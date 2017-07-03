package store

import (
	"github.com/WeisswurstSystems/WWM-BB/database"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type Store interface {
	Count() (int, error)
	FindAll() ([]meeting.Meeting, error)
	Create(meeting meeting.Meeting) (meeting.Meeting, error)
}

func Count() (int, error) {
	return database.Meetings.Find(nil).Count()
}

func FindAll() ([]meeting.Meeting, error) {
	var results []meeting.Meeting
	err := database.Meetings.Find(nil).All(&results)
	return results, err
}

func Create(meeting meeting.Meeting) (meeting.Meeting, error) {
	err := database.Meetings.Insert(&meeting)
	return meeting, err
}
