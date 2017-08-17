package construction

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter/command"
	"github.com/WeisswurstSystems/WWM-BB/meeting/adapter/query"
	"github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/closemeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/createmeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/putproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/removeproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setbuyer"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setplace"
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
	closemeeting.Interactor{MeetingStore, UserUseCases.AuthenticateUseCase},
	putproduct.Interactor{MeetingStore, UserUseCases.AuthenticateUseCase},
	removeproduct.Interactor{MeetingStore, UserUseCases.AuthenticateUseCase},
	setbuyer.Interactor{MeetingStore, UserUseCases.AuthenticateUseCase},
	setplace.Interactor{MeetingStore, UserUseCases.AuthenticateUseCase},
}

var MeetingCommand = command.CommandHandler{
	Interactor: &MeetingUseCases,
}
var MeetingQuery = query.QueryHandler{
	MeetingStore: MeetingStore,
}