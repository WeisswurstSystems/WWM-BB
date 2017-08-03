package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter"
	"github.com/WeisswurstSystems/WWM-BB/meeting/application"
	"github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	"log"
	"os"
)

var Store = driver.NewMongoStore()

var Interactor = application.Interactor{
	MeetingStore: Store,
	LOG:          *log.New(os.Stdout, "TestCloseMeeting: ", log.LstdFlags),
}

var Command = adapter.CommandHandler{
	MeetingInteractor: &Interactor,
}
var Query = adapter.QueryHandler{
	MeetingStore: Store,
}
