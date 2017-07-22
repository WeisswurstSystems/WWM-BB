package usecase

import (
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/user"
)

type Interactor struct {
	UserStore   user.Store
	MailService mail.Service
}
