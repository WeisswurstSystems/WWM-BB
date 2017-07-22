package user

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm/mail"

	"github.com/WeisswurstSystems/WWM-BB/user/middleware"
	"github.com/WeisswurstSystems/WWM-BB/user/store"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user/webhandler"
)

var Store = store.NewMongoStore()

var Interactor = usecase.Interactor{
	UserStore:   Store,
	MailService: mail.MailService,
}

var Command = webhandler.CommandHandler{
	UserInteractor: Interactor,
}

var Query = webhandler.QueryHandler{
	UserStore: Store,
}
var Middleware = middleware.MiddlewareHandler{
	UserStore: Store,
}
