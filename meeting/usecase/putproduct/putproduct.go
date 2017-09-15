package putproduct

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

type PutProductUseCase interface {
	PutProduct(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.MeetingID `json:"meetingID"`
	Product           meeting.Product `json:"product"`
	Login             user.Login      `json:"login"`
}

func (i Interactor) PutProduct(req Request) error {

	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}

	u, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	if !u.HasMail(m.Creator, m.Buyer) {
		return meeting.ErrNotAllowed
	}

	m.Offer = m.Offer.Put(req.Product)
	err = i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}
