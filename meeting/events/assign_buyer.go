package events

import (
	"errors"
)

type AssignBuyer struct {
	MeetingID string `json:"meetingID"`
	Buyer     string `json:"buyer"`
}

func (event AssignBuyer) Execute() error {
	return errors.New("Not yet implemented!")
}
