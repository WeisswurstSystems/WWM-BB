package construction

import (
	mailDriver "github.com/WeisswurstSystems/WWM-BB/mail/driver"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/adapter"
	"github.com/WeisswurstSystems/WWM-BB/user/application"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
)

var UserStore = driver.NewMongoStore()

var MailService = mailDriver.NewSMTPService()

var UserAuthentication user.Authentication

var UserInteractor = application.Interactor{
	UserStore:   UserStore,
	MailService: MailService,
}

var UserCommand = adapter.CommandHandler{
	UserInteractor: &UserInteractor,
}
var UserQuery = adapter.QueryHandler{
	UserStore: UserStore,
}
var UserMiddleware = adapter.MiddlewareHandler{
	UserStore: UserStore,
}
