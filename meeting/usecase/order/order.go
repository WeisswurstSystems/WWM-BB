package order

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

// OrderUseCase can add a single order for a customer.
type OrderUseCase interface {
	Order(Request) error
}

// Request with the own login and the order item to add.
type Request struct {
	Login     user.Login        `json:"login"`
	Item      meeting.OrderItem `json:"orderItem"`
	MeetingID meeting.MeetingID `json:"meetingID"`
}

// Interactor for doing the order logic.
type Interactor struct {
	authenticate.AuthenticateUseCase
	meeting.Store
}

// Order adds the item to the costumers order in the given meeting.
func (i *Interactor) Order(req Request) error {
	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}

	u, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	m.AddOrderItemForCustomer(req.Item, meeting.CustomerMail(u.Mail))
	return i.Save(m)
}
