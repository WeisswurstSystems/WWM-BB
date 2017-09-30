package mail

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

func init() {
	_, err := invitationTemplate.Parse(invitationMessageTemplate)
	if err != nil {
		panic(err)
	}
}

func NewNotificiationMail(
	notificationTopic string,
	notificationShortMessage string,
	notificationMessage string,
	meeting meeting.Meeting) (mail Mail) {

	smtpData := SmtpTemplateData{
		Subject:       notificationTopic,
		BodyShortText: notificationShortMessage,
		Body:          notificationMessage,
	}

	var receiverList []meeting.CustomerMail

	for _, orderItem := range meeting.Orders {
		receiverList = append(receiverList, orderItem.Customer)
	}

	var err error

	mail.Content, err = NewContent(smtpData)
	if err != nil {
		panic(err)
	}
	mail.Receivers = []string(receiverList)
	return mail
}
