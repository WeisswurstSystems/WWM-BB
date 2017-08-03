package events

import (
	"errors"
)

type Notify struct {
	MeetingID string `json:"meetingID"`
	Message   string `json:"message"`
}

func (event Notify) Execute() error {
	return errors.New("Not yet implemented!")
}
