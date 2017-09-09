package meeting

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
	"time"
)

type MeetingID string

type Meeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Creator   string    `json:"creator"`
	Buyer     string    `json:"buyer"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
	Orders    []Order   `json:"orders"`
	Offer     Offer     `json:"offer"`
}

type ReducedMeeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
}

func (m Meeting) Reduced() ReducedMeeting {
	return ReducedMeeting{
		ID:        m.ID,
		Place:     m.Place,
		Date:      m.Date,
		CloseDate: m.CloseDate,
		Closed:    m.Closed,
	}
}

func AllReduced(meetings []Meeting) []ReducedMeeting {
	list := make([]ReducedMeeting, 0, len(meetings))
	for _, v := range meetings {
		list = append(list, v.Reduced())
	}
	return list
}

type ReadStore interface {
	FindAll() ([]Meeting, error)
	FindAllReduced() ([]ReducedMeeting, error)
	FindOne(id MeetingID) (Meeting, error)
}

type Store interface {
	ReadStore
	WriteStore
}

type WriteStore interface {
	Save(meeting Meeting) error
}

var (
	ErrMeetingNotFound = wwm.Error{Code: http.StatusNotFound, Message: "The meeting does not exist"}
	ErrNotAllowed      = wwm.Error{Code: http.StatusUnauthorized, Message: "Not allowed on this meeting"}
)
