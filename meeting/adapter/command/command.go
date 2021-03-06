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
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/toggleorderpayed"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/invite"
	"github.com/gorilla/mux"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/notify"
	"log"
)

type Interactor interface {
	closemeeting.CloseMeetingUseCase
	createmeeting.CreateMeetingUseCase
	putproduct.PutProductUseCase
	removeproduct.RemoveProductUseCase
	setbuyer.SetBuyerUseCase
	setplace.SetPlaceUseCase
	invite.InviteUseCase
	notify.NotifyUseCase
	toggleorderpayed.ToggleOrderPayedUseCase
}

type CommandHandler struct {
	Interactor
}

func (ch *CommandHandler) CloseMeeting(w http.ResponseWriter, req *http.Request) error {
	var e closemeeting.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.CloseMeeting(e)
}

func (ch *CommandHandler) CreateMeeting(w http.ResponseWriter, req *http.Request) error {
	var e createmeeting.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	err := wwm.DecodeBody(req.Body, &e)
	e.Meeting.Creator = e.Login.Mail
	if err != nil {
		return err
	}
	return ch.Interactor.CreateMeeting(e)
}

func (ch *CommandHandler) PutProduct(w http.ResponseWriter, req *http.Request) error {
	var e putproduct.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.PutProduct(e)
}

func (ch *CommandHandler) RemoveProduct(w http.ResponseWriter, req *http.Request) error {
	var e removeproduct.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.RemoveProduct(e)
}

func (ch *CommandHandler) SetBuyer(w http.ResponseWriter, req *http.Request) error {
	var e setbuyer.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetBuyer(e)
}

func (ch *CommandHandler) SetPlace(w http.ResponseWriter, req *http.Request) error {
	var e setplace.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.SetPlace(e)
}

func (ch *CommandHandler) ToggleOrderPayed(w http.ResponseWriter, req *http.Request) error {
	var e toggleorderpayed.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = getIDFromRequest(req)

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.ToggleOrderPayed(e)
}

func (ch *CommandHandler) Invite(w http.ResponseWriter, req *http.Request) error {
	var e invite.Request
	e.MeetingID = getIDFromRequest(req)

	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.Invite(e)
}

func (ch *CommandHandler) Notify(w http.ResponseWriter, req *http.Request) error {
	var e notify.Request
	e.Login.Mail, e.Login.Password, _ = req.BasicAuth()
	e.MeetingID = meeting.MeetingID(getIDFromRequest(req))

	err := wwm.DecodeBody(req.Body, &e)
	if err != nil {
		return err
	}
	return ch.Interactor.Notify(e)
}

func getIDFromRequest(request *http.Request) meeting.MeetingID {
	id, ok := mux.Vars(request)["meetingId"]
	if !ok {
		log.Fatal("Request had no meetingId in path!")
	}

	return meeting.MeetingID(id)
}
