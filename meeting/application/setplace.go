package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type SetPlace struct {
	Place   string
	Meeting meeting.MeetingID
}

func (i *Interactor) SetPlace(e SetPlace) error {
	m, err := i.MeetingStore.FindOne(e.Meeting)
	if err != nil {
		return err
	}
	m.Place = e.Place
	err = i.MeetingStore.Save(m)
	if err != nil {
		return err
	}

	i.LOG.Printf("did %v", e)
	return nil
}
