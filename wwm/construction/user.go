package construction

import (
	mailDriver "github.com/WeisswurstSystems/WWM-BB/mail/driver"
	"github.com/WeisswurstSystems/WWM-BB/user/adapter/command"
	"github.com/WeisswurstSystems/WWM-BB/user/adapter/query"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/setUpPayPal"
)

var UserStore = driver.NewMongoStore()

var MailService = mailDriver.NewSMTPService()

var UserUseCases = struct {
	authenticate.AuthenticateUseCase
	activate.ActivateUseCase
	register.RegisterUseCase
	setUpPayPal.SetUpPayPalUseCase
}{
	AuthenticateUseCase: authenticate.Interactor{
		ReadStore: UserStore,
	},
	ActivateUseCase: activate.Interactor{
		Store: UserStore,
	},
	RegisterUseCase: register.Interactor{
		Store:       UserStore,
		MailService: MailService,
	},
}
var UserCommand = command.CommandHandler{
	Interactor: &UserUseCases,
}
var UserQuery = query.QueryHandler{
	Store: UserStore,
}
