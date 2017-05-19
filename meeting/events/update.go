package events

import (
	"errors"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type Update struct {
	Meeting meeting.Meeting `json:"meeting"`
}

func (event Update) Execute() error {
	return errors.New("Not yet implemented!")
}
