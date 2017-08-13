package putproduct

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type PutProductUseCase interface {
	PutProduct(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.Product
	meeting.MeetingID
	user.Login
}

func (i Interactor) PutProduct(req Request) error {
	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}
	user, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	if !isAllowed(user, m) {
		return meeting.ErrNotAllowed
	}

	m.Offer = m.Offer.Put(req.Product)
	err = i.Save(m)
	if err != nil {
		return err
	}
	usecase.LOG.Printf("did %v", req)
	return nil
}

func isAllowed(u user.User, m meeting.Meeting) bool {
	if u.Mail == m.Creator {
		return true
	}
	if u.Mail == m.Buyer {
		return true
	}
	return util.Contains(u.Roles, "admin")
}
