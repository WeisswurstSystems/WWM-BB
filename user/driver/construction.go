package driver

import (
	mail "github.com/WeisswurstSystems/WWM-BB/mail/driver"

	"github.com/WeisswurstSystems/WWM-BB/user/adapter"
	"github.com/WeisswurstSystems/WWM-BB/user/application"
)

var Store = NewMongoStore()

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
