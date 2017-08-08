package command

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/closemeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/createmeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/putproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/removeproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setbuyer"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setplace"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

type Interactor interface {
	closemeeting.CloseMeetingUseCase
	createmeeting.CreateMeetingUseCase
	putproduct.PutProductUseCase
	removeproduct.RemoveProductUseCase
	setbuyer.SetBuyerUseCase
	setplace.SetPlaceUseCase
}

type CommandHandler struct {
	Interactor
}

func (ch *CommandHandler) CloseMeeting(w http.ResponseWriter, req *http.Request) error {
	var e closemeeting.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.CloseMeeting(e)
}

func (ch *CommandHandler) CreateMeeting(w http.ResponseWriter, req *http.Request) error {
	var e createmeeting.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.CreateMeeting(e)
}

func (ch *CommandHandler) PutProduct(w http.ResponseWriter, req *http.Request) error {
	var e putproduct.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.PutProduct(e)
}

func (ch *CommandHandler) RemoveProduct(w http.ResponseWriter, req *http.Request) error {
	var e removeproduct.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.RemoveProduct(e)
}

func (ch *CommandHandler) SetBuyer(w http.ResponseWriter, req *http.Request) error {
	var e setbuyer.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetBuyer(e)
}

func (ch *CommandHandler) SetPlace(w http.ResponseWriter, req *http.Request) error {
	var e setplace.Request
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetPlace(e)
}
