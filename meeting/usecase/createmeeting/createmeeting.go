package createmeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"time"
)

type CreateMeetingUseCase interface {
	CreateMeeting(request Request) error
}

type Interactor struct {
	meeting.WriteStore
}

type Request struct {
	Place   string
	Date    time.Time
	Creator string
}

func (i Interactor) CreateMeeting(req Request) error {
	var m meeting.Meeting

	m.Creator = req.Creator
	m.Date = req.Date
	m.Place = req.Place
	m.ID = meeting.MeetingID(util.GetUID(12))

	err := i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}
