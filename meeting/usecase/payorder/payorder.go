package payorder

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"errors"
	"fmt"
)

type PayOrderUsecase interface {
	PayOrder(request Request) error
}

type Interactor struct {
	meeting.Store
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.MeetingID `json:"meetingID"`
	Order meeting.Order `json:"order"`
	Login user.Login  `json:"login"`
}

func (i Interactor) PayOrder(req Request) error {
	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}

	u, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	if !u.HasMail(m.Creator, m.Buyer) {
		return meeting.ErrNotAllowed
	}

	index := util.IndexOf(len(m.Orders), func(i int) bool {
		return m.Orders[i].Customer == req.Order.Customer
	})

	if index == -1 {
		return errors.New(fmt.Sprintf("Order with customer %v not found in Meeting %v ", req.Order.Customer, m.ID))
	}

	m.Orders[index].Payed = true
	err = i.Save(m)
	if err != nil {
		return err
	}

	usecase.LOG.Printf("did %v", req)
	return nil
}
