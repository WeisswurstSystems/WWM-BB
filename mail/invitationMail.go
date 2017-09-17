package mail

import (
	"bytes"
	"text/template"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
)

type invitationData struct {
	Usermail       string
	MeetingCreator string
	MeetingID      string
	MeetingDate    string
	MeetingPlace   string
}

const invitationTopic = "Einladung zum Meeting am {{.MeetingPlace}}"

const invitationMessageTemplate = `
Hallo {{.Usermail}}!

{{.MeetingCreator}} hat dich zu einem Treffen eingeladen.

Wann? {{.MeetingDate}}
Wo? {{.MeetingDate}}

Unter dem folgenden Link kannst du beitreten und deine Bestellung aufgeben
https://weisswurstsystems.github.io/WWM-ITM/{{.MeetingID}}

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
		MeetingID: string(meeting.ID),
		Usermail: usermail,
		MeetingCreator:meeting.Creator,
		MeetingDate: meeting.Date.Format("2006-01-02 15:04:05"),
		MeetingPlace: meeting.Place,
	}

	var message bytes.Buffer
	err := invitationTemplate.Execute(&message, data)
	if err != nil {
		panic(err)
	}

	mail.Content, err = NewContent(invitationTopic, message.String())
	if err != nil {
		panic(err)
	}
	mail.Receivers = []string{usermail}
	return mail
}