package construction

import (
	mailDriver "github.com/WeisswurstSystems/WWM-BB/mail/driver"
	"github.com/WeisswurstSystems/WWM-BB/user/adapter/command"
	"github.com/WeisswurstSystems/WWM-BB/user/adapter/query"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/changePassword"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/deleteAccount"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/setUpPayPal"
)

var UserStore = driver.NewMongoStore()

var MailService = mailDriver.NewSMTPService()

var UserAuthenticateUseCase = authenticate.Interactor{
	ReadStore: UserStore,
}

var UserUseCases = struct {
	authenticate.AuthenticateUseCase
	activate.ActivateUseCase
	register.RegisterUseCase
	setUpPayPal.SetUpPayPalUseCase
	changePassword.ChangePasswordUseCase
	deleteAccount.DeleteAccountUseCase
}{
	AuthenticateUseCase: UserAuthenticateUseCase,
	ActivateUseCase: activate.Interactor{
		Store: UserStore,
	},
	RegisterUseCase: register.Interactor{
		Store:       UserStore,
		MailService: MailService,
	},
	ChangePasswordUseCase: changePassword.Interactor{
		Store:               UserStore,
		AuthenticateUseCase: UserAuthenticateUseCase,
	},
	DeleteAccountUseCase: deleteAccount.Interactor{
		Store:               UserStore,
		AuthenticateUseCase: UserAuthenticateUseCase,
	},
}

var UserCommand = command.CommandHandler{
	Interactor: &UserUseCases,
}
var UserQuery = query.QueryHandler{
	Store: UserStore,
}
