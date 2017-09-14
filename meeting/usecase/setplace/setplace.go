package setplace

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type SetPlaceUseCase interface {
	SetPlace(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	Place             string `json:"place"`
	meeting.MeetingID `json:"meetingID"`
	user.Login        `json:"login"`
}

func (i Interactor) SetPlace(req Request) error {
	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}
	user, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	if !util.IsMeetingCreatorOrBuyer(user, m) {
		return meeting.ErrNotAllowed
	}

	m.Place = req.Place
	err = i.Save(m)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("did %v", req)
	return nil
}
