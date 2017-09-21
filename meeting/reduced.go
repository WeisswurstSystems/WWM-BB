package meeting

import "time"

// ReducedMeeting contains only limited (most important) data of a meeting.
type ReducedMeeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
}

// Reduced version of the meeting.
func (m Meeting) Reduced() ReducedMeeting {
	return ReducedMeeting{
		ID:        m.ID,
		Place:     m.Place,
		Date:      m.Date,
		CloseDate: m.CloseDate,
		Closed:    m.Closed,
	}
}

// AllReduced of a meeting list.
func AllReduced(meetings []Meeting) []ReducedMeeting {
	list := make([]ReducedMeeting, 0, len(meetings))
	for _, v := range meetings {
		list = append(list, v.Reduced())
	}
	return list
}
