package removeproduct

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type RemoveProductUseCase interface {
	RemoveProduct(request Request) error
}

type Interactor struct {
	meeting.Store
	user.Authentication
}

type Request struct {
	ProductName meeting.ProductName
	Meeting     meeting.MeetingID
}

func (i Interactor) RemoveProduct(req Request) error {
	m, err := i.FindOne(req.Meeting)
	if err != nil {
		return err
	}

	if !isAllowed(i.CurrentUser(), m) {
		return meeting.ErrNotAllowed
	}

	m.Offer = m.Offer.Remove(req.ProductName)
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
