package setbuyer

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

type SetBuyerUseCase interface {
	SetBuyer(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}
type Request struct {
	Buyer             string `json:"buyer"`
	meeting.MeetingID `json:"meetingID"`
	user.Login        `json:"login"`
}

func (i Interactor) SetBuyer(req Request) error {
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

	m.Buyer = req.Buyer
	err = i.Save(m)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("did %v", req)
	return nil
}
