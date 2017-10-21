package mail

import (
	"bytes"
	"text/template"
	"github.com/WeisswurstSystems/WWM-BB/constants"
)

type registrationData struct {
	Usermail       string
}

const registrationTopic = "Deine Registrierung"

const registrationMessageTemplate = `
Hallo {{.Usermail}}!
<br><br>
Vielen Dank für deine Registrierung.<br>
Um Dein Konto zu aktivieren, klicke bitte auf den unten stehenden Button.
<br><br>
Viel Spaß beim bestellen!
`

var registrationTemplate = template.New("registration")

func init() {
	_, err := registrationTemplate.Parse(registrationMessageTemplate)
	if err != nil {
		panic(err)
	}
}

func NewRegistrationMail(registrationId string, usermail string) (mail Mail) {
	data := registrationData{Usermail: usermail}

	var message bytes.Buffer
	err := registrationTemplate.Execute(&message, data)
	if err != nil {
		panic(err)
	}

	smtpData := SmtpTemplateData{
		BodyButtonLink: constants.BASEURL_ACTIVATE_ACCOUNT + registrationId,
		BodyButtonText: "Aktivieren",
		Subject: registrationTopic,
		BodyShortText: "Um deine Registrierung zu bestätigen, musst du nur noch eine Sache erledigen.",
		Body: message.String(),
	}

	mail.Content, err = NewContent(smtpData)
	if err != nil {
		panic(err)
	}
	mail.Receivers = []string{usermail}
	return mail
}
