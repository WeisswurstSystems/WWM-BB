package driver

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type memoryStore struct {
	list map[meeting.MeetingID]meeting.Meeting
}

func NewMemoryStore() meeting.Store {
	store := memoryStore{
		list: make(map[meeting.MeetingID]meeting.Meeting),
	}
	return &store
}

func (m *memoryStore) Count() (int, error) {
	return len(m.list), nil
}

func (m *memoryStore) Has(id meeting.MeetingID) (bool, error) {
	_, ok := m.list[id]
	return ok, nil
}

func (m *memoryStore) FindAll() ([]meeting.Meeting, error) {
	meetings := make([]meeting.Meeting, 0, len(m.list))
	for _, v := range m.list {
		meetings = append(meetings, v)
	}
	return meetings, nil
}

func (m *memoryStore) FindAllReduced() ([]meeting.ReducedMeeting, error) {
	meetings := make([]meeting.ReducedMeeting, 0, len(m.list))
	for _, v := range m.list {
		meetings = append(meetings, v.Reduced())
	}
	return meetings, nil
}

func (m *memoryStore) FindOne(id meeting.MeetingID) (meeting.Meeting, error) {
	meet, ok := m.list[id]
	if !ok {
		return meeting.Meeting{}, meeting.ErrMeetingNotFound
	}
	return meet, nil
}

func (m *memoryStore) Save(meeting meeting.Meeting) error {
	m.list[meeting.ID] = meeting
	return nil
}
