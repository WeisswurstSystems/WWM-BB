package closemeeting

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/go-errors/errors"
)

type Mock struct {
	Meeting meeting.Meeting
	Saved   meeting.Meeting
}

func NewMocking() (CloseMeetingUseCase, *Mock) {
	mock := new(Mock)
	return Interactor{
		Store:          mock,
		Authentication: mock,
	}, mock
}

func (mock *Mock) CurrentUser() user.User {
	return user.User{
		Roles: []string{"admin"},
	}
}

func (Mock) FindAll() ([]meeting.Meeting, error) {
	panic("implement me")
}

func (Mock) FindAllReduced() ([]meeting.ReducedMeeting, error) {
	panic("implement me")
}

func (mock *Mock) FindOne(id meeting.MeetingID) (meeting.Meeting, error) {
	if id == mock.Meeting.ID {
		return mock.Meeting, nil
	}
	return meeting.Meeting{}, errors.New("meeting not found")
}

func (mock *Mock) Save(meeting meeting.Meeting) error {
	mock.Saved = meeting
	return nil
}
