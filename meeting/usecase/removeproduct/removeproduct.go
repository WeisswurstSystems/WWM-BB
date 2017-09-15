package removeproduct

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

type RemoveProductUseCase interface {
	RemoveProduct(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.ProductName `json:"productName"`
	meeting.MeetingID   `json:"meetingID"`
	Login               user.Login `json:"login"`
}

func (i Interactor) RemoveProduct(req Request) error {
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

	m.Offer = m.Offer.Remove(req.ProductName)
	err = i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}
