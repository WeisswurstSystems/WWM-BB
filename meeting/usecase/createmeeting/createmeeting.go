package createmeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/user"
)

type CreateMeetingUseCase interface {
	CreateMeeting(request Request) error
}

type Interactor struct {
	meeting.WriteStore
	authenticate.AuthenticateUseCase
}

type Request struct {
	Meeting meeting.Meeting `json:"meeting"`
	Login user.Login  `json:"login"`
}

func (i Interactor) CreateMeeting(req Request) error {
	m := req.Meeting

	_, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	m.ID = meeting.MeetingID(util.GetUID(12))

	err = i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}
