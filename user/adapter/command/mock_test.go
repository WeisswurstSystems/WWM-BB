package command

import (
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/changePassword"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/deleteAccount"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/setUpPayPal"
)

type Mock struct {
	Requests struct {
		Register       register.Request
		Activate       activate.Request
		SetUpPayPal    setUpPayPal.Request
		ChangePassword changePassword.Request
		DeleteAccount  deleteAccount.Request
	}
}

func NewMockCommandHandler() (CommandHandler, *Mock) {
	mock := new(Mock)
	return CommandHandler{
		Interactor: mock,
	}, mock
}

func (m *Mock) Register(req register.Request) error {
	m.Requests.Register = req
	return nil
}

func (m *Mock) Activate(req activate.Request) error {
	m.Requests.Activate = req
	return nil
}

func (m *Mock) SetUpPayPal(req setUpPayPal.Request) error {
	m.Requests.SetUpPayPal = req
	return nil
}

func (m *Mock) ChangePassword(req changePassword.Request) error {
	m.Requests.ChangePassword = req
	return nil
}

func (m *Mock) DeleteAccount(req deleteAccount.Request) error {
	m.Requests.DeleteAccount = req
	return nil
}
