package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type CloseMeeting struct {
	Meeting meeting.MeetingID
}

func (i *Interactor) CloseMeeting(e CloseMeeting) error {
	m, err := i.MeetingStore.FindOne(e.Meeting)
	if err != nil {
		return err
	}

	m.Closed = true
	err = i.MeetingStore.Save(m)
	if err != nil {
		return err
	}

	i.LOG.Printf("did %v", e)
	return nil
}
