package application

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	"log"
	"os"
)

func NewMockInteractor() Interactor {
	return Interactor{
		MeetingStore: driver.NewMemoryStore(),
		LOG:          *log.New(os.Stdout, "TestCloseMeeting: ", log.LstdFlags),
	}
}
