package webhandler

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/event"
)

type UserInteractor interface {
	Activate(e event.Activate) error
	Register(e event.Register) error
}

type CommandHandler struct {
	UserInteractor UserInteractor
}

type QueryHandler struct {
	UserStore user.Store
}
