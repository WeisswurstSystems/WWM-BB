package adapter

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/application"
)

type MiddlewareHandler struct {
	UserStore user.Store
}

type UserInteractor interface {
	Activate(e application.Activate) error
	Register(e application.Register) error
}

type CommandHandler struct {
	UserInteractor UserInteractor
}

type QueryHandler struct {
	UserStore user.Store
}
