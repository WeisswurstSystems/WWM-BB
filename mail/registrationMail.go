package mail

import (
	"bytes"
	"text/template"
)

type registrationData struct {
	Usermail       string
	RegistrationID string
}

const registrationTopic = "Deine Registrierung bei der Weisswurst-Verwaltung"

const registrationMessageTemplate = `
Hallo {{.Usermail}}!

Vielen Dank für deine Registrierung.
Um Dein Konto zu aktivieren, klicke bitte auf den folgenden Link:
http://wwm-bb.herokuapp.com/users/register/{{.RegistrationID}}

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
	data := registrationData{Usermail: usermail, RegistrationID: registrationId}

	var message bytes.Buffer
	err := registrationTemplate.Execute(&message, data)
	if err != nil {
		panic(err)
	}

	mail.Content, err = NewContent(registrationTopic, message.String())
	if err != nil {
		panic(err)
	}
	mail.Receivers = []string{usermail}
	return mail
}
