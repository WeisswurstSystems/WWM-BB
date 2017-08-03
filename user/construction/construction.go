package construction

import (
	mail "github.com/WeisswurstSystems/WWM-BB/mail/construction"

	"github.com/WeisswurstSystems/WWM-BB/user/adapter"
	"github.com/WeisswurstSystems/WWM-BB/user/application"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
)

var Store = driver.NewMongoStore()

var Interactor = application.Interactor{
	UserStore:   Store,
	MailService: mail.MailService,
}

var Command = adapter.CommandHandler{
	UserInteractor: &Interactor,
}
var Query = adapter.QueryHandler{
	UserStore: Store,
}
var Middleware = adapter.MiddlewareHandler{
	UserStore: Store,
}
