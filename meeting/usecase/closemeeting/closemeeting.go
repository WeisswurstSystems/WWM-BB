package closemeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type CloseMeetingUseCase interface {
	CloseMeeting(request Request) error
}

type Interactor struct {
	user.Authentication
	meeting.Store
}

type Request struct {
	Meeting meeting.MeetingID
}

func (i Interactor) CloseMeeting(req Request) error {
	m, err := i.FindOne(req.Meeting)
	if err != nil {
		return err
	}

	if !isAllowed(i.CurrentUser(), m) {
		return meeting.ErrNotAllowed
	}

	m.Closed = true
	err = i.Save(m)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("did %v", req)
	return nil
}

func isAllowed(u user.User, m meeting.Meeting) bool {
	if u.Mail == m.Creator {
		return true
	}
	return util.Contains(u.Roles, "admin")
}
