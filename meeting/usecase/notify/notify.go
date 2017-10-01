package notify

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"errors"
	"log"
)

type NotifyUseCase interface {
	Notify(request Request) error
}

type Interactor struct {
	meeting.ReadStore
	MailService mail.Service
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.MeetingID          `json:"meetingID"`
	NotifyTopic     string     `json:"topic"`
	NotifyShortText string     `json:"shortText"`
	NotifyLongText  string     `json:"longText"`
	Login           user.Login `json:"login"`
}

func (i Interactor) Notify(req Request) error {
	log.Printf("[%v] notify#Notify: Received request.", req.Login.Mail)
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

	mMail := mail.NewNotificiationMail(req.NotifyTopic, req.NotifyShortText, req.NotifyLongText, m)
	log.Printf("Sending E-Mail to the following list: %v", mMail.Receivers)
	err = i.MailService.Send(mMail)
	if err != nil {
		log.Fatalf("[%v] notify#Notify: Error sending Notification-Mail: %v", req.Login.Mail, err)
		return errors.New("Mail could not be send.")
	}

	log.Printf("[%v] notify#Notify: Request completed.", req.Login.Mail)
	return nil
}
