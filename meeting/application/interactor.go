package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"log"
)

type Interactor struct {
	MeetingStore meeting.Store
	LOG          log.Logger
}
