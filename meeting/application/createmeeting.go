package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"time"
)

type CreateMeeting struct {
	Place   string
	Date    time.Time
	Creator string
}

func (i *Interactor) CreateMeeting(e CreateMeeting) error {
	var m meeting.Meeting

	m.Creator = e.Creator
	m.Date = e.Date
	m.Place = e.Place
	m.ID = meeting.MeetingID(util.GetUID(12))

	err := i.MeetingStore.Save(m)
	if err != nil {
		return err
	}
	i.LOG.Printf("did %v", e)
	return nil
}
