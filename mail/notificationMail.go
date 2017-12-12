package mail

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/constants"
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
		BodyButtonLink: constants.BASEURL_GOTO_MEETING + string(meeting.ID),
		BodyButtonText: "Gehe zu Meeting",
		Subject:       notificationTopic,
		BodyShortText: notificationShortMessage,
		Body:          notificationMessage,
	}

	var receiverList []string

	for _, orderItem := range meeting.Orders {
		receiverList = append(receiverList, string(orderItem.Customer))
	}

	var err error

	mail.Content, err = NewContent(smtpData)
	if err != nil {
		panic(err)
	}
	mail.Receivers = receiverList
	return mail
}
