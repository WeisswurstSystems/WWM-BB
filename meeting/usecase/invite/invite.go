package invite

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"errors"
	"fmt"
	"log"
)

type InviteUseCase interface {
	Invite(request Request) error
}

type Interactor struct {
	meeting.ReadStore
	MailService mail.Service
	authenticate.AuthenticateUseCase
}

type Request struct {
	meeting.MeetingID    `json:"meetingID"`
	UserMails []string   `json:"userMails"`
	Login     user.Login `json:"login"`
}

func (i Interactor) Invite(req Request) error {
	m, err := i.FindOne(req.MeetingID)
	if err != nil {
		return err
	}

	u, err := i.AuthenticateUseCase.Authenticate(req.Login)
	if err != nil {
		return err
	}

	if !u.HasMail(m.Creator) {
		return meeting.ErrNotAllowed
	}

	var failedMails []string

	for _, usermail := range req.UserMails {
		m := mail.NewInvitationMail(m, usermail)
		err := i.MailService.Send(m)
		if err != nil {
			log.Fatalf("Error sending mail to %v: %v", usermail, err)
			failedMails = append(failedMails, usermail)
		}
	}

	if len(failedMails) != 0 {
		return errors.New(fmt.Sprintf("Erorrs sending Mail to the followeing receipients: %v", failedMails))
	}

	return nil
}
