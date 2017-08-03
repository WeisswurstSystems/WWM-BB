package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type PutProduct struct {
	Product meeting.Product
	Meeting meeting.MeetingID
}

func (i *Interactor) PutProduct(e PutProduct) error {
	m, err := i.MeetingStore.FindOne(e.Meeting)
	if err != nil {
		return err
	}

	m.Offer = m.Offer.Put(e.Product)
	err = i.MeetingStore.Save(m)
	if err != nil {
		return err
	}
	i.LOG.Printf("did %v", e)
	return nil
}
