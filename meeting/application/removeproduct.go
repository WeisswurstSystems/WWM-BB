package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type RemoveProduct struct {
	ProductName meeting.ProductName
	Meeting     meeting.MeetingID
}

func (i *Interactor) RemoveProduct(e RemoveProduct) error {
	m, err := i.MeetingStore.FindOne(e.Meeting)
	if err != nil {
		return err
	}

	m.Offer = m.Offer.Remove(e.ProductName)
	err = i.MeetingStore.Save(m)
	if err != nil {
		return err
	}
	i.LOG.Printf("did %v", e)
	return nil
}
