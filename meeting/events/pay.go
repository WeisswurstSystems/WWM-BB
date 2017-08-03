package events

import (
	"errors"
)

type Pay struct {
	MeetingID string `json:"meetingID"`
	UserID    string `json:"userID"`
}

func (event Pay) Execute() error {
	return errors.New("Not yet implemented!")
}
