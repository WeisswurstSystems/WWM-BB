package createmeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type CreateMeetingUseCase interface {
	CreateMeeting(request Request) error
}

type Interactor struct {
	meeting.WriteStore
}

type Request struct {
	Meeting meeting.Meeting `json:"meeting"`
}

func (i Interactor) CreateMeeting(req Request) error {
	m := req.Meeting

	m.ID = meeting.MeetingID(util.GetUID(12))

	err := i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}
