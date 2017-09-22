package query

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"errors"
)

func NewMocking() (QueryHandler, *Mock) {
	mock := new(Mock)
	return QueryHandler{MeetingStore: mock}, mock
}

type Mock struct {
	All []meeting.Meeting
	One meeting.Meeting
}

func (mock *Mock) FindAll() ([]meeting.Meeting, error) {
	return mock.All, nil
}

func (mock *Mock) FindAllReduced() ([]meeting.ReducedMeeting, error) {
	return meeting.AllReduced(mock.All), nil
}

func (mock *Mock) FindOne(id meeting.MeetingID) (meeting.Meeting, error) {
	for _, m := range mock.All {
		if m.ID == id {
			return m, nil
		}
	}
	return meeting.Meeting{}, errors.New("id not found")
}
