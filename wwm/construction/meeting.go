package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter"
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter/command"
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter/query"
	"github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/closemeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/createmeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/putproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/removeproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setbuyer"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setplace"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"log"
	"os"
)

var MeetingStore = driver.NewMongoStore()

var MeetingUseCases = struct {
	createmeeting.CreateMeetingUseCase
	closemeeting.CloseMeetingUseCase
	putproduct.PutProductUseCase
	removeproduct.RemoveProductUseCase
	setbuyer.SetBuyerUseCase
	setplace.SetPlaceUseCase
}{
	createmeeting.Interactor{MeetingStore},
	closemeeting.Interactor{UserAuthentication, MeetingStore},
	putproduct.Interactor{MeetingStore, UserAuthentication},
	removeproduct.Interactor{MeetingStore, UserAuthentication},
	setbuyer.Interactor{MeetingStore, UserAuthentication},
	setplace.Interactor{MeetingStore, UserAuthentication},
}

var MeetingCommand = command.CommandHandler{
	Interactor: &MeetingUseCases,
}
var MeetingQuery = query.QueryHandler{
	MeetingStore: MeetingStore,
}
