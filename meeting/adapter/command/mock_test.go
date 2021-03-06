package command

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/closemeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/createmeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/putproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/removeproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setbuyer"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setplace"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/toggleorderpayed"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/invite"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/notify"
)

type Mock struct {
	Requests struct {
		CreateMeeting    createmeeting.Request
		CloseMeeting     closemeeting.Request
		PutProduct       putproduct.Request
		RemoveProduct    removeproduct.Request
		SetBuyer         setbuyer.Request
		SetPlace         setplace.Request
		Invite           invite.Request
		Notify           notify.Request
		ToggleOrderPayed toggleorderpayed.Request
	}
}

func NewMockCommandHandler() (CommandHandler, *Mock) {
	mock := new(Mock)
	return CommandHandler{
		Interactor: mock,
	}, mock
}

func (mock *Mock) CloseMeeting(request closemeeting.Request) error {
	mock.Requests.CloseMeeting = request
	return nil
}

func (mock *Mock) CreateMeeting(request createmeeting.Request) error {
	mock.Requests.CreateMeeting = request
	return nil
}

func (mock *Mock) PutProduct(request putproduct.Request) error {
	mock.Requests.PutProduct = request
	return nil
}

func (mock *Mock) RemoveProduct(request removeproduct.Request) error {
	mock.Requests.RemoveProduct = request
	return nil
}

func (mock *Mock) SetBuyer(request setbuyer.Request) error {
	mock.Requests.SetBuyer = request
	return nil
}

func (mock *Mock) SetPlace(request setplace.Request) error {
	mock.Requests.SetPlace = request
	return nil
}

func (mock *Mock) Invite(request invite.Request) error {
	mock.Requests.Invite = request
	return nil
}

func (mock *Mock) Notify(request notify.Request) error {
	mock.Requests.Notify = request
	return nil
}

func (mock *Mock) ToggleOrderPayed(request toggleorderpayed.Request) error {
	mock.Requests.ToggleOrderPayed = request
	return nil
}
