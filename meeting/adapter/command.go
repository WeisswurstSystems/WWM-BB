package adapter

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting/application"
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
)

func (ch *CommandHandler) CloseMeeting(w http.ResponseWriter, req *http.Request) error {
	var e application.CloseMeeting
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.CloseMeeting(e)
}

func (ch *CommandHandler) CreateMeeting(w http.ResponseWriter, req *http.Request) error {
	var e application.CreateMeeting
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.CreateMeeting(e)
}

func (ch *CommandHandler) PutProduct(w http.ResponseWriter, req *http.Request) error {
	var e application.PutProduct
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.PutProduct(e)
}

func (ch *CommandHandler) RemoveProduct(w http.ResponseWriter, req *http.Request) error {
	var e application.RemoveProduct
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.RemoveProduct(e)
}

func (ch *CommandHandler) SetBuyer(w http.ResponseWriter, req *http.Request) error {
	var e application.SetBuyer
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.SetBuyer(e)
}

func (ch *CommandHandler) SetPlace(w http.ResponseWriter, req *http.Request) error {
	var e application.SetPlace
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.MeetingInteractor.SetPlace(e)
}
