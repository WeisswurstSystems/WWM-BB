package events

import (
	"errors"

	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type Update meeting.Meeting

func (event Update) Execute() error {
	return errors.New("Not yet implemented!")
}
