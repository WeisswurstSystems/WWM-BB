package mail

import (
	"bytes"
	"text/template"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/constants"
)

type invitationData struct {
	Usermail       string
	MeetingCreator string
	MeetingDate    string
	MeetingPlace   string
}

const invitationTopic = "Du wurdest eingeladen!"

const invitationMessageTemplate = `
Hallo {{.Usermail}}!
<br><br>
{{.MeetingCreator}} hat dich zu einem Treffen eingeladen.
<br><br>
<b>Wann?</b> {{.MeetingDate}}<br>
<b>Wo?</b> {{.MeetingPlace}}
<br><br>
Unter dem folgenden Link kannst du beitreten und deine Bestellung aufgeben.

Viel Spaß beim bestellen!
`

var invitationTemplate = template.New("invitation")

func init() {
	_, err := invitationTemplate.Parse(invitationMessageTemplate)
	if err != nil {
		panic(err)
	}
}

func NewInvitationMail(meeting meeting.Meeting, usermail string) (mail Mail) {
	data := invitationData{
		Usermail: usermail,
		MeetingCreator:meeting.Creator,
		MeetingDate: meeting.Date.Format("02.01.2006 15:04"),
		MeetingPlace: meeting.Place,
	}

	var message bytes.Buffer
	err := invitationTemplate.Execute(&message, data)
	if err != nil {
		panic(err)
	}

	smtpData := SmtpTemplateData{
		BodyButtonLink: constants.BASEURL_GOTO_MEETING + string(meeting.ID),
		BodyButtonText: "Meeting beitreten",
		Subject: invitationTopic,
		BodyShortText: "Du wurdest von " + usermail + " zu einem Meeting eingeladen",
		Body: message.String(),
	}

	mail.Content, err = NewContent(smtpData)
	if err != nil {
		panic(err)
	}
	mail.Receivers = []string{usermail}
	return mail
}