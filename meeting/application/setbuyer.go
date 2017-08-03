package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type SetBuyer struct {
	Buyer   string
	Meeting meeting.MeetingID
}

func (i *Interactor) SetBuyer(e SetBuyer) error {
	m, err := i.MeetingStore.FindOne(e.Meeting)
	if err != nil {
		return err
	}
	m.Buyer = e.Buyer
	err = i.MeetingStore.Save(m)
	if err != nil {
		return err
	}

	i.LOG.Printf("did %v", e)
	return nil
}
