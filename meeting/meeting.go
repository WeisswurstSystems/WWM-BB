package meeting

import (
	"net/http"
	"time"

	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// MeetingID uniquely identifies a meeting.
type MeetingID string

// Meeting describes a planned meeting for multiple persons.
type Meeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Creator   string    `json:"creator"`
	Buyer     string    `json:"buyer"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
	Orders    Orders    `json:"orders"`
	Offer     Offer     `json:"offer"`
}

// ReadStore can query meetings.
type ReadStore interface {
	FindAll() ([]Meeting, error)
	FindAllReduced() ([]ReducedMeeting, error)
	FindOne(id MeetingID) (Meeting, error)
}

// Store can read and write meetings.
type Store interface {
	ReadStore
	WriteStore
}

// WriteStore can save (or change) a meeting.
type WriteStore interface {
	Save(meeting Meeting) error
}

var (
	// ErrMeetingNotFound if a meeting is not found in a store.
	ErrMeetingNotFound = wwm.Error{Code: http.StatusNotFound, Message: "The meeting does not exist"}
	// ErrNotAllowed if a user is not allowed to do something on a meeting.
	ErrNotAllowed = wwm.Error{Code: http.StatusUnauthorized, Message: "Not allowed on this meeting"}
)
